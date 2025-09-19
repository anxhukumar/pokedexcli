package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ApiCall(url string) (LocationResponse, error) {
	// create a new request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		error := fmt.Errorf("failed to make get request")
		return LocationResponse{}, error
	}

	// make the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		error := fmt.Errorf("failed to make get request")
		return LocationResponse{}, error
	}
	defer res.Body.Close()

	// decoding response
	var location LocationResponse
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&location); err != nil {
		error := fmt.Errorf("failed to make get request")
		return LocationResponse{}, error
	}

	return location, nil

}
