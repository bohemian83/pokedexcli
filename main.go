package main

import (
	"github.com/bohemian83/pokedexcli/internal/pokeapi"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(60*time.Second, 5*time.Minute)
	cfg := &config{pokeapiClient: pokeClient}
	startRepl(cfg)
}
