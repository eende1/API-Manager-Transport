package github

import (
	"archive/zip"
	"encoding/xml"
	"errors"
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
	"strings"
	"time"
)

type Sync struct {
	Proxies map[string][]byte
	TenantName string
	LogMessage string
}

var tenantMap map[string]string = map[string]string{"dev": "dfc3ccb1f", "qa": "d8b3bfb89", "prod": "d9cfb42fa", "sandbox": "d6c83d68e"}

func tenantGet(tenant string) string {
	return tenantMap[strings.ToLower(tenant)]
}

type APIProxy struct {
	Name string `xml:"name"`
}

type APIProxies struct {
	APIs []APIProxy `xml:"entry>content>properties"`
}

func getAllAPINames(tenantName, auth string) (APIProxies, error) {
	var a APIProxies
	client := &http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://produs2apiportalapimgmtpphx-%s.us2.hana.ondemand.com/apiportal/api/1.0/Management.svc/APIProxies", tenantGet(tenantName)), nil)
	if err != nil {
		return a, err
	}

	req.Header.Add("Authorization", "Basic "+auth)
	resp, err := client.Do(req)
	if err != nil {
		return a, err
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return a, err
	}

	if resp.StatusCode != 200 {
		return a, errors.New("returned non 200 response")
	}

	err = xml.Unmarshal(respBytes, &a)
	if err != nil {
		return a, err
	}

	return a, nil
}

func GetAllAPIZip(tenantName, auth string) (map[string][]byte) {
	APIProxies, err := getAllAPINames(tenantName, auth)
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan bool, len(APIProxies.APIs))
	m := make(map[string][]byte)
	for _, a := range APIProxies.APIs {
		fmt.Println(a.Name)
		go GetAPIZip(done, tenantName, a.Name, auth, m)
	}

	for i := 0; i < len(APIProxies.APIs); i++ {
		<-done
	}
	return m
}
// UNEXPORT THIS AT SOME POINT
func GetAPIZip(c chan bool, tenantName, apiName, auth string, m map[string][]byte) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://produs2apiportalapimgmtpphx-%s.us2.hana.ondemand.com/apiportal/api/1.0/Transport.svc/APIProxies?name=%s", tenantGet(tenantName), apiName), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Basic "+auth)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("returned non 200 response")
	}

	m[apiName] = respBytes
	c <- true

	return respBytes, nil
}

type Repo struct {
	fs billy.Filesystem
	r *git.Repository
	worktree *git.Worktree
}

func InitializeGithubRepo() (Repo, error){
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

func (g *Repo) SyncAPIs(proxies map[string]([]byte), tenantName, logMessage string) {
	
	for name, proxy := range proxies {
		err := unzip(proxy, &(g.fs), fmt.Sprintf("%s/%s", tenantName, name))
		if err != nil {
			log.Fatal(err)
		}
		// Include a log message unless an empty message was passed.
		if logMessage != "" {
			err = logChange(logMessage, fmt.Sprintf("%s/%s/%s", tenantName, name, "change_log.txt"), &(g.fs))
			if err != nil {
				log.Fatal(err)
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

func unzip(proxy []byte, fs *billy.Filesystem, stem string) error {
	file, err := ioutil.TempFile("", "zip")
	if err != nil {
		log.Fatal(err)
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
			//fmt.Println(path)
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
