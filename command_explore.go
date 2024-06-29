package main

import (
	"fmt"
	"time"
)

func commandExplore(cfg *config, args []string) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextAreaURL)
	if err != nil {
		return err
	}
	return nil
}
