package external

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDelegationsFrom(t *testing.T) {
	// Create a mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/operations/delegations" {
			t.Errorf("Expected to request '/operations/delegations', got: %s", r.URL.Path)
		}
		if r.URL.RawQuery != "limit=10000&offset=0" {
			t.Errorf("Expected query 'limit=10000&offset=0', got: %s", r.URL.RawQuery)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[
			{"timestamp": "2023-01-01T00:00:00Z", "sender": {"address": "tz1abc"}, "level": 100, "amount": 1000},
			{"timestamp": "2023-01-02T00:00:00Z", "sender": {"address": "tz1ghi"}, "level": 101, "amount": 2000}
		]`))
	}))
	defer server.Close()

	client := NewTezosClient(server.URL)

	delegations, err := client.GetDelegationsFrom(0)
	assert.NoError(t, err)
	assert.Len(t, delegations, 2)
	assert.Equal(t, "2023-01-01T00:00:00Z", delegations[0].Timestamp)
	assert.Equal(t, "tz1abc", delegations[0].Sender.Address)
	assert.Equal(t, 100, delegations[0].Level)
	assert.Equal(t, 1000, delegations[0].Amount)
}

func TestGetDelegationsFromError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	client := NewTezosClient(server.URL)

	delegations, err := client.GetDelegationsFrom(0)
	assert.Error(t, err)
	assert.Nil(t, delegations)
}
