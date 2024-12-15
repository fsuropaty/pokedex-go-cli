package main

import "fmt"

func commandExplore(cfg *config) error {

	fmt.Println()
	fmt.Printf("Explore %s...\n", *cfg.param)

	locationPokemonResp, err := cfg.pokeapiClient.GetLocationPokemon(cfg.param)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon :")
	for _, loc := range locationPokemonResp.PokemonEncounters {

		fmt.Printf(" - %s\n", loc.Pokemon.Name)
	}

	return nil

}
