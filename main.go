package main

import (
	"math/rand"
	"time"

	"github.com/fsuropaty/go-pokedexcli/internal/pokeapi"
)

func main() {
	rand.NewSource(time.Now().UnixNano())
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
		pokedex:       make(map[string]pokeapi.RespPokemon),
	}

	repl(cfg)
}
