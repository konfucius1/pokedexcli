package main

import (
	"errors"
	"fmt"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a valid pokemon name")
	}

	name := args[0]

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	res, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	chance := res.BaseExperience
	fmt.Println(chance)

	return nil
}
