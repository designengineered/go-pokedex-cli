package main

import (
	"fmt"
)

func commandMap(cfg *config) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextAreaURL)
	if err != nil {
		return err
	}
	fmt.Println("")
	if cfg.pageCount == 0 {
		fmt.Println("---------------------------------------------------------")
		fmt.Printf("Location Areas(1):\n")
		fmt.Println("---------------------------------------------------------")
	}
	if cfg.pageCount != 0 {
		fmt.Printf("Location Areas(%d):\n", cfg.pageCount+1)
		fmt.Println("")
	}
	for _, result := range resp.Results {
		fmt.Printf("- %s\n", result.Name)
	}
	fmt.Println("---------------------------------------------------------")
	// if there is a next area, update next and prev pointers
	cfg.nextAreaURL = resp.Next
	if cfg.nextAreaURL == nil {
		return fmt.Errorf("no next area to go back to")
	}
	cfg.prevAreaURL = resp.Previous
	cfg.pageCount++
	return nil
}
