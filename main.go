package main

import (
	"apiinfo"
	"apitesting"
	"apitransport"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"github"
	"fmt"
	"time"
	//"github.com/pkg/profile"
)

func main() {

	//defer profile.Start(profile.MemProfile).Stop()
	
	if os.Getenv("UNAUTHORIZEDID_SECRET") == "" {
		panic("No Secret for unauthorized client id in environement. Please specif a secret on the environment variable UNAUTHORIZEDID_SECRET.")
	}
	if os.Getenv("SCPI_AUTH") == "" {
		panic("No credentials for SCPI in environemnt. Please specify credentials on the environment variable SCPI_AUTH.")
	}
	if os.Getenv("GITHUB_TOKEN") == "" {
		panic("No Github token in environment. Please specify a token on the environment variable GITHUB_TOKEN.")
	}

	syncIn := make(chan github.Sync)
	syncOut := make(chan error)

	go func(syncIn chan github.Sync, syncOut chan error) {
		apimRepo, err := github.InitializeGithubRepo();
		//scpiAuth := os.Getenv("SCPI_AUTH")
		if err != nil {
			panic("Failed to initialize github repository.")
		}

		for {
			toSync := <- syncIn
			start := time.Now()
			
			apimRepo.SyncAPIs(toSync.Proxies, toSync.TenantName, toSync.LogMessage)
			syncOut <- nil
			elapsed := time.Since(start)
			fmt.Printf("sync took %s", elapsed)
		}
	}(syncIn, syncOut)
	
	go func(syncIn chan github.Sync, syncOut chan error) {
		tenant := "sandbox"
		for {
			fmt.Println(tenant)
			apiProxies := github.GetAllAPIZip(tenant, os.Getenv("SCPI_AUTH"))
			syncIn <- github.Sync{apiProxies, tenant, ""}
			<- syncOut
			tenant = advanceTenant(tenant)
			time.Sleep(20 * time.Second)
		}
	}(syncIn, syncOut)
	/*
	c := make(chan bool)
	m := make(map[string][]byte)
	go github.GetAPIZip(c, "sandbox", "API_NIKE_TMP_SYNC_TEST", os.Getenv("SCPI_AUTH"), m)
	<-c
	fmt.Println("go zip")
	syncIn <- github.Sync{m, "sandbox", "logtest"}
	<- syncOut
	*/
	r := mux.NewRouter()
	r.HandleFunc("/api/test", apitesting.Handler).Methods("POST")
	r.HandleFunc("/api/transport", apitransport.CreateTransportHandler(syncIn, syncOut)).Methods("POST")
	r.HandleFunc("/{tenant}/Management.svc/{target}", apiinfo.Handler).Methods("GET")
	
	//Uncomment for docker builds
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("/go/bin/dist/api-manager/")))
	//Uncomment for local builds
	//r.PathPrefix("/").Handler(http.FileServer(http.Dir("./dist/api-manager/")))
	http.ListenAndServe(":80", r)
}

func advanceTenant(tenant string) string {
	switch tenant {
	case "sandbox":
		return "dev"
	case "dev":
		return "qa"
	case "qa":
		return "sandbox"
	/*	
	case "prod":
		return "sandbox"
        */
	}
	return ""
}
