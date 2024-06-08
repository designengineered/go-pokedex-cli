package main

import (
	"fmt"
	"log"
	"time"
)

func commandMap(cfg *config) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextAreaURL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("")
	if cfg.pageCount == 0 {
		fmt.Printf("Location Areas(1):")
		fmt.Println("")
		time.Sleep(time.Second / 4)
	}
	if cfg.pageCount != 0 {
		fmt.Printf("Location Areas(%d):\n", cfg.pageCount+1)
		time.Sleep(time.Second / 4)
		fmt.Println("")
	}
	for _, result := range resp.Results {
		fmt.Printf("- %s\n", result.Name)
	}
	// if there is a next area, update next and prev pointers
	cfg.nextAreaURL = resp.Next
	if cfg.nextAreaURL == nil {
		return fmt.Errorf("no next area to go back to")
	}
	cfg.prevAreaURL = resp.Previous
	cfg.pageCount++
	return nil
}
