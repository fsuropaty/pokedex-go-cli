package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config) error {

	fmt.Printf("Throwing a Pokeball at %s...\n", *cfg.param)
	fmt.Println()

	pokemonResp, err := cfg.pokeapiClient.GetPokemon(cfg.param)
	if err != nil {
		return err
	}

	if !calculateProbability(pokemonResp.BaseExperience) {
		fmt.Printf("%s escaped!\n", pokemonResp.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemonResp.Name)
	cfg.pokedex[pokemonResp.Name] = pokemonResp

	fmt.Println()

	return nil

}

func calculateProbability(baseExp int) bool {
	catchRate := 100 - (float64(baseExp) / 4)

	if catchRate < 0.0 {
		catchRate = 5.0
	}

	randomNum := rand.Intn(100)

	return float64(randomNum) < catchRate

}
