package main

import "fmt"

func commandPokedex(cfg *config) error {

	pokemon := cfg.pokedex
	if len(pokemon) == 0 {
		return fmt.Errorf("You have not caught any Pokemon")
	}

	fmt.Println("Your Pokedex:")
	for _, pk := range pokemon {
		fmt.Printf(" - %s\n", pk.Name)
	}

	fmt.Println()

	return nil

}
