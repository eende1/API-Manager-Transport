package apitesting

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const oktaURL string = "https://nike-qa.oktapreview.com/oauth2/ausa0mcornpZLi0C40h7/v1/token"

const NumTests = 3

type ApiTest struct {
	Url           string
	Token         string
	MetaDataPath  string
	Tenant        string
	APIName       string
	TokenClientID string `json:"cid"`
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

func ParseApiTest(r *http.Request) (ApiTest, error) {
	var u ApiTest
	if r.Body == nil {
		return u, errors.New("Request does not have body.")
	}

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		return u, err
	}

	u.Url = fmt.Sprintf("https://nikescp%s.apimanagement.us2.hana.ondemand.com%s", u.Tenant, u.Url)

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
	return u, nil
}

func (a *ApiTest) ExecuteTests(c chan TestResult) {
	go UnauthorizedClientTest(c, a.Url, a.MetaDataPath, "unauthorized client test")
	go CallAPI(c, a.Url, a.MetaDataPath, a.Token, "api authentication test")
	go KVMAuthorizationTest(c, strings.ToLower(a.Tenant), a.APIName, a.TokenClientID,
		os.Getenv("SCPI_AUTH"), "kvm authorization test")
}

func CallAPI(c chan TestResult, url, metaDataPath, token, name string) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url+metaDataPath, nil)
	if err != nil {
		c <- TestResult{name, false, err}
		return
	}

	req.Header.Add("Authorization", "Bearer "+token)
	resp, err := client.Do(req)

	if err != nil {
		c <- TestResult{name, false, err}
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		c <- TestResult{name, false, nil}
		return
	}

	c <- TestResult{name, true, nil}
}

type OktaResponse struct {
	Token string `json:"access_token"`
}

func GenerateToken(clientID, secret string) (string, error) {
	data := url.Values{}
	data.Set("client_id", clientID)
	data.Set("client_secret", secret)
	data.Set("grant_type", "client_credentials")

	client := &http.Client{}
	req, err := http.NewRequest("POST", oktaURL, strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", errors.New("Non 200 response Code")
	}

	var r OktaResponse
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return "", err
	}

	return r.Token, nil
}

func UnauthorizedClientTest(c chan TestResult, url, metaDataPath, name string) {
	token, err := GenerateToken("nike.sapae.unauthorizedid", os.Getenv("UNAUTHORIZEDID_SECRET"))

	if err != nil {
		c <- TestResult{name, false, err}
		return
	}
	c2 := make(chan TestResult)
	go CallAPI(c2, url, metaDataPath, token, name)
	result := <-c2
	if result.Err != nil {
		c <- TestResult{name, false, result.Err}
		return
	}
	// CallAPI test should fail, so return opposite of CallAPI's result
	c <- TestResult{name, !result.Pass, nil}
}
