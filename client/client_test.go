package client

import "testing"

func TestAuthClient(t *testing.T) {
	valid, err := AuthCheck("http://localhost:8080/api/valid", "test")
	if err == nil {
		t.Error("fake validation")
	}
	if valid {
		t.Error("valid not work")
	}
}