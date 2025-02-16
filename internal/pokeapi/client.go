package pokeapi

import (
	"encoding/json"
	"net/http"
)

var api_url = "https://pokeapi.co/api/v2/"

// get makes a GET request to endpoint and stores the decoded
// resopnse in the value pointed to by v
//
// unless it was found in the cache
func get(endpoint string, v any) error {
	full_url := api_url + endpoint

	req, err := http.NewRequest(http.MethodGet, full_url, nil)
	if err != nil {
		return err
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	dec := json.NewDecoder(res.Body)
	return dec.Decode(v)
}
