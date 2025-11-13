package main

import (
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("Usage: inspect <pokemon name>\n")
	}

	pokemonName := args[0]
	pokemon, exists := cfg.pokeapiClient.Pokedex.Pokemons[pokemonName]

	if exists {
		fmt.Printf("Name: %v\n", pokemonName)
		fmt.Printf("Height: %v\n", pokemon.Height)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, types := range pokemon.Types {
			fmt.Printf(" - %v\n", types.Type.Name)
		}
		return nil
	}

	fmt.Printf("You have not yet caught %v.\n", pokemonName)

	return nil
}
