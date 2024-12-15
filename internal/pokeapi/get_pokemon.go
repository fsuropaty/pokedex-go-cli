package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name *string) (RespPokemon, error) {
	url := baseURL + "/pokemon/"
	pokemon := RespPokemon{}

	if name != nil {
		url += *name
	}

	if data, ok := c.cache.Get(url); ok {
		err := json.Unmarshal(data, &pokemon)
		return pokemon, err
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemon{}, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	if err != nil {
		return RespPokemon{}, err
	}

	c.cache.Add(url, data)

	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return RespPokemon{}, err
	}

	return pokemon, err
}
