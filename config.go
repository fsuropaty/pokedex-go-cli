package main

import "github.com/fsuropaty/go-pokedexcli/internal/pokeapi"

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	param            *string
	pokedex          map[string]pokeapi.RespPokemon
}
