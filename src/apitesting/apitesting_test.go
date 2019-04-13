package apitesting

import (
	"os"
	"testing"
)

func TestAPICall(t *testing.T) {
	c := make(chan TestResult)
	go CallAPI(c, "https://postman-echo.com/get", "", "", "api call test")
	result := <-c
	if result.Err != nil {
		t.Errorf("returned an error: %s", result.Err)
	}
	if !result.Pass {
		t.Errorf("response code was not 200")
	}
}

func TestGenerateToken(t *testing.T) {
	token, err := GenerateToken("nike.sapae.unauthorizedid",
		"n4YFiCwufDiUzUO7tQVjcccsU3nmPt9W5aiVEFWGgskFVcSJ9v9XN98eqCE3dOOW")
	t.Log(token)
	if err != nil {
		t.Errorf("returned an error: %s", err)
	}
}

func TestUnauthorizedClientTest(t *testing.T) {
	c := make(chan TestResult)
	go UnauthorizedClientTest(c, "https://postman-echo.com/get", "", "unathorized client test")
	result := <-c
	if result.Err != nil {
		t.Errorf("returned an error: %s", result.Err)
	}
	if result.Pass {
		t.Errorf("should have received false")
	}

	go UnauthorizedClientTest(c, "https://nikescpdev.apimanagement.us2.hana.ondemand.com/DeliveryDetails", "/$metadata", "unathorized client test")
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
