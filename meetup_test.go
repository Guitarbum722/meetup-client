package meetup

import (
	"net/http"
	"testing"
)

// validateClientType
func validateClientType(c Clienter) bool {
	if _, ok := c.(*Client); !ok {
		return false
	}
	return true
}

const apiKey = "asdfasdf12341234"

// TestNewClient
func TestNewClient(t *testing.T) {
	copts := ClientOpts{
		APIKey: apiKey,
	}
	c := NewClient(&copts)
	if !validateClientType(c) {
		t.Error("returned value NOT of type Clienter")
	}
}

// TestNewClient_with_HTTP
func TestNewClient_with_HTTP(t *testing.T) {
	copts := ClientOpts{
		APIKey:     apiKey,
		HTTPClient: &http.Client{},
	}
	c := NewClient(&copts)
	if !validateClientType(c) {
		t.Error("returned value NOT of type Clienter")
	}
}

// TestNewClient_Nil_Opts
func TestNewClient_Nil_Opts(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("code never panicked")
		}
	}()
	NewClient(nil)
}
