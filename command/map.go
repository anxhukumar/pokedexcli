package command

import (
	"fmt"

	"github.com/anxhukumar/pokedexcli/pokeapi"
)

func commandMap(client *pokeapi.Client, config *pokeapi.Config, areaName string) error {
	mapSlice, err := client.GetLocationAreas(pokeapi.Next())
	if err != nil {
		return err
	}
	for _, v := range mapSlice.Results {
		fmt.Println(v.Name)
	}
	return nil
}

func commandMapB(client *pokeapi.Client, config *pokeapi.Config, areaName string) error {
	mapSlice, err := client.GetLocationAreas(pokeapi.Previous())
	if err != nil {
		return err
	}
	for _, v := range mapSlice.Results {
		fmt.Println(v.Name)
	}
	return nil
}
