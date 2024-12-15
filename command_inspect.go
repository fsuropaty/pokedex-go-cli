package main

import "fmt"

func commandInspect(cfg *config) error {

	pokemon, exists := cfg.pokedex[*cfg.param]
	if !exists {
		return fmt.Errorf("You have not caught %s", *cfg.param)
	}

	fmt.Println("Name: ", pokemon.Name)
	fmt.Println("Height: ", pokemon.Height)
	fmt.Println("Weight: ", pokemon.Weight)
	fmt.Println("Stats: ")
	for _, st := range pokemon.Stats {
		fmt.Printf(" -%s: %d\n", st.Stat.Name, st.BaseStat)
	}

	fmt.Println("Type(s): ")
	for _, ty := range pokemon.Types {
		fmt.Printf(" -%s\n", ty.Type.Name)
	}

	fmt.Println()

	return nil

}
