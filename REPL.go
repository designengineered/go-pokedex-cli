package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startREPL() {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to Cello's Pokedex")
	fmt.Print(" Pokedex > ")

	// reading input
	for reader.Scan() {
		input := reader.Text()
		// get the cleaned input array
		cleaned := cleanInput(input)
		// handle empty return from user
		if len(cleaned) == 0 {
			fmt.Println("Empty input. Please Try Again.")
			fmt.Println("")
			fmt.Print(" Pokedex > ")
			continue
		}
		// set first word as command trigger
		commandName := cleaned[0]

		// switching to using a map of commands

		availableCommands := getCommands()
		// is current command in this list?
		command, ok := availableCommands[commandName]
		// default non-command behavior
		if !ok {
			fmt.Println(commandName, "is not a valid command. Try again.")
			fmt.Println("")
			fmt.Print(" Pokedex > ")
			continue
		}
		// handle possible command errors
		err := command.callback()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Print(" Pokedex > ")
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
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
