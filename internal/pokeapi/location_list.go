package pokeapi

import (
	"net/http"
	"io"
	"encoding/json"
)

const base_url = "https://pokeapi.co/api/v2"

func (c *Client) ListLocations(pageURL *string) (locationArea, error) {
	url := base_url + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.mycache.Get(url); ok {
		locationRes := locationArea{}
		err := json.Unmarshal(val, &locationRes)
		if err != nil {
			return locationRes, err
		}
		return locationRes, nil
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

func (c *Client) GetLocation(argument *string) (habitat, error) {
	url := base_url + "/location-area/" + *argument

	if val, ok := c.mycache.Get(url); ok {
		locationInfo := habitat{}
		err := json.Unmarshal(val, &locationInfo)
		if err != nil {
			return locationInfo, err
		}
		return locationInfo, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return habitat{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return habitat{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return habitat{}, err
	}

	locationInfo := habitat{}
	err = json.Unmarshal(data, &locationInfo)
	if  err != nil {
		return habitat{}, err
	}

	return locationInfo, nil

}

