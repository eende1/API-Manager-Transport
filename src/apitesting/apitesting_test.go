package apitesting

import (
	"os"
	"testing"
	"apiproxy"
)

func TestUnauthorizedClientTest(t *testing.T) {
	c := make(chan TestResult)
	a := apiproxy.APIProxy{}
	a.Url = "https://postman-echo.com/get"

	test := APITest{}
	test.APIProxy = a
	test.Method = "GET"

	go (&test).UnauthorizedClientTest(c, "unauthorized client test")

	result := <-c
	if result.Err != nil {
		t.Errorf("returned an error: %s", result.Err)
	}
	if result.Pass {
		t.Errorf("should have received false")
	}

	b := apiproxy.APIProxy{}
	b.Url = "https://nikescpdev.apimanagement.us2.hana.ondemand.com/DeliveryDetails/$metadata"

	test2 := APITest{}
	test2.APIProxy = b
	test2.Method = "GET"
	go (&test2).UnauthorizedClientTest(c, "unathorized client test")

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

func TestAPICallTest(t *testing.T) {
	c := make(chan TestResult)
	a := apiproxy.APIProxy{}
	a.Url = "https://postman-echo.com/get"

	test := APITest{}
	test.APIProxy = a
	test.Method = "GET"
	go (&test).APICallTest(c, "API Call Test")
	result := <-c
	if result.Err != nil {
		t.Errorf("returned an error: %s", result.Err)
	}

	if !result.Pass {
		t.Error("test should have passed, but it failed")
	}
}
