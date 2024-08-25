package utils

import (
	"encoding/json"
	"fmt"
	"regexp"
)

func AccessTokenIsValid(accessToken string) error {
	patternToMatch := `^ey[a-zA-Z0-9._-]{94}\.[a-zA-Z0-9._-]+$`
	regex := regexp.MustCompile(patternToMatch)
	valid := regex.Match([]byte(accessToken))

	if !valid {
		return fmt.Errorf("provided access token does not comply with the regex patter: %s", patternToMatch)
	}

	return nil
}

func CheckStatusCode(statusCode int, rawBody []byte) error {

	if statusCode == 200 {
		return nil
	}

	body := map[string]interface{}{}
	if err := json.Unmarshal(rawBody, &body); err != nil {
		return err
	}

	return fmt.Errorf("failed to make request to the ups api: %s", body)
}
