package apiinfo

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"os"
	"tenant"
	"encoding/xml"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	auth := os.Getenv("SCPI_AUTH")
	vars := mux.Vars(r)

	tenantName := "dev"
	if vars["tenant"] == "prodapi" {
		tenantName = "prod"
	}

	if vars["tenant"] == "qaapi" {
		tenantName = "qa"
	}

	resp, err := APIInfoCall(tenantName, vars["target"], auth)
	if err != nil {
		http.Error(w, fmt.Sprintf("{'status':%s}", err), 400)
	}
	w.Write(resp)
}

func APIInfoCall(tenantName, target, auth string) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://produs2apiportalapimgmtpphx-%s.us2.hana.ondemand.com/apiportal/api/1.0/Management.svc/%s", tenant.Get(tenantName), target), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Basic "+auth)
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("returned non 200 response")
	}
	return respBytes, nil
}

type APIProxy struct {
	Name string `xml:"name"`
}

type APIProxies struct {
	APIs []APIProxy `xml:"entry>content>properties"`
}

func GetAllAPINames(tenantName, auth string) (APIProxies, error) {
	var a APIProxies
	client := &http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://produs2apiportalapimgmtpphx-%s.us2.hana.ondemand.com/apiportal/api/1.0/Management.svc/APIProxies", tenant.Get(tenantName)), nil)
	if err != nil {
		return a, err
	}

	req.Header.Add("Authorization", "Basic "+auth)
	resp, err := client.Do(req)
	if err != nil {
		return a, err
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return a, err
	}

	if resp.StatusCode != 200 {
		return a, errors.New("returned non 200 response")
	}

	err = xml.Unmarshal(respBytes, &a)
	if err != nil {
		return a, err
	}

	return a, nil
}
