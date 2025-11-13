package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemonData(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName

	cachedVal, exists := c.cache.Get(url)
	if exists {
		response := Pokemon{}
		err := json.Unmarshal(cachedVal, &response)

		if err != nil {
			return Pokemon{}, fmt.Errorf("error: decoding response: %v", err)
		}

		return response, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error: fetching data: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return Pokemon{}, fmt.Errorf("error: response failed with status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error: reading response body: %v", err)
	}

	c.cache.Add(url, body)

	response := Pokemon{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error: decoding response: %v", err)
	}

	return response, nil
}
