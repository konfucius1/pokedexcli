package main

import (
	"errors"
	"fmt"
	"github.com/konfucius1/pokedexcli/internal/pokeapi"
)

func logPokemon(pokemon pokeapi.Pokemon) {
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)

	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" - %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, stat := range pokemon.Types {
		fmt.Printf(" - %s\n", stat.Type.Name)
	}
}

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("inspect requires a valid pokemon name")
	}

	pokemonName := args[0]

	pokemon, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		return errors.New("you have not caught that pokemon")
	}

	logPokemon(pokemon)

	return nil
}
