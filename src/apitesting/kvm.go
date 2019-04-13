package apitesting

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"tenant"
)

func KVMAuthorizationTest(c chan TestResult, tenantName, kvmName, clientID, auth, name string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://produs2apiportalapimgmtpphx-%s.us2.hana.ondemand.com/apiportal/api/1.0/Management.svc/KeyMapEntries('%s')/keyMapEntryValues", tenant.Get(tenantName), kvmName), nil)
	if err != nil {
		c <- TestResult{name, false, err}
		return
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Basic "+auth)
	resp, err := client.Do(req)

	if err != nil {
		c <- TestResult{name, false, err}
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		c <- TestResult{name, false, errors.New("Returned non 200 response")}
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))

	if err != nil {
		c <- TestResult{name, false, err}
		return
	}

	res := strings.Contains(string(body), fmt.Sprintf("name='%s'", clientID))
	c <- TestResult{name, res, nil}
}
