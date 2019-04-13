package apiinfo

import (
	"os"
	"testing"
)

func TestAPIInfoCall(t *testing.T) {
	auth := os.Getenv("SCPI_AUTH")

	_, err := APIInfoCall("dev", "APIProxies", auth)
	if err != nil {
		t.Errorf("returned an error: %s", err)
	}
}
