package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocations(pageURL *string) (RespLocations, error) {
	url := baseURL + "/location"
	if pageURL != nil {
		url = *pageURL
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

	locations := RespLocations{}
	err = json.Unmarshal(data, &locations)
	if err != nil {
		return RespLocations{}, err
	}

	return locations, err
}
