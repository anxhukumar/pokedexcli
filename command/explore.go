package command

import (
	"fmt"

	"github.com/anxhukumar/pokedexcli/pokeapi"
)

func commandExplore(client *pokeapi.Client, config *pokeapi.Config, areaName string, pokemon string) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	locationAreaStruct, err := client.GetPokemonInArea(url, areaName)
	if err != nil {
		return err
	}
	fmt.Println("Exploring " + areaName + "...")
	fmt.Println("Found Pokemon:")
	for _, v := range locationAreaStruct.PokemonEncounters {
		fmt.Println(" - " + v.Pokemon.Name)
	}
	return nil
}
