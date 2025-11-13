package main

import "fmt"

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("Usage: explore <area name>\n")
	}

	areaName := args[0]

	response, err := cfg.pokeapiClient.GetAreaPokemons(areaName)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %v...\n", areaName)
	fmt.Println("Found Pokemon:")

	for _, pokemon := range response.PokemonEncounters {
		fmt.Printf(" - %v\n", pokemon.Pokemon.Name)
	}

	return nil
}
