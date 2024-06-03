package main

import (
	"fmt"
)

func commandHelp() error {
	fmt.Printf("\nWelcome to the Pokedex!")
	fmt.Printf("\nUsage:\n\n")

	availableCommands := getAvailableCommands()

	for command, details := range availableCommands {
		fmt.Printf("%s: %s\n", command, details.description)
	}

	fmt.Printf("\n")

	return nil
}
