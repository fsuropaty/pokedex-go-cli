package main

import (
	"time"

	"github.com/fsuropaty/go-pokedexcli/internal/pokeapi"
)

func main() {

	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	repl(cfg)
}
