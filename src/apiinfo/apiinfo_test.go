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

func TestGetAPIInfo(t *testing.T) {
	auth := os.Getenv("SCPI_AUTH")
	APIProxies, err := GetAPIInfo("dev", "API_NIKE_CONVERSE_GET_ORDER_STATUS", auth)
	t.Log(APIProxies.APIs[0])
	if err != nil {
		t.Errorf("returned an error: %s", err)
	}
	if APIProxies.APIs[0].Url != "/Converse/OrderStatus" {
		t.Error("returned incorrect url")
	}
	if APIProxies.APIs[0].Name != "API_NIKE_CONVERSE_GET_ORDER_STATUS" {
		t.Errorf("returned incorrect Name, got %s", APIProxies.APIs[0].Name)
	}
}

func TestGetAllAPIInfo(t *testing.T) {
	auth := os.Getenv("SCPI_AUTH")
	APIProxies, err := GetAllAPIInfo("dev", auth)
	if err != nil {
		t.Errorf("returned an error: %s", err)
	}
	if len(APIProxies.APIs) < 2 {
		t.Error("not enough proxies, something went wrong")
	}
	for _, api := range(APIProxies.APIs) {
		t.Logf("%s: %s %s", api.Name, api.Url, api.Description)
	}
}

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

func TestGetAllAPIZip(t *testing.T) {
	auth := os.Getenv("SCPI_AUTH")
	if auth == "" {
		t.Error("No SCPI_AUTH in environment")
	}
	m, err := GetAllAPIZip("sandbox", auth)

	if err != nil {
		t.Errorf("returned an error: %s", err)
	}

	if len(m) < 1 {
		t.Errorf("not enough proxies returned")
	}
}
