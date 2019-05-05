package apitesting

import (
	"os"
	"testing"
)

func TestUnauthorizedClientTest(t *testing.T) {
	c := make(chan TestResult)
	go UnauthorizedClientTest(c, "https://postman-echo.com/get", "GET", "unathorized client test")
	result := <-c
	if result.Err != nil {
		t.Errorf("returned an error: %s", result.Err)
	}
	if result.Pass {
		t.Errorf("should have received false")
	}

	go UnauthorizedClientTest(c, "https://nikescpdev.apimanagement.us2.hana.ondemand.com/DeliveryDetails/$metadata", "GET", "unathorized client test")
	result = <-c
	if result.Err != nil {
		t.Errorf("returned an error: %s", result.Err)
	}
	if !result.Pass {
		t.Errorf("should have received false")
	}
}

func TestKVMAuthorizationTest(t *testing.T) {
	auth := os.Getenv("SCPI_AUTH")
	c := make(chan TestResult)
	go KVMAuthorizationTest(c, "dev", "API_NIKE_PRA_GET_SHIPMENT_DETAILS", "nike.sapapi.ntntest", auth, "kvm test")
	result := <-c
	if result.Err != nil {
		t.Errorf("returned an error: %s", result.Err)
	}

	if !result.Pass {
		t.Errorf("Got false, should have gotten true")
	}

	go KVMAuthorizationTest(c, "dev", "API_NIKE_PRA_GET_SHIPMENT_DETAILS", "should fail", auth, "kvm test")
	result = <-c
	if result.Err != nil {
		t.Errorf("returned an error: %s", result.Err)
	}

	if result.Pass {
		t.Errorf("Got true, should have gotten false")
	}

	go KVMAuthorizationTest(c, "dev", "doesntexist", "should fail", auth, "kvm test")
	result = <-c
	if result.Err != nil {
		t.Errorf("returned an error: %s", result.Err)
	}

	if result.Pass {
		t.Errorf("Got true, should have gotten false")
	}
}

func TestAPICall(t *testing.T) {
	resp, err := APICall("http://postman-echo.com/get", "no auth", "GET")
	if err != nil {
		t.Errorf("returned an error: %s", err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("returned non 200 response code")
	}
}

func TestAPICallTest(t *testing.T) {
	c := make(chan TestResult)
	go APICallTest(c, "http://postman-echo.com/get", "notokenneeded", "GET",  "API Call Test")
	result := <-c
	if result.Err != nil {
		t.Errorf("returned an error: %s", result.Err)
	}

	if !result.Pass {
		t.Error("test should have passed, but it failed")
	}
}
