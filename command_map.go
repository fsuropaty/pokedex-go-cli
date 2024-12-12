package main

import (
	"fmt"
	"github.com/fsuropaty/go-pokedexcli/internal/pokeapi"
)

func commandMap(cfg *Config) error {

	url := "https://pokeapi.co/api/v2/location-area"

	if cfg.Next != "" {
		url = cfg.Next
	}

	client := &pokeapi.Client{}
	locations, err := client.GetRequest(url)
	if err != nil {
		return err
	}

	cfg.Next = locations.Next

	if locations.Previous != nil {
		cfg.Previous = *locations.Previous
	} else {
		cfg.Previous = ""
	}

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(cfg *Config) error {
	if cfg.Previous == "" {
		fmt.Println("You are on the first page")
		return nil
	}

	client := &pokeapi.Client{}
	locations, err := client.GetRequest(cfg.Previous)
	if err != nil {
		return err
	}

	cfg.Next = locations.Next
	if locations.Previous != nil {
		cfg.Previous = *locations.Previous
	} else {
		cfg.Previous = ""
	}

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
