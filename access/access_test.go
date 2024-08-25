package access

import (
	"testing"
)

func TestGenerateAccessTokenSucceeds(t *testing.T) {
	_, err := GenerateAccessToken("", "", "")

	if err != nil {
		t.Errorf("errored out while generating access token: %s", err)
	}
}

func TestGenerateAccessTokenFails(t *testing.T) {
	_, err := GenerateAccessToken("TEST", "TEST", "12345")
	if err != nil {
		t.Errorf("errored out while generating access token: %s", err)
	}

}
