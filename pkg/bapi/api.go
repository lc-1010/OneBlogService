package bapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	APP_KEY    = "APP_KEY"
	APP_SECRET = "APP_SECRET"
)

// API blog api sdk
type API struct {
	URL    string
	Client *http.Client
}

// AccessToken
type AccessToken struct {
	Token string `json:"token"`
}

// NewAPI creates a new API instance.
//
// It takes a URL string as a parameter and returns a pointer to an API object.
func NewAPI(url string) *API {
	return &API{URL: url, Client: http.DefaultClient}
}

func (a *API) getAccessToken(c *gin.Context) (string, error) {

	body, err := a.httpGet(c, fmt.Sprintf("%s?app_key=%s&app_secret=%s", "auth", APP_KEY, APP_SECRET))
	if err != nil {
		return "", err
	}
	var accessToken AccessToken
	_ = json.Unmarshal(body, &accessToken)
	return accessToken.Token, nil
}

func (a *API) httpGet(ctx context.Context, path string) ([]byte, error) {
	resp, err := http.Get(fmt.Sprint("%s/%s", a.URL, path))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body, nil
}
