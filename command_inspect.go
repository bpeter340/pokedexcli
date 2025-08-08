package main

import (
	"errors"
	"fmt"
)

func commandInspect(c *config) error {
	if len(c.args) < 1 {
		return errors.New("Must have a pokemon argument")
	}

	name := c.args[0]

	pokemon, exists := c.user.Pokedex[name]
	if !exists {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, s := range pokemon.Stats {
		fmt.Printf("  -%v: %v\n", s.Stat.Name, s.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %v\n", t.Type.Name)
	}

	return nil
}
