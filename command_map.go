package main

import (
	"errors"
)

func commandMap(c *config) error {

	locations, err := c.pokeapiClient.ListLocations(c.NextLocationUrl)
	if err != nil {
		return err
	}

	c.NextLocationUrl = locations.Next
	c.PreviousLocationUrl = locations.Previous

	for _, r := range locations.Results {
		println(r.Name)
	}

	return nil
}

func commandMapb(c *config) error {

	if c.PreviousLocationUrl == nil {
		return errors.New("you're on the first page")
	}

	locations, err := c.pokeapiClient.ListLocations(c.PreviousLocationUrl)
	if err != nil {
		return err
	}

	c.NextLocationUrl = locations.Next
	c.PreviousLocationUrl = locations.Previous

	for _, r := range locations.Results {
		println(r.Name)
	}

	return nil
}
