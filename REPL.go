package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func startREPL(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Print(" Pokedex > ")

	// reading input
	for reader.Scan() {
		input := reader.Text()
		// get the cleaned input array
		cleaned := cleanInput(input)
		// handle empty return from user
		if len(cleaned) == 0 {
			fmt.Println("\nEmpty input. Please Try Again.\n")
			fmt.Print(" Pokedex > ")
			continue
		}
		// command & arg logic
		commandName := cleaned[0]
		args := cleaned[1:]

		availableCommands := getCommands()
		// is current command in this list?
		command, ok := availableCommands[commandName]
		// default non-command behavior
		if !ok {
			fmt.Printf("\n%s is not a valid command. Try again.\n\n", commandName)
			fmt.Print(" Pokedex > ")
			continue
		}
		// pass config and args to the command function
		err := command.callback(cfg, args)
		if err != nil {
			fmt.Println(err)
		}

		time.Sleep(time.Second / 4)
		fmt.Println("")
		fmt.Print(" Pokedex > ")

	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, args []string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays a map area with 20 locations, \nrun again to display next map area",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Goes back to the previous map area",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Explore a location area for pokemon",
			callback:    commandExplore,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func cleanInput(str string) []string {
	// parse the user input, clean it up, return all the tokens in the string as individual words
	// lowercase the input
	lowered := strings.ToLower(str)
	// use fields method
	words := strings.Fields(lowered)
	// make a slice of the lowered input
	return words
}
