package main

import (
	"fmt"
)

func commandHelp(cfg *config) error {
	fmt.Printf("\nWelcome to the Pokedex!")
	fmt.Printf("\nUsage:\n\n")

	availableCommands := getAvailableCommands()

	for _, details := range availableCommands {
		fmt.Printf("%s: %s\n", details.name, details.description)
	}

	fmt.Printf("\n")

	return nil
}
