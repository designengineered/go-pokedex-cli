package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreaResp, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint
	// if there is a page url, append it to the full url
	if pageURL != nil {
		fullURL = *pageURL
	}
	var locationAreaResp LocationAreaResp

	req, err := http.NewRequest("GET", fullURL, nil)

	if err != nil {
		return locationAreaResp, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return locationAreaResp, err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		return locationAreaResp, fmt.Errorf("error code: %d", resp.StatusCode)

	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return locationAreaResp, err
	}
	err = json.Unmarshal(data, &locationAreaResp)
	if err != nil {
		return locationAreaResp, err
	}
	return locationAreaResp, nil
}
