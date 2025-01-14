package api_test

import (
	"testing"

	"github.com/nickemma/cryptomaster_and_web_server/api"
)

func TestApiCall(t *testing.T) {
	_, err := api.GetRate("")

	if err == nil {
		t.Error("Error was not found")
	}
}
