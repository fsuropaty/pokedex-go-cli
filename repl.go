package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name          string
	description   string
	callback      func(cfg *config) error
	requiresParam bool
	paramName     string
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
			if command.requiresParam {
				if len(words) < 2 {
					err := fmt.Errorf("Please input a %s", command.paramName)
					fmt.Println(err)
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
			name:          "help",
			description:   "Display a help message",
			callback:      commandHelp,
			requiresParam: false,
			paramName:     "",
		},

		"catch": {
			name:          "catch",
			description:   "Catch a Pokemon",
			callback:      commandCatch,
			requiresParam: true,
			paramName:     "Pokemon name",
		},

		"inspect": {
			name:          "inspect",
			description:   "inspect caught Pokemon",
			callback:      commandInspect,
			requiresParam: true,
			paramName:     "Pokemon name",
		},

		"explore": {
			name:          "explore",
			description:   "Display a list of all the Pokemon located in the location",
			callback:      commandExplore,
			requiresParam: true,
			paramName:     "Location name",
		},

		"map": {
			name:          "map",
			description:   "Display 20 location areas in Pokemon World",
			callback:      commandMap,
			requiresParam: false,
			paramName:     "",
		},

		"mapb": {
			name:          "mapb",
			description:   "Display 20 previous location areas",
			callback:      commandMapb,
			requiresParam: false,
			paramName:     "",
		},

		"exit": {
			name:          "exit",
			description:   "Exit the Pokedex",
			callback:      commandExit,
			requiresParam: false,
			paramName:     "",
		},
	}
}
