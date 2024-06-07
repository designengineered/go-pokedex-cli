package main

import (
	"fmt"
	"log"
	"time"
)

func commandMapB(cfg *config) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevAreaURL)
	if cfg.prevAreaURL == nil {
		return fmt.Errorf("no previous area to go back to")
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("")
	fmt.Println("Location Areas:")
	fmt.Println("")
	time.Sleep(time.Second / 4)
	for _, result := range resp.Results {
		fmt.Printf("- %s\n", result.Name)
	}
	// if there is a next area, update next and prev pointers
	cfg.nextAreaURL = resp.Next
	cfg.prevAreaURL = resp.Previous
	return nil
}
