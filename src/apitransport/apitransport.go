// Package apitransport provides functions for transport API Proxies.
package apitransport

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"tenant"
	"github"
	"encoding/json"
	"github.com/apex/log"
	"apitesting"
	"apiproxy"
	"strings"
	"io/ioutil"
)

const conversionIflowURL = "https://l5347-iflmap.hcisbp.us2.hana.ondemand.com/http/metadatatoswagger"

func CreateTransportHandler(tenantLock *tenant.Lock, syncIn chan github.Sync, syncOut chan error, devportalIn chan apiproxy.APIProxy) (func (w http.ResponseWriter, r *http.Request)) {
return func (w http.ResponseWriter, r *http.Request) {
	apiTest, err := apitesting.ParseApiTest(r)
		ctx := log.WithFields(log.Fields{
		"path": apiTest.Path,
		"method": apiTest.Method,
		"email": apiTest.Email,
		"api": apiTest.APIProxy.Name,
		"tenant": apiTest.APIProxy.Tenant,
		"url": apiTest.APIProxy.Url,
	})
	defer ctx.Trace("Transport").Stop(&err)
	if err != nil {
		http.Error(w, fmt.Sprintf("{'status':'%s'}", err), 400)
		return
	}

	numTests := apitesting.NumTests + 2
	testResults := make(chan apitesting.TestResult, numTests)

	apiTest.ExecuteTests(testResults)

	// Assume highest tenant, degrade if otherwise
	kvmAuthTenant := "prod"
	if apiTest.APIProxy.Tenant == "dev" {
		kvmAuthTenant = "qa"
	}

	go apitesting.KVMAuthorizationTest(testResults, kvmAuthTenant, apiTest.APIProxy.Name, apiTest.TokenClientID,
		os.Getenv("SCPI_AUTH"), "authorized in target tenant")
	go apitesting.LDAPAuthenticationTest(testResults, apiTest.Email, apiTest.Password, "ldap authentication test")

	results := apitesting.Responses{}
	transport := true
	for i := 0; i < numTests; i++ {
		results = append(results, <-testResults)
		transport = transport && results[len(results)-1].Pass
	}
	close(testResults)

	if transport {
		t := apitesting.TestResult{"transport", false, nil}
		resp, err := Transport(apiTest, tenantLock, syncIn, syncOut)

		if err == nil && resp {
			t.Pass = true
			devportalIn <- apiTest.APIProxy
		}
		results = append(results, t)
	}

	resultJson, err := json.Marshal(results)
	if err != nil {
		http.Error(w, fmt.Sprintf("{'error':'%s'}", err.Error()), 500)
		return
	}
	w.Write([]byte(resultJson))
}
}

// Tranports an API Proxy from one tenant to the next while holding locks, then Sync with Github.
func Transport(apiTest apitesting.APITest, tenantLock *tenant.Lock, syncIn chan github.Sync, syncOut chan error) (bool, error) {
	lock, ok := (*tenantLock).Map[apiTest.APIProxy.Tenant]
	if !ok {
		return false, errors.New("Failed to get lock for this tenant in transport")
	}
	(*lock).Lock()
	defer (*lock).Unlock()
	c := make(chan error)
	go apiTest.APIProxy.GetZip(c)
	err := <-c
	if err != nil {
		return false, err
	}

	err = apiTest.APIProxy.Transport()
	if err != nil {
		return false, err
	}

	OpenAPISpec := ""
	if apiTest.APIProxy.OData {
		OpenAPISpec, err = convertMetaDatatoOpenAPI(apiTest.Result)
		if err != nil {
			fmt.Println("failed to convert metadata")
		}
	}

	proxies := apiproxy.APIProxies{}
	proxies.APIs = append(proxies.APIs, apiTest.APIProxy)
	syncIn <- github.Sync{proxies, fmt.Sprintf("API Transported by %s", apiTest.Email), OpenAPISpec}
	<- syncOut

 	return true, nil
}

func convertMetaDatatoOpenAPI(metaData string) (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", conversionIflowURL, strings.NewReader(metaData))
	if err != nil {
		return "", err
	}
	req.Header.Add("Authorization", "Basic "+os.Getenv("SCPI_AUTH"))
	resp, err := client.Do(req)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Iflow failed to convert metadata to swagger: %s", err))
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(bodyBytes), nil
}
