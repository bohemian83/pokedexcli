package main

import (
	"fmt"
)

func commandHelp(cfg *config, args ...string) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, cmd := range getCommands() {
		fmt.Printf("%s, %s\n", cmd.name, cmd.description)
	}
	return nil
}
