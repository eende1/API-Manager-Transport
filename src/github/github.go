package github

import (
	"errors"
	"archive/zip"
	"fmt"
	"gopkg.in/src-d/go-billy.v4"
	"gopkg.in/src-d/go-billy.v4/memfs"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	githttp "gopkg.in/src-d/go-git.v4/plumbing/transport/http"
	"gopkg.in/src-d/go-git.v4/storage/memory"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"io"
	"path/filepath"
	"time"
	"sync"
	"tenant"
	"apiinfo"
)

type Sync struct {
	Proxies map[string][]byte
	TenantName string
	LogMessage string
}

type Repo struct {
	fs billy.Filesystem
	r *git.Repository
	worktree *git.Worktree
}

func GithubTenantSync(tenantLock *tenant.Lock, syncIn chan Sync, syncOut chan error) {
	tenantName := "sandbox"
	for {
		time.Sleep(20 * time.Second)
		fmt.Println(tenantName)
		// Acquire lock on tenant
		lock, ok := (*tenantLock).Map[tenantName]
		if !ok {
			continue
		}
		(*lock).Lock()

		apiProxies, err := getAllAPIZip(tenantName, os.Getenv("SCPI_AUTH"))
		if err != nil {
			(*lock).Unlock()
			tenantName = tenant.Advance(tenantName)
			continue
		}
		syncIn <- Sync{apiProxies, tenantName, ""}
		<- syncOut
		(*lock).Unlock()
		tenantName = tenant.Advance(tenantName)
	}
}

func StartGithubHandler(syncIn chan Sync, syncOut chan error) {
	apimRepo, err := initializeGithubRepo();
	if err != nil {
		panic("Failed to initialize github repository.")
	}

	for {
		toSync := <- syncIn
		apimRepo.SyncAPIs(toSync.Proxies, toSync.TenantName, toSync.LogMessage)
		syncOut <- nil
	}
}

func (g *Repo) SyncAPIs(proxies map[string]([]byte), tenantName, logMessage string) {	
	for name, proxy := range proxies {
		err := unzip(proxy, &(g.fs), fmt.Sprintf("%s/%s", tenantName, name))
		if err != nil {
			log.Printf("Failed to unzip in SyncAPIS\n")
		}
		// Include a log message unless an empty message was passed.
		if logMessage != "" {
			err = logChange(logMessage, fmt.Sprintf("%s/%s/%s", tenantName, name, "change_log.txt"), &(g.fs))
			if err != nil {
				log.Printf("Failed to log change in SyncAPIS\n")
				return
			}
		}
	}
	g.worktree.AddGlob("*")

	_, err := g.worktree.Commit("example go-git commit", &git.CommitOptions{
		Author: &object.Signature{
			Name: "Backup Process",
			When: time.Now(),
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	err = g.r.Push(&git.PushOptions{
		Auth: &githttp.BasicAuth{
			Username: "tandr9",
			Password: os.Getenv("GITHUB_TOKEN"),
		},
	})
	if err != nil {
		fmt.Println(err)
	}
}

func initializeGithubRepo() (Repo, error){
	fs := memfs.New()
	storer := memory.NewStorage()

	g, err := git.Clone(storer, fs, &git.CloneOptions{
		URL: "https://api.github.nike.com/scp/APIM-Backup",
		Auth: &githttp.BasicAuth{
			Username: "tandr9",
			Password: os.Getenv("GITHUB_TOKEN"),
		},
	})
	if err != nil {
		return Repo{fs, g, nil}, err
	}

	w, err := g.Worktree()

	return Repo{fs, g, w}, err
}

type safeMap struct {
	mux sync.Mutex
	m   map[string][]byte
}

func getAllAPIZip(tenantName, auth string) (map[string][]byte, error) {
	APIProxies, err := apiinfo.GetAllAPINames(tenantName, auth)
	if err != nil {
		return nil, err
	}
	done := make(chan bool, len(APIProxies.APIs))
	//m := make(map[string][]byte)
	sm := safeMap{m: make(map[string][]byte)}
	for _, a := range APIProxies.APIs {
		go getAPIZip(done, tenantName, a.Name, auth, &sm)
	}

	success := true
	for i := 0; i < len(APIProxies.APIs); i++ {
		success = success && <-done
	}
	close(done)
	if !success {
		return nil, errors.New("Encountered error getting an API Zip")
	}
	return sm.m, nil
}

func getAPIZip(c chan bool, tenantName, apiName, auth string, sm *safeMap) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://produs2apiportalapimgmtpphx-%s.us2.hana.ondemand.com/apiportal/api/1.0/Transport.svc/APIProxies?name=%s", tenant.Get(tenantName), apiName), nil)
	if err != nil {
		c <- false
		return
	}

	req.Header.Add("Authorization", "Basic "+auth)
	resp, err := client.Do(req)
	if err != nil {
		c <- false
		return
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c <- false
		return
	}

	if resp.StatusCode != 200 {
		c <- false
		return
	}
	(*sm).mux.Lock()
	(*sm).m[apiName] = respBytes
	(*sm).mux.Unlock()
	c <- true
}

func unzip(proxy []byte, fs *billy.Filesystem, stem string) error {
	file, err := ioutil.TempFile("", "zip")
	if err != nil {
		return err
	}
	defer os.Remove(file.Name())
	err = ioutil.WriteFile(file.Name(), proxy, 0755)
	if err != nil {
		return err
	}
	r, err := zip.OpenReader(file.Name())
	if err != nil {
		return err
	}

	defer func() {
		if err := r.Close(); err != nil {
			panic(err)
		}
	}()

	(*fs).MkdirAll(stem, 0755)

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer func() {
			if err := rc.Close(); err != nil {
				panic(err)
			}
		}()

		path := filepath.Join(stem, f.Name)
		if f.FileInfo().IsDir() {
			(*fs).MkdirAll(path, f.Mode())
		} else {
			(*fs).MkdirAll(filepath.Dir(path), f.Mode())
			f, err := (*fs).OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer func() {
				if err := f.Close(); err != nil {
					panic(err)
				}
			}()
			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, f := range r.Reader.File {
		err := extractAndWriteFile(f)
		if err != nil {
			return err
		}
	}
	return nil
}

func writeOpenAPISpec(spec, filename string, fs *billy.Filesystem) error {
	f, err := (*fs).OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err = f.Write([]byte(spec)); err != nil {
		return err
	}
	return nil
}

func logChange(change, filename string, fs *billy.Filesystem) error {
	f, err := (*fs).OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}

	defer f.Close()
	change = fmt.Sprintf("%s:\t%s\n", time.Now().Format("2006-01-02 15:04:05"), change)
	if _, err = f.Write([]byte(change)); err != nil {
		return err
	}
	return nil
}
