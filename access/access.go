package access

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/SamCederwall/ups-go/internal/utils"
)

var upsApiVersion = "v1"
var baseEndpointUrl = fmt.Sprintf("https://wwwcie.ups.com/security/%s", upsApiVersion)

type AccessToken struct {
	TokenType    string `json:"token_type"`
	IssuedAt     string `json:"issued_at"`
	ClientId     string `json:"client_id"`
	AccessToken  string `json:"access_token"`
	Scope        string `json:"scope"`
	ExpiresIn    string `json:"expires_in"`
	RefreshCount string `json:"refresh_count"`
	Status       string `json:"status"`
}

// Generate an access token that can be used downstream with the rest of the UPS Client.
// Takes a username and password, which is then base 64 encoded and
// passed as a header to the `oauth/token` endpoint.
func GenerateAccessToken(username string, password string, upsAccountNumber string) (*AccessToken, error) {
	const grantType string = "client_credentials"

	encodedAuth := b64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", username, password)))

	data := url.Values{}
	data.Set("grant_type", grantType)

	request, err := http.NewRequest("POST", fmt.Sprintf("%s/oauth/token", baseEndpointUrl), strings.NewReader(data.Encode()))

	if err != nil {
		return nil, err
	}

	request.Header.Set("content-type", "application/x-www-form-urlencoded")
	request.Header.Set("x-merchant-id", upsAccountNumber)
	request.Header.Set("Authorization", fmt.Sprintf("Basic %s", encodedAuth))

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	statusCode := res.StatusCode
	rawBody, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	if err := utils.CheckStatusCode(statusCode, rawBody); err != nil {
		return nil, err
	}

	var accessToken AccessToken
	err = json.Unmarshal(rawBody, &accessToken)

	if err != nil {
		return nil, err
	}

	return &accessToken, nil

}
