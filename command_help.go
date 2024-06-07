package main

import (
	"fmt"
	"time"
)

func commandHelp(cfg *config) error {
	fmt.Println("")
	fmt.Println("Hey! This is the Pokedex help menu.")
	fmt.Println("")
	fmt.Println("Your available commands are:")
	fmt.Println("------------------------------------")
	availableCommands := getCommands()
	for _, cmd := range availableCommands {
		fmt.Printf("%s: %s.", cmd.name, cmd.description)
		fmt.Println("")
		fmt.Println("------------------------------------")
		time.Sleep(time.Second / 4)
	}
	return nil
}
