package main

import (
	"fmt"
	"time"
)

func commandExplore(cfg *config) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextAreaURL)
