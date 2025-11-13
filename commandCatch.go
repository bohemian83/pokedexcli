package main

import (
	"fmt"
	"math/rand/v2"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("Usage: catch <pokemon name>\n")
	}

	pokemonName := args[0]
	_, exists := cfg.pokeapiClient.Pokedex.Pokemons[pokemonName]
	if exists {
		fmt.Printf("You have already caught %v!\n", pokemonName)
		return nil
	}

	response, err := cfg.pokeapiClient.GetPokemonData(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", pokemonName)
	catchChance := rand.IntN(200) + 50

	if catchChance < response.BaseExperience {
		fmt.Printf("%v escaped! Catch chance: %v, Base Experience: %v\n", pokemonName, catchChance, response.BaseExperience)
		return nil
	}

	fmt.Printf("%v was caught!\n", pokemonName)
	cfg.pokeapiClient.Pokedex.Pokemons[pokemonName] = response

	return nil
}
