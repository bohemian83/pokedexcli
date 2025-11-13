package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetAreaPokemons(areaURL string) (ResponsePokemonList, error) {
	url := baseURL + "/location-area/" + areaURL
	if areaURL == "" {
		return ResponsePokemonList{}, fmt.Errorf("error: area not defined")
	}

	cachedVal, exists := c.cache.Get(url)
	if exists {
		response := ResponsePokemonList{}
		err := json.Unmarshal(cachedVal, &response)

		if err != nil {
			return ResponsePokemonList{}, fmt.Errorf("error: decoding response: %v", err)
		}

		return response, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return ResponsePokemonList{}, fmt.Errorf("error: fetching data: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return ResponsePokemonList{}, fmt.Errorf("error: response failed with status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return ResponsePokemonList{}, fmt.Errorf("error: reading response body: %v", err)
	}

	c.cache.Add(url, body)

	response := ResponsePokemonList{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return ResponsePokemonList{}, fmt.Errorf("error: decoding response: %v", err)
	}

	return response, nil
}
