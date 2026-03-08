package pokeapi

import (
	"net/http"
	"io"
	"encoding/json"
)

const base_url = "https://pokeapi.co/api/v2"

func (c Client) ListLocations(pageURL *string) (locationArea, error) {
	url := base_url + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return locationArea{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return locationArea{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return locationArea{}, err
	}

	locationRes := locationArea{}
	err = json.Unmarshal(data, &locationRes)
	if  err != nil {
		return locationArea{}, err
	}

	return locationRes, nil

}