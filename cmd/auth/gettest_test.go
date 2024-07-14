package auth

import (
	"os/user"
	"testing"

	"github.com/zalando/go-keyring"
)

func TestGetToken(t *testing.T) {

	keyring.MockInit()
	user, _ := user.Current()
	err := keyring.Set("durpcli", user.Username, "password")
	if err != nil {
		t.Fatal(err)
	}

	token := gettoken()
	if token != "password" {
		t.Error("password was not expected string")
	}
}
