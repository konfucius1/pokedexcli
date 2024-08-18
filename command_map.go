package main

import (
	"fmt"
	"github.com/konfucius1/pokedexcli/internal/pokeapi"
)

func commandMap() error {
	res, err := pokeapi.ListLocations()
	if err != nil {
		return err
	}

	for _, loc := range res.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
