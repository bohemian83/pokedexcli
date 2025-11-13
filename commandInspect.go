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
		fmt.Printf("-hp: %v\n", pokemon.Stats[0].BaseStat)
		fmt.Printf("-attack: %v\n", pokemon.Stats[1].BaseStat)
		fmt.Printf("-defense: %v\n", pokemon.Stats[2].BaseStat)
		fmt.Printf("-special-attack: %v\n", pokemon.Stats[3].BaseStat)
		fmt.Printf("-special-defense: %v\n", pokemon.Stats[4].BaseStat)
		fmt.Printf("-speed: %v\n", pokemon.Stats[5].BaseStat)
		fmt.Println("Types:")
		for _, types := range pokemon.Types {
			fmt.Printf(" - %v\n", types.Type.Name)
		}
		return nil
	}

	fmt.Printf("You have not yet caught %v.\n", pokemonName)

	return nil
}
