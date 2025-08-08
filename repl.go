package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/bpeter340/pokedexcli/internal/pokeapi"
)

type config struct {
	NextLocationUrl     *string
	PreviousLocationUrl *string
	args                []string
	pokeapiClient       pokeapi.Client
	user                pokeapi.User
}

func (c *config) setArgs(data []string) {
	c.args = data
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		cfg.setArgs(words[1:])

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(c *config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"catch": {
			name:        "catch <pokemon name>",
			description: "Attempt to a pokemon",
			callback:    commandCatch,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Explore a location",
			callback:    commandExplore,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"inspect": {
			name:        "inspect <pokemon name>",
			description: "Displays information about a pokemon such as type, stats, etc...",
			callback:    commandInspect,
		},
		"map": {
			name:        "map",
			description: "Get next page of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get previous page of locations",
			callback:    commandMapb,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Display a list of caught pokemon",
			callback:    commandPokedex,
		},
	}
}
