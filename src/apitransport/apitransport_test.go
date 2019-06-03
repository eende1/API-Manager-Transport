package apitransport

import (
	"os"
	"testing"
	"github"
	"tenant"
	"apiproxy"
	"apitesting"
)

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

	a := apiproxy.APIProxy{}
	a.Name = "TRACE_JENKINS_TEST"
	a.Tenant = "dev"
	a.Auth = auth

	test := apitesting.APITest{}
	test.APIProxy = a
	test.TokenClientID = "gotestcid"
	_, err := Transport(test, &locks, syncIn, syncOut)
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

		a := apiproxy.APIProxy{}
	a.Name = "TRACE_JENKINS_TEST"
	a.Tenant = "dev"
	a.Auth = auth

	test := apitesting.APITest{}
	test.APIProxy = a
	test.TokenClientID = "gotestcid"
	_, err := Transport(test, &locks, syncIn, syncOut)
	if err != nil {
		t.Errorf("returned an error: %s", err)
	}
}
