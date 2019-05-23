package main

import (
	"apiinfo"
	"apiproxy"
	"apitesting"
	"apitransport"
	"github.com/gorilla/mux"
	"github.com/apex/log"
	"github.com/apex/log/handlers/json"
	"net/http"
	"os"
	"github"
	"tenant"
	"devportal"
)

func main() {

	if os.Getenv("UNAUTHORIZEDID_SECRET") == "" {
		panic("No Secret for unauthorized client id in environement. Please specif a secret on the environment variable UNAUTHORIZEDID_SECRET.")
	}
	if os.Getenv("SCPI_AUTH") == "" {
		panic("No credentials for SCPI in environemnt. Please specify credentials on the environment variable SCPI_AUTH.")
	}
	if os.Getenv("GITHUB_TOKEN") == "" {
		panic("No Github token in environment. Please specify a token on the environment variable GITHUB_TOKEN.")
	}
	if os.Getenv("DEVPORTAL_SECRET") == "" {
		panic("No Github token in environment. Please specify a token on the environment variable GITHUB_TOKEN.")
	}

	logFile, err := os.OpenFile("/logging/logs.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer logFile.Close()
	log.SetHandler(json.New(logFile))

	syncIn := make(chan github.Sync)
	syncOut := make(chan error)
	locks := tenant.InitializeTenantLocks()

	go github.StartGithubHandler(syncIn,syncOut)
	go github.GithubTenantSync(&locks, syncIn, syncOut)

	devportalIn := make(chan apiproxy.APIProxy)
	go devportal.Handler(devportalIn)

	r := mux.NewRouter()
	r.HandleFunc("/api/test/{tenant}/{name}", apitesting.Handler).Methods("POST")
	r.HandleFunc("/api/transport/{tenant}/{name}", apitransport.CreateTransportHandler(&locks, syncIn, syncOut, devportalIn)).Methods("POST")
	r.HandleFunc("/api/{tenant}/Management.svc/{target}", apiinfo.Handler).Methods("GET")

	//Uncomment for docker builds
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("/go/bin/dist/api-manager/")))
	//Uncomment for local builds
	//r.PathPrefix("/").Handler(http.FileServer(http.Dir("./dist/api-manager/")))
	http.ListenAndServe(":80", r)
}
