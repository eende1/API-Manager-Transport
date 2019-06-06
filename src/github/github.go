// Package github provides functions for syncing Proxies to a github repository
package github

import (
	"archive/zip"
	"fmt"
	"gopkg.in/src-d/go-billy.v4"
	"gopkg.in/src-d/go-billy.v4/memfs"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
	"gopkg.in/src-d/go-git.v4/storage/memory"
	"github.com/apex/log"
	"io/ioutil"
	"os"
	"io"
	"path/filepath"
	"time"
	"tenant"
	"apiproxy"
	"errors"
)

const tenantSyncIntervalMinutes = 15

type Sync struct {
	Proxies apiproxy.APIProxies
	LogMessage string
	OpenAPISpec string
}

type Repo struct {
	fs billy.Filesystem
	r *git.Repository
	worktree *git.Worktree
}

func GithubTenantSync(tenantLock *tenant.Lock, syncIn chan Sync, syncOut chan error) {
	tenantName := "sandbox"
	for {
		time.Sleep(tenantSyncIntervalMinutes * time.Minute)
		// Acquire lock on tenant
		lock, ok := (*tenantLock).Map[tenantName]
		ctx := log.WithFields(log.Fields{
			"tenant": tenantName,
		})
		if !ok {
			ctx.Trace("Github Sync").Error("Github Sync: Failed to get lock")
			continue
		}
		(*lock).Lock()

		apiProxies, err := apiproxy.GetAll(tenantName, os.Getenv("SCPI_AUTH"))
		if err != nil {
			(*lock).Unlock()
			tenantName = tenant.Advance(tenantName)
			ctx.Trace("Github Sync: Failed to get proxies").Stop(&err)
			continue
		}

		err = apiProxies.PopulateZips()
		if err != nil {
			(*lock).Unlock()
			tenantName = tenant.Advance(tenantName)
			ctx.Trace("Github Sync: Failed to get zips").Stop(&err)
			continue
		}

		syncIn <- Sync{apiProxies, "", ""}
		err = <-syncOut
		(*lock).Unlock()


		if err != nil {
			ctx.WithField("API Count", len(apiProxies.APIs)).Trace("Github Sync: Failed to Sync APIs").Stop(&err)
		} else {
			ctx.WithField("API Count", len(apiProxies.APIs)).Info("Github Sync")
		}

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
		err = apimRepo.SyncAPIs(toSync.Proxies, toSync.LogMessage, toSync.OpenAPISpec)
		syncOut <- err
		if err != nil {
			log.WithField("error", err.Error()).Error("Failed to synch Proxies")
		}
	}
}

func (g *Repo) SyncAPIs(proxies apiproxy.APIProxies, logMessage string, openAPISpec string) error {
	for _, proxy := range proxies.APIs {
		err := unzipInGitRepository(proxy.Zip, &(g.fs), fmt.Sprintf("%s/%s", proxy.Tenant, proxy.Name))
		if err != nil {
			//log.Printf("Failed to unzip in SyncAPIS\n")
			return errors.New("Failed to unzip proxy")
		}

		err = logChange(logMessage, fmt.Sprintf("%s/%s/%s", proxy.Tenant, proxy.Name, "change_log.txt"), &(g.fs))
		if err != nil {
			//log.Printf("Failed to log change in SyncAPIS\n")
			return errors.New("Failed to log change")
		}

		err = writeOpenAPISpec(openAPISpec, fmt.Sprintf("%s/%s/%s", proxy.Tenant, proxy.Name, "spec.json"), &(g.fs))
		if err != nil {
			//log.Printf("Failed to write openAPISpec SyncAPIS\n")
			return errors.New("Failed to write open API Spec")
		}
	}
	g.commitRepo("commit by sync process")
	return nil
}

func (g *Repo) commitRepo(commitMessage string) {
	g.worktree.AddGlob("*")

	_, err := g.worktree.Commit(commitMessage, &git.CommitOptions{
		Author: &object.Signature{
			Name: "Backup Process",
			When: time.Now(),
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	err = g.r.Push(&git.PushOptions{
		Auth: &http.BasicAuth{
			Username: "tandr9",
			Password: os.Getenv("GITHUB_TOKEN"),
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
}

func initializeGithubRepo() (Repo, error){
	fs := memfs.New()
	storer := memory.NewStorage()

	g, err := git.Clone(storer, fs, &git.CloneOptions{
		URL: "https://api.github.nike.com/scp/APIM-Backup",
		Auth: &http.BasicAuth{
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

func unzipInGitRepository(proxy []byte, fs *billy.Filesystem, stem string) error {
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
	if spec == "" {
		return nil
	}
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
	if change == "" {
		return nil
	}

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
