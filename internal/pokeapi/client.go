package pokeapi

import (
	"net/http"
	"time"

	"github.com/bohemian83/pokedexcli/internal/pokecache"
)

type Client struct {
	Pokedex    Pokedex
	httpClient http.Client
	cache      pokecache.Cache
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		Pokedex: Pokedex{Pokemons: make(map[string]Pokemon)},
		cache:   pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
