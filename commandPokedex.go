package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {

	fmt.Println("Your Pokedex:")
	if len(cfg.pokeapiClient.Pokedex.Pokemons) > 0 {
		for _, pokemon := range cfg.pokeapiClient.Pokedex.Pokemons {
			fmt.Printf(" - %v\n", pokemon.Name)
		}
		return nil
	}
	fmt.Println("<empty>")
	return nil
}
