package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocationPokemon(name *string) (RespLocationPokemon, error) {
	url := baseURL + "/location-area/"
	locationPokemon := RespLocationPokemon{}

	if name != nil {
		url += *name
	}

	if data, ok := c.cache.Get(url); ok {
		err := json.Unmarshal(data, &locationPokemon)
		return locationPokemon, err
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationPokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationPokemon{}, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	if err != nil {
		return RespLocationPokemon{}, err
	}

	c.cache.Add(url, data)

	err = json.Unmarshal(data, &locationPokemon)
	if err != nil {
		return RespLocationPokemon{}, err
	}

	return locationPokemon, err
}

