package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationPokemon(locationArea string) (AreaPokemonResp, error) {
	endpoint := "location-area/" + locationArea
	fullURL := baseURL + endpoint
	var areaPokemonResp AreaPokemonResp

	//check cache first
	cache, found := c.cache.Get(fullURL)
	if found {
		err := json.Unmarshal(cache, &areaPokemonResp)
		if err != nil {
			return areaPokemonResp, err
		}
		return areaPokemonResp, nil
	}
	//if not found in cache, make request

	req, err := http.NewRequest("GET", fullURL, nil)

	if err != nil {
		return areaPokemonResp, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return areaPokemonResp, err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		return areaPokemonResp, fmt.Errorf("error code: %d", resp.StatusCode)

	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return areaPokemonResp, err
	}
	err = json.Unmarshal(data, &areaPokemonResp)
	if err != nil {
		return areaPokemonResp, err
	}

	//add to cache
	c.cache.Add(fullURL, data)

	return areaPokemonResp, nil
}
