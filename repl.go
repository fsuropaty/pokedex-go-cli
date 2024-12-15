package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fsuropaty/go-pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
}

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	param            *string
}

func repl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := getCommands()[commandName]
		if exists {
			if command.name == "explore" {
				if len(words) < 2 {
					fmt.Println("Please provide the location")
					continue
				}

				cfg.param = &words[1]
			}
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}

	}
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	words := strings.Fields(lower)
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{

		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    commandHelp,
		},

		"explore": {
			name:        "explore",
			description: "Display a list of all the Pokemon located in the location",
			callback:    commandExplore,
		},

		"map": {
			name:        "map",
			description: "Display 20 location areas in Pokemon World",
			callback:    commandMap,
		},

		"mapb": {
			name:        "mapb",
			description: "Display 20 previous location areas",
			callback:    commandMapb,
		},

		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
