package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type Client struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) GetRequest(url string) (*Client, error) {

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var data Client

	err = json.Unmarshal(body, &data)

	if err != nil {
		return nil, err
	}

	return &data, nil

}
