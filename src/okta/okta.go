package okta

import (
	"net/http"
	"errors"
	"encoding/json"
	"net/url"
	"strings"
)

const oktaQAURL string = "https://nike-qa.oktapreview.com/oauth2/ausa0mcornpZLi0C40h7/v1/token"
const oktaProductionURL string = "https://nike.okta.com/oauth2/aus27z7p76as9Dz0H1t7/v1/token"

type oktaResponse struct {
	Token string `json:"access_token"`
}

func GenerateToken(clientID, secret string, production bool) (string, error) {
	data := url.Values{}
	data.Set("client_id", clientID)
	data.Set("client_secret", secret)
	data.Set("grant_type", "client_credentials")

	oktaURL := oktaQAURL
	if production {
		oktaURL = oktaProductionURL
	}

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

	var r oktaResponse
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return "", err
	}

	return r.Token, nil
}
