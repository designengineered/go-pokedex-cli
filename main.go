package main

import "github.com/laztaxon/go-pokedex-cli/internal/pokeapi"

type config struct {
	pokeapiClient pokeapi.Client
	nextAreaURL   *string
	prevAreaURL   *string
	pageCount     int
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(),
		nextAreaURL:   nil,
		prevAreaURL:   nil,
		pageCount:     0,
	}

	startREPL(&cfg)
}
