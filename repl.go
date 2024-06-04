package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	availableCmds := getAvailableCommands()

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()

		cmdName, noCmdSpecified := cleanInput(input)
		if noCmdSpecified != nil {
			continue
		}

		command, ok := availableCmds[cmdName]
		if !ok {
			fmt.Printf("%s: command not found\n", input)
		}

		err := command.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(input string) (commandName string, noCommand error) {
	loweredInput := lowerInput(input)
	words := separateInput(loweredInput)

	if len(words) == 0 {
		return "", fmt.Errorf("no command specified")
	}

	return words[0], nil
}

func lowerInput(text string) (loweredInput string) {
	return strings.ToLower(text)
}

func separateInput(output string) (words []string) {
	return strings.Fields(output)
}

func getAvailableCommands() map[string]cliCommand {
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
		"map": {
			name:        "map",
			description: "Displays next pokemon locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapB",
			description: "Displays previous pokemon locations",
			callback:    commandMapB,
		},
	}
}
