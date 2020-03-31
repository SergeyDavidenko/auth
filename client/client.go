package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// AuthCheck ...
func AuthCheck(url, token string) (bool error) {
	if token == "" {
		return false, fmt.Errorf("token not set")
	}
	if url == "" {
		return false, fmt.Errorf("auth url not set")
	}
	var tokenJSON TokenValidator
	tokenJSON.Token = token
	jsonStr, errMarshal := json.Marshal(tokenJSON)
	if errMarshal != nil {
		return false, fmt.Errorf("cannot json marshal struct")
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.Status != 200 {
		return false, fmt.Errorf("status conde not qe 200")
	}
	return true, nil
}

// TokenValidator ...
type TokenValidator struct {
	Token string `json:"token"`
}
