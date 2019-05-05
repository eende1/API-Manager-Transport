package apiproxy

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"tenant"
	"encoding/xml"
	"bytes"
	"encoding/base64"
)

const transportURL = "https://produs2apiportalapimgmtpphx-%s.us2.hana.ondemand.com/apiportal/api/1.0/Transport.svc/APIProxies"

type APIProxy struct {
	OData bool
	Tenant string
	Name string `xml:"content>properties>name"`
	Url string `xml:"link>inline>feed>entry>content>properties>base_path"`
	Provider string `xml:"content>properties>FK_PROVIDERNAME"`
	Description string `xml:"content>properties>description"`
	Zip []byte
	Auth string
}

type APIProxies struct {
	APIs []APIProxy `xml:"entry"`
}

func Get(tenantName, name, auth string) (APIProxy, error) {
	a := APIProxy{}
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://produs2apiportalapimgmtpphx-%s.us2.hana.ondemand.com/apiportal/api/1.0/Management.svc/APIProxies('%s')?$expand=proxyEndPoints", tenant.Get(tenantName), name), nil)
	if err != nil {
		return a, err
	}
	req.Header.Add("Authorization", "Basic "+auth)
	resp, err := client.Do(req)
	if err != nil {
		return a, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return a, errors.New(fmt.Sprintf("returned non 200 response: %s", resp.Status))
	}

	if err != nil {
		return a, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return a, err
	}

	err = xml.Unmarshal(body, &a)
	if err != nil {
		return a, err
	}

	if (a.Url == "") {
		return a, errors.New(fmt.Sprintf("Failed to get API Proxies"))
	}

	a.Tenant = tenantName
	a.Auth = auth
	a.Url = fmt.Sprintf("https://nikescp%s.apimanagement.us2.hana.ondemand.com%s", a.Tenant, a.Url)
	a.OData = true
	return a, nil
}

func GetAll(tenantName, auth string) (APIProxies, error) {
	a := APIProxies{}
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://produs2apiportalapimgmtpphx-%s.us2.hana.ondemand.com/apiportal/api/1.0/Management.svc/APIProxies?$expand=proxyEndPoints", tenant.Get(tenantName)), nil)
	if err != nil {
		return a, err
	}
	req.Header.Add("Authorization", "Basic "+auth)
	resp, err := client.Do(req)
	if err != nil {
		return a, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return a, errors.New(fmt.Sprintf("returned non 200 response: %s", resp.Status))
	}

	if err != nil {
		return a, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return a, err
	}

	err = xml.Unmarshal(body, &a)
	if err != nil {
		return a, err
	}
	if (len(a.APIs) == 0) {
		return a, errors.New(fmt.Sprintf("Failed to get API Proxies"))
	}

	for i := 0; i < len(a.APIs); i++ {
		a.APIs[i].Tenant = tenantName
		a.APIs[i].Auth = auth
		a.APIs[i].Url = fmt.Sprintf("https://nikescp%s.apimanagement.us2.hana.ondemand.com%s", a.APIs[i].Tenant, a.APIs[i].Url)
		a.APIs[i].OData = true
	}

	return a, nil
}

func (p *APIProxies) PopulateZips() error {
	c := make(chan error, len(p.APIs))
	for i := 0; i < len(p.APIs); i++ {
		go p.APIs[i].GetZip(c)
	}

	var encounteredErrors []error
	for i := 0; i < len(p.APIs); i++ {
		encounteredErrors = append(encounteredErrors, <-c)
	}
	close(c)

	for i := 0; i < len(encounteredErrors); i++ {
		if encounteredErrors[i] != nil {
			return errors.New(fmt.Sprintf("Encountered error getting an Zip: %s", encounteredErrors[i]))
		}
	}

	return nil
}

func (api *APIProxy) GetZip(c chan error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://produs2apiportalapimgmtpphx-%s.us2.hana.ondemand.com/apiportal/api/1.0/Transport.svc/APIProxies?name=%s", tenant.Get(api.Tenant), api.Name), nil)
	if err != nil {
		c <- err
	}
	req.Header.Add("Authorization", "Basic "+api.Auth)
	resp, err := client.Do(req)

	if err != nil {
		c <- err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		c <- errors.New(fmt.Sprintf("Returned non 200 response getting APIProxy for %s in %s: %s", api.Name, api.Tenant, resp.Status))
	}

	api.Zip, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		c <- err
	}
	c <- nil
}

// Tranports an API Proxy from one tenant to the next.
func (api *APIProxy) Transport() error {
	//Advance tenant by one stage
	if api.Tenant == "dev" {
		api.Tenant = "qa"
	} else if api.Tenant == "qa" {
		api.Tenant = "prod"
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf(transportURL, tenant.Get(api.Tenant)), nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "Basic "+api.Auth)
	req.Header.Add("x-csrf-token", "fetch")
	resp, err := client.Do(req)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	str := base64.StdEncoding.EncodeToString(api.Zip)
	req, err = http.NewRequest("POST", fmt.Sprintf(transportURL, tenant.Get(api.Tenant)), bytes.NewBuffer([]byte(str)))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/octet-stream")
	req.Header.Add("Authorization", "Basic "+api.Auth)
	req.Header.Add("x-csrf-token", resp.Header["X-Csrf-Token"][0])
	cookies := resp.Cookies()

	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}

	resp, err = client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New((fmt.Sprintf("returned non 200 response: %d", resp.StatusCode)));
	}

	return nil
}
