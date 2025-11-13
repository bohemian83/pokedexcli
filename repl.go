package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/bohemian83/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	nextURL       *string
	previousURL   *string
	pokeapiClient pokeapi.Client
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}

		cmd, existsCmd := getCommands()[input[0]]
		args := []string{}
		if len(input) > 1 {
			args = input[1:]
		}

		if existsCmd {
			err := cmd.callback(cfg, args...)
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

func getCommands() map[string]cliCommand {
	commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message and available commands.",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Shows the next 20 location areas.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Shows the previous 20 location areas.",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Shows a list of pokemons at a specified location.",
			callback:    commandExplore,
		},
	}

	return commands
}

func cleanInput(text string) []string {
	cleaned := strings.Fields(strings.ToLower(text))
	return cleaned
}
