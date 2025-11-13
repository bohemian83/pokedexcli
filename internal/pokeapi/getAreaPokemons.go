package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetAreaPokemons(areaURL string) (Location, error) {
	url := baseURL + "/location-area/" + areaURL

	cachedVal, exists := c.cache.Get(url)
	if exists {
		response := Location{}
		err := json.Unmarshal(cachedVal, &response)

		if err != nil {
			return Location{}, fmt.Errorf("error: decoding response: %v", err)
		}

		return response, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return Location{}, fmt.Errorf("error: fetching data: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return Location{}, fmt.Errorf("error: response failed with status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Location{}, fmt.Errorf("error: reading response body: %v", err)
	}

	c.cache.Add(url, body)

	response := Location{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return Location{}, fmt.Errorf("error: decoding response: %v", err)
	}

	return response, nil
}
