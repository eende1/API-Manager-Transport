package apitransport

import (
	"os"
	"testing"
)

func TestGetAPIProxy(t *testing.T) {
	auth := os.Getenv("SCPI_AUTH")

	_, err := GetAPIProxy("dev", "TRACE_JENKINS_TEST", auth)
	if err != nil {
		t.Errorf("returned an error: %s", err)
	}
	//t.Log(string(res))
}

func TestTransport(t *testing.T) {
	auth := os.Getenv("SCPI_AUTH")
	_, err := Transport("dev", "TRACE_JENKINS_TEST", auth)
	if err != nil {
		t.Errorf("returned an error: %s", err)
	}
}
