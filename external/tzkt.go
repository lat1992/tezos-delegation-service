package external

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/lat1992/tezos-delegation-service/structs"
)

type TezosClient struct {
	url    string
	client *http.Client
}

func NewTezosClient(url string) *TezosClient {
	return &TezosClient{
		url: url + "/operations/delegations",
		client: &http.Client{
			Timeout: 15 * time.Second,
		},
	}
}

// GetDelegationsFrom retrieves a batch of delegations starting from the specified offset
func (tc *TezosClient) GetDelegationsFrom(offset int) ([]structs.Delegation, error) {
	req, err := http.NewRequest(http.MethodGet, tc.url+"?limit=10000&offset="+strconv.Itoa(offset), nil)
	if err != nil {
		return nil, fmt.Errorf("create new request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := tc.client.Do(req)
	if err != nil {
		delegations, _ := tc.GetDelegationsFrom(offset)
		return delegations, nil
	}

	var delegations []structs.Delegation
	if err := json.NewDecoder(resp.Body).Decode(&delegations); err != nil {
		return nil, fmt.Errorf("decoding json data: %w", err)
	}
	return delegations, nil
}
