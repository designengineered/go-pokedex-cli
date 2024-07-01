package main

import (
	"fmt"
	"strings"
	"time"
)

func commandExplore(cfg *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("Please provide a location area.")
	}
	// support mulitple unhyphenated words
	areaName := strings.Join(args, "-")

	fmt.Println("")
	fmt.Printf("Exploring %s...\n", areaName)
	time.Sleep(time.Second / 4)
	resp, err := cfg.pokeapiClient.ListLocationPokemon(areaName)
	if err != nil {
		fmt.Println("")
		fmt.Println("Area not found.")
		return err
	}
	fmt.Println("")
	fmt.Println("==================================================")
	fmt.Println("Found the following Pokemon:")
	fmt.Println("==================================================")
	for _, encounter := range resp.PokemonEncounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
		fmt.Println("---------------------------------------------------------")
	}
	return nil
}
