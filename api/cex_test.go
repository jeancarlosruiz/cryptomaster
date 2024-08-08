package api_test

import (
	"testing"

	"jeanruiz.dev/go/cryptopmaster/api"
)

func TestAPIcall(t *testing.T) {
	_, err := api.GetRate("")
	if err == nil {
		t.Errorf("error was found")
	}
}
