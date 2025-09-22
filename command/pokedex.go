package command

import (
	"fmt"

	"github.com/anxhukumar/pokedexcli/pokeapi"
)

func commandPokedex(client *pokeapi.Client, config *pokeapi.Config, areaName string, pokemon string) error {
	fmt.Println("Your Pokedex:")
	for key := range pokeapi.PokemonCollection {
		fmt.Println(" - " + key)
	}
	return nil
}
