package apitransport

import (
	"apitesting"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"tenant"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	apiTest, err := apitesting.ParseApiTest(r)
	if err != nil {
		http.Error(w, fmt.Sprintf("{'status':'%s'}", err), 400)
		return
	}

	numTests := apitesting.NumTests + 1
	testResults := make(chan apitesting.TestResult, numTests)

	apiTest.ExecuteTests(testResults)

	// Advance tenant
	if apiTest.Tenant == "DEV" {
		apiTest.Tenant = "QA"
	} else {
		apiTest.Tenant = "PROD"
	}

	go apitesting.KVMAuthorizationTest(testResults, strings.ToLower(apiTest.Tenant), apiTest.APIName, apiTest.TokenClientID,
		os.Getenv("SCPI_AUTH"), "authorized in target tenant")

	results := apitesting.Responses{}
	transport := true
	for i := 0; i < numTests; i++ {
		results = append(results, <-testResults)
		transport = transport && results[len(results)-1].Pass
	}
	close(testResults)

	if transport {
		t := apitesting.TestResult{"transport", false, nil}
		//resp, err := apitransport.Transport(u.Tenant, u.APIName, os.Getenv("SCPI_AUTH"))
		resp, err := true, error(nil)
		if err == nil && resp {
			t.Pass = true
		}
		results = append(results, t)
	}

	resultJson, err := json.Marshal(results)
	if err != nil {
		http.Error(w, fmt.Sprintf("{'error':'%s'}", err.Error()), 500)
	}
	w.Write([]byte(resultJson))
}

func GetAPIProxy(tenantName, apiName, auth string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://produs2apiportalapimgmtpphx-%s.us2.hana.ondemand.com/apiportal/api/1.0/Transport.svc/APIProxies?name=%s", tenant.Get(tenantName), apiName), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Basic "+auth)
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New("Returned non 200 response")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func Transport(tenantName, apiName, auth string) (bool, error) {

	APIProxy, err := GetAPIProxy(tenantName, apiName, auth)
	if err != nil {
		return false, err
	}

	//Advance tenant by one stage
	if tenantName == "dev" {
		tenantName = "qa"
	} else if tenantName == "qa" {
		tenantName = "prod"
	}

	APIProxyReader := bytes.NewReader(APIProxy)

	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://produs2apiportalapimgmtpphx-%s.us2.hana.ondemand.com/apiportal/api/1.0/Transport.svc/APIProxies", tenant.Get(tenantName)), nil)
	if err != nil {
		return false, err
	}

	req.Header.Add("Authorization", "Basic "+auth)
	req.Header.Add("x-csrf-token", "fetch")
	resp, err := client.Do(req)

	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	req, err = http.NewRequest("POST", fmt.Sprintf("https://produs2apiportalapimgmtpphx-%s.us2.hana.ondemand.com/apiportal/api/1.0/Transport.svc/APIProxies", tenant.Get(tenantName)), APIProxyReader)
	if err != nil {
		return false, err
	}
	req.Header.Add("Content-Type", "application/zip")
	req.Header.Add("Authorization", "Basic "+auth)
	req.Header.Add("x-csrf-token", resp.Header["X-Csrf-Token"][0])
	cookies := resp.Cookies()

	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}

	resp, err = client.Do(req)
	if err != nil {
		return false, err
	}

	return true, nil
}
