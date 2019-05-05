package apiproxy

import (
	"os"
	"testing"
)

func TestGet(t *testing.T) {
	auth := os.Getenv("SCPI_AUTH")
	api, err := Get("dev","API_NIKE_CONVERSE_GET_ORDER_STATUS", auth)
	t.Log(api)
	if err != nil {
		t.Errorf("returned an error: %s", err)
	}
	if api.Url != "https://nikescpdev.apimanagement.us2.hana.ondemand.com/Converse/OrderStatus" {
		t.Error("returned incorrect url")
	}
	if api.Name != "API_NIKE_CONVERSE_GET_ORDER_STATUS" {
		t.Errorf("returned incorrect Name, got %s", api.Name)
	}
	if api.Provider != "PO_HTTPS" {
				t.Errorf("returned incorrect Provider, got %s", api.Provider)
	}
	if api.Description != "<p>This service gets sales order status from Converse system.</p>" {
		t.Errorf("returned incorrect description, got %s", api.Description)
	}
}

func TestGetAll(t *testing.T) {
	auth := os.Getenv("SCPI_AUTH")
	api, err := GetAll("dev", auth)
	t.Log(api)
	if err != nil {
		t.Errorf("returned an error: %s", err)
	}
}

func TestGetZip(t *testing.T) {
	auth := os.Getenv("SCPI_AUTH")
	api, err := Get("dev","API_NIKE_CONVERSE_GET_ORDER_STATUS", auth)
	t.Log(api)
	if err != nil {
		t.Errorf("returned an error getting api: %s", err)
	}
	c := make(chan error)
	go api.GetZip(c)
	err = <- c
	if err != nil {
		t.Errorf("returned an error getting zip: %s", err)
	}
}

func TestPopulateZips(t *testing.T) {
	auth := os.Getenv("SCPI_AUTH")
	proxies, err := GetAll("dev", auth)
	if err != nil {
		t.Errorf("returned an error getting proxies: %s", err)
	}
	err = proxies.PopulateZips()
	if err != nil {
		t.Errorf("returned an error getting zips: %s", err)
	}
}

/*
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
*/
