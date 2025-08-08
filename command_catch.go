package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"time"
)

func commandCatch(c *config) error {
	rand.Seed(time.Now().UnixNano())

	if len(c.args) < 1 {
		return errors.New("You must have a pokemon argument")
	}
	name := c.args[0]
	fmt.Printf("Throwing a Pokeball at %v...\n", name)

	pokemon, err := c.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}
	const (
		minXP = 20
		maxXP = 700
	)

	normalized := normalize(pokemon.BaseExperience, maxXP, minXP)
	adjusted := math.Pow(normalized, 3)
	chance := int(adjusted * 1000)
	roll := rand.Intn(1000)

	if roll < chance {
		c.user.Pokedex[pokemon.Name] = pokemon
		fmt.Printf("%v was caught!\n", pokemon.Name)
		fmt.Println("You may now inspect it with the inspect command.")
		return nil
	}

	fmt.Printf("%v escaped!\n", pokemon.Name)

	return nil
}

func normalize(value, max, min int) float64 {
	if value > max {
		value = max
	}
	if value < min {
		value = min
	}
	return float64(max-value) / float64(max-min)
}
