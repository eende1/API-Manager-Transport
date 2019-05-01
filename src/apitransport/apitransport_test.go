package apitransport

import (
	"os"
	"testing"
	"github"
	"tenant"
)

func TestGetAPIProxy(t *testing.T) {
	auth := os.Getenv("SCPI_AUTH")
	if auth == "" {
		t.Errorf("No SCPI_AUTH in environment")
	}

	res, err := GetAPIProxy("dev", "TRACE_JENKINS_TEST", auth)
	if err != nil {
		t.Errorf("returned an error: %s", err)
	}
	t.Log(string(res))
}

func TestTransport(t *testing.T) {
	auth := os.Getenv("SCPI_AUTH")
	if auth == "" {
		t.Errorf("No SCPI_AUTH in environment")
	}
	syncIn := make(chan github.Sync)
	syncOut := make(chan error)

	locks := tenant.InitializeTenantLocks()
	
	go func(syncIn chan github.Sync, syncOut chan error) {
		_ = <-syncIn
		syncOut <- nil
	}(syncIn, syncOut)

	_, err := Transport("dev", "TRACE_JENKINS_TEST", "gotestcid", auth, &locks, syncIn, syncOut)
	if err != nil {
		t.Errorf("returned an error: %s", err)
	}
}

func TestTransportWithUpdate(t *testing.T) {
	auth := os.Getenv("SCPI_AUTH")
	if auth == "" {
		t.Errorf("No SCPI_AUTH in environment")
	}
	syncIn := make(chan github.Sync)
	syncOut := make(chan error)

	locks := tenant.InitializeTenantLocks()

	go github.StartGithubHandler(syncIn,syncOut)
	go github.GithubTenantSync(&locks, syncIn, syncOut)
	_, err := Transport("dev", "TRACE_JENKINS_TEST", "gotestcid", auth, &locks, syncIn, syncOut)
	if err != nil {
		t.Errorf("returned an error: %s", err)
	}
}

