package apitransport

import (
	"os"
	"testing"
	"github"
)

func TestGetAPIProxy(t *testing.T) {
	auth := os.Getenv("SCPI_AUTH")

	res, err := GetAPIProxy("dev", "TRACE_JENKINS_TEST", auth)
	if err != nil {
		t.Errorf("returned an error: %s", err)
	}
	t.Log(string(res))
}

func TestTransport(t *testing.T) {
	auth := os.Getenv("SCPI_AUTH")
	syncIn := make(chan github.Sync)
	syncOut := make(chan error)
	
	go func(syncIn chan github.Sync, syncOut chan error) {
		_ = <-syncIn
		syncOut <- nil
	}(syncIn, syncOut)

	_, err := Transport("dev", "TRACE_JENKINS_TEST", "gotestcid", auth, syncIn, syncOut)
	if err != nil {
		t.Errorf("returned an error: %s", err)
	}
}

func TestTransportWithUpdate(t *testing.T) {
	auth := os.Getenv("SCPI_AUTH")
	syncIn := make(chan github.Sync)
	syncOut := make(chan error)

	go func(syncIn chan github.Sync, syncOut chan error) {
		apimRepo, err := github.InitializeGithubRepo();
		//scpiAuth := os.Getenv("SCPI_AUTH")
		if err != nil {
			panic("Failed to initialize github repository.")
		}

		toSync := <- syncIn
		apimRepo.SyncAPIs(toSync.Proxies, toSync.TenantName, toSync.LogMessage)
		syncOut <- nil
	}(syncIn, syncOut)

	_, err := Transport("dev", "TRACE_JENKINS_TEST", "gotestcid", auth, syncIn, syncOut)
	if err != nil {
		t.Errorf("returned an error: %s", err)
	}
}

