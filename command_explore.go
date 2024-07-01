package main

import (
	"fmt"
	"strings"
	// "time"
)

func commandExplore(cfg *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("please provide a location area")
	}
	// support mulitple unhyphenated words
	areaName := strings.Join(args, "-")

	fmt.Printf("Exploring %s\n", areaName)
	// resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextAreaURL, areaName)
	// if err != nil {
	// 	return err
	// }
	return nil
}
