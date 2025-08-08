package main

import (
	"errors"
	"fmt"
)

func commandPokedex(c *config) error {
	if len(c.user.Pokedex) == 0 {
		return errors.New("No pokemon have been caught yet")
	}

	fmt.Println("Your Pokedex:")
	for key, _ := range c.user.Pokedex {
		fmt.Printf(" - %v\n", key)
	}
	return nil
}
