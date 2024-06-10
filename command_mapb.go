package main

import (
	"fmt"
	"time"
)

func commandMapB(cfg *config) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevAreaURL)
	if cfg.prevAreaURL == nil {
		return fmt.Errorf("no previous area to go back to")
	}
	if err != nil {
		return err
	}
	fmt.Println("")
	if cfg.pageCount == 0 {
		fmt.Println("---------------------------------------------------------")
		fmt.Printf("Location Areas(1):\n")
		fmt.Println("---------------------------------------------------------")
		time.Sleep(time.Second / 4)
	}
	if cfg.pageCount != 0 {
		fmt.Printf("Location Areas(%d):\n", cfg.pageCount-1)
		time.Sleep(time.Second / 4)
		fmt.Println("")
	}
	time.Sleep(time.Second / 4)
	for _, result := range resp.Results {
		fmt.Printf("- %s\n", result.Name)
	}
	// if there is a next area, update next and prev pointers
	cfg.nextAreaURL = resp.Next
	cfg.prevAreaURL = resp.Previous
	cfg.pageCount--
	return nil
}
