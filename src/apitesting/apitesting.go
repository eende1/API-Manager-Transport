package apitesting

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"io/ioutil"
	"github.com/gorilla/mux"
	"okta"
	"apiproxy"
)

const NumTests = 3

type APITest struct {
	APIProxy apiproxy.APIProxy
	Token   string
	Payload string
	Path  string
	Method string

	Email string
	Password string
	TokenClientID string `json:"cid"`
	Result string
}

type TestResult struct {
	Name string
	Pass bool
	Err  error
}

type Responses []TestResult

func Handler(w http.ResponseWriter, r *http.Request) {
	apiTest, err := ParseApiTest(r)
	if err != nil {
		http.Error(w, fmt.Sprintf("{'error':'%s'}", err), 400)
		return
	}
	testResults := make(chan TestResult, NumTests)
	apiTest.ExecuteTests(testResults)

	results := Responses{}
	for i := 0; i < NumTests; i++ {
		results = append(results, <-testResults)
	}
	close(testResults)

	resultJson, err := json.Marshal(results)
	if err != nil {
		http.Error(w, fmt.Sprintf("{'status':'%s'}", err.Error()), 500)
		return
	}
	w.Write([]byte(resultJson))
}

func ParseApiTest(r *http.Request) (APITest, error) {
	var u APITest
	if r.Body == nil {
		return u, errors.New("Request does not have body.")
	}
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		return u, err
	}

	vars := mux.Vars(r)
	u.APIProxy, err = apiproxy.Get(strings.ToLower(vars["tenant"]), vars["name"], os.Getenv("SCPI_AUTH"))

	if err != nil {
		return u, err
	}

	jwt := strings.Split(u.Token, ".")
	if len(jwt) < 3 {
		return u, errors.New("Invalid Token.")
	}

	// Format piece of token so it is valid base64
	for len(jwt[1])%4 != 0 {
		jwt[1] += "="
	}
	jwtJsonString, err := base64.StdEncoding.DecodeString(jwt[1])
	if err != nil {
		return u, errors.New("Invalid Token.")
	}

	jwtJsonStringReader := bytes.NewReader(jwtJsonString)
	err = json.NewDecoder(jwtJsonStringReader).Decode(&u)
	if err != nil {
		return u, errors.New("Invalid Token.")
	}
	u.Method = "GET"
	return u, nil
}

func (a *APITest) ExecuteTests(c chan TestResult) {
	go a.UnauthorizedClientTest(c, "unauthorized client test")
	go a.APICallTest(c, "api authentication test")
	go KVMAuthorizationTest(c, a.APIProxy.Tenant, a.APIProxy.Name, a.TokenClientID,
		a.APIProxy.Auth, "kvm authorization test")
}

func (a *APITest) APICallTest(c chan TestResult, name string) {
	client := &http.Client{}
	req, err := http.NewRequest(a.Method, a.APIProxy.Url+a.Path, strings.NewReader(a.Payload))
	if err != nil {
		c <- TestResult{name, false, err}
		return
	}

	req.Header.Add("Authorization", "Bearer " + a.Token)
	resp, err := client.Do(req)
	if err != nil {
		c <- TestResult{name, false, err}
		return
	}
	defer resp.Body.Close()
	if err != nil {
		c <- TestResult{name, false, err}
		return
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c <- TestResult{name, false, err}
		return
	}

	a.Result = string(bodyBytes)
	c <- TestResult{name, (resp.StatusCode >= 200 && resp.StatusCode < 300), nil}
}

func (a *APITest) UnauthorizedClientTest(c chan TestResult, name string) {
	token, err := okta.GenerateToken("nike.sapae.unauthorizedid", os.Getenv("UNAUTHORIZEDID_SECRET"), false)

	if err != nil {
		c <- TestResult{name, false, err}
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(a.Method, a.APIProxy.Url+a.Path, strings.NewReader(a.Payload))
	if err != nil {
		c <- TestResult{name, false, err}
		return
	}

	req.Header.Add("Authorization", "Bearer " + token)
	resp, err := client.Do(req)
	if err != nil {
		c <- TestResult{name, false, err}
		return
	}
	resp.Body.Close()

	if err != nil {
		c <- TestResult{name, false, err}
		return
	}
	// CallAPI should retun a non 200 response
	c <- TestResult{name, !(resp.StatusCode >= 200 && resp.StatusCode < 300), nil}
}
