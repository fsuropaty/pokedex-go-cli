package main

import "fmt"

func commandHelp(cfg *config) error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, c := range getCommands() {
		fmt.Printf("%s: %s\n", c.name, c.description)

	}
	fmt.Println()
	return nil
}
