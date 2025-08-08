package main

import (
	"time"

	"github.com/bpeter340/pokedexcli/internal/pokeapi"
)

func main() {
	client := pokeapi.NewClient(5*time.Second, 5*time.Second)
	cfg := &config{
		pokeapiClient: client,
		args:          make([]string, 0),
		user:          pokeapi.NewUser(),
	}

	startRepl(cfg)
}
