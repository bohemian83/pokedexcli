package main

import (
	"fmt"
)

func commandMap(cfg *config, args ...string) error {
	response, err := cfg.pokeapiClient.GetLocationAreas(cfg.nextURL)
	if err != nil {
		return err
	}

	for _, result := range response.Results {
		fmt.Println(result.Name)
	}

	cfg.nextURL = response.Next
	cfg.previousURL = response.Previous

	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.previousURL == nil {
		fmt.Println("You are on the first page.")
		return nil
	}
	response, err := cfg.pokeapiClient.GetLocationAreas(cfg.previousURL)
	if err != nil {
		return err
	}

	for _, result := range response.Results {
		fmt.Println(result.Name)
	}

	cfg.nextURL = response.Next
	cfg.previousURL = response.Previous

	return nil
}
