package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationAreas(pageURL *string) (ResponseLocationAreas, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	cachedVal, exists := c.cache.Get(url)
	if exists {
		response := ResponseLocationAreas{}
		err := json.Unmarshal(cachedVal, &response)

		if err != nil {
			return ResponseLocationAreas{}, fmt.Errorf("error decoding response: %v", err)
		}

		return response, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return ResponseLocationAreas{}, fmt.Errorf("error fetching data: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return ResponseLocationAreas{}, fmt.Errorf("response failed with status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return ResponseLocationAreas{}, fmt.Errorf("error reading response body: %v", err)
	}

	c.cache.Add(url, body)

	response := ResponseLocationAreas{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return ResponseLocationAreas{}, fmt.Errorf("error decoding response: %v", err)
	}

	return response, nil
}
