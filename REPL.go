package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startREPL() {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Print("Pokedex >")

	// reading input
	for reader.Scan() {
		input := reader.Text()
		//handling input commands
		fmt.Println("Pokedex >", input)

		if input == "exit" {
			break
		}
		fmt.Print("Pokedex > ")
	}

	if err := reader.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
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
