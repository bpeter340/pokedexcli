package main

import (
	"errors"
	"fmt"
)

func commandExplore(c *config) error {
	if len(c.args) < 1 {
		return errors.New("location name argument is required")
	}

	argument := c.args[0]
	locationArea, err := c.pokeapiClient.GetLocation(argument)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", locationArea.Name)
	for _, encounter := range locationArea.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}
