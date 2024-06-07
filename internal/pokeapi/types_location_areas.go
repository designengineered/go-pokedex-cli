package pokeapi

type LocationAreaResp struct {
	Count int `json:"count"`
	//url it's a pointer because it can be null.
	Next *string `json:"next"`
	//url
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
