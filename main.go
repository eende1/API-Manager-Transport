package main

import (
	"apiinfo"
	"apitesting"
	"apitransport"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	
	if os.Getenv("UNAUTHORIZEDID_SECRET") == "" {
		panic("No Secret for unauthorized client id in environement")
	}
	if os.Getenv("SCPI_AUTH") == "" {
		panic("No credentials for SCPI in environemnt")
	}
	r := mux.NewRouter()
	r.HandleFunc("/api/test", apitesting.Handler).Methods("POST")
	r.HandleFunc("/api/transport", apitransport.Handler).Methods("POST")
	r.HandleFunc("/{tenant}/Management.svc/{target}", apiinfo.Handler).Methods("GET")
	
	// Uncomment for docker builds
	//r.PathPrefix("/").Handler(http.FileServer(http.Dir("/go/bin/dist/api-manager/")))
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./dist/api-manager/")))
	http.ListenAndServe(":80", r)
}
