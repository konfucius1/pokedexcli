package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func ListLocations() (LocationAreaResponse, error) {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area")
	if err != nil {
		return LocationAreaResponse{}, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	data, err := io.ReadAll(res.Body)

	if res.StatusCode != http.StatusOK {
		return LocationAreaResponse{}, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	if err != nil {
		return LocationAreaResponse{}, err
	}

	locationAreaResponse := LocationAreaResponse{}
	err = json.Unmarshal(data, &locationAreaResponse)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	return locationAreaResponse, nil
}
