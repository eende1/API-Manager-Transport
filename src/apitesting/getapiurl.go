package apitesting

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"tenant"
	"encoding/xml"
	"net/url"
)

type APIProxy struct {
	Url string `xml:"base_path"`
}

type APIProxies struct {
	APIs []APIProxy `xml:"entry>content>properties"`
}

func GetAPIURL(tenantName, apiName, auth string) (string, error){
	client := &http.Client{}
	query := url.PathEscape(fmt.Sprintf("$filter=FK_API_NAME eq '%s'", apiName))
	req, err := http.NewRequest("GET", fmt.Sprintf("https://produs2apiportalapimgmtpphx-%s.us2.hana.ondemand.com/apiportal/api/1.0/Management.svc/APIProxyEndPoints?%s", tenant.Get(tenantName), query), nil)
	if err != nil {
		return "", err
	}
	
	req.Header.Add("Authorization", "Basic "+auth)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != 200 {
		return "", errors.New("returned non 200 response")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	a := APIProxies{}
	err = xml.Unmarshal(body, &a)
	if err != nil {
		return "", err
	}
	if (len(a.APIs) == 0) {
		return "", errors.New("Could not find API Name")
	}
	
	return a.APIs[0].Url, nil
}

