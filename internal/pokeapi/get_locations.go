package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocations(pageURL *string) (RespLocations, error) {
	url := baseURL + "/location-area"
	locations := RespLocations{}

	if pageURL != nil {
		url = *pageURL
	}

	if data, ok := c.cache.Get(url); ok {
		err := json.Unmarshal(data, &locations)
		return locations, err
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocations{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocations{}, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	if err != nil {
		return RespLocations{}, err
	}

	c.cache.Add(url, data)

	err = json.Unmarshal(data, &locations)
	if err != nil {
		return RespLocations{}, err
	}

	return locations, err
}
