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
	fmt.Println(".")
	time.Sleep(time.Second / 4)
	fmt.Println("..")
	time.Sleep(time.Second / 4)
	fmt.Println("...")
	time.Sleep(time.Second / 2)

	fmt.Println("")
	time.Sleep(time.Second / 4)
	fmt.Println("Welcome to Cello's Pokedex")
	time.Sleep(time.Second / 4)
	fmt.Println("")
	time.Sleep(time.Second / 2)
	fmt.Print(" Pokedex > ")

	// reading input
	for reader.Scan() {
		input := reader.Text()
		// get the cleaned input array
		cleaned := cleanInput(input)
		// handle empty return from user
		if len(cleaned) == 0 {
			fmt.Println("")
			time.Sleep(time.Second / 4)
			fmt.Println("Empty input. Please Try Again.")
			time.Sleep(time.Second / 4)
			fmt.Println("")
			fmt.Print(" Pokedex > ")
			continue
		}
		// set first word as command trigger
		commandName := cleaned[0]
		// set the rest of the words as the command parameters
		commandParams := cleaned[1:]

		// switching to using a map of commands

		availableCommands := getCommands()
		// is current command in this list?
		command, ok := availableCommands[commandName]
		// default non-command behavior
		if !ok {
			fmt.Println("")
			time.Sleep(time.Second / 4)
			fmt.Println(commandName, "is not a valid command. Try again.")
			time.Sleep(time.Second / 4)
			fmt.Println("")
			fmt.Print(" Pokedex > ")
			continue
		}
		// if there are parameters, pass them to the command
		if len(commandParams) > 0 {
			err := command.callback(cfg, commandParams)
			if err != nil {
				fmt.Println(err)
			}
		}
		// handle possible command errors
		err := command.callback(cfg)
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
	callback    func(*config) error
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
