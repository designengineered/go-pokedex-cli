package main

import (
	"bufio"
	"fmt"
	"os"
)

func startREPL() {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Println("Pokedex >")

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
