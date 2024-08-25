package client

import (
	"net/http"

	"github.com/SamCederwall/ups-go/internal/utils"
)

// The UPS API utilizes OAuth as their authorization process. Because of this, access tokens exist on
// temporarily and expire. This package provides a method for generating an access token (`GenerateAccessToken`) for use with
// the rest of the API Client, however, it is not ideal that a new access token is generated for each
// instance of the client.
type UpsClient struct {
	accessToken string
	HttpClient  *http.Client
}

// Builds a new UPS Client for making requests to the UPS API.
func New(accessToken string) (*UpsClient, error) {
	if err := utils.AccessTokenIsValid(accessToken); err != nil {
		return nil, err
	}

	client := UpsClient{
		accessToken: accessToken,
		HttpClient:  http.DefaultClient,
	}

	return &client, nil
}
