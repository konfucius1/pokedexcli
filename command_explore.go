package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	name := args[0]

	fmt.Printf("Exploring %s...\n", name)

	res, err := cfg.pokeapiClient.GetLocation(name)
	if err != nil {
		return err
	}

	for _, pokemon := range res.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
