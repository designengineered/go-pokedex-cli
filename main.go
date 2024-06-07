package main

import "github.com/laztaxon/go-pokedex-cli/internal/pokeapi"

type config struct {
	pokeapiClient pokeapi.Client
	nextAreaURL   *string
	prevAreaURL   *string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(),
		nextAreaURL:   nil,
		prevAreaURL:   nil,
	}

	startREPL(&cfg)
}
