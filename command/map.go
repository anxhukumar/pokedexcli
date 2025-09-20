package command

import (
	"fmt"

	"github.com/anxhukumar/pokedexcli/pokeapi"
)

func commandMap(client *pokeapi.Client, config *pokeapi.Config) error {
	mapSlice, err := client.ApiCall(pokeapi.Next())
	if err != nil {
		return err
	}
	for _, v := range mapSlice.Results {
		fmt.Println(v.Name)
	}
	return nil
}

func commandMapB(client *pokeapi.Client, config *pokeapi.Config) error {
	mapSlice, err := client.ApiCall(pokeapi.Previous())
	if err != nil {
		return err
	}
	for _, v := range mapSlice.Results {
		fmt.Println(v.Name)
	}
	return nil
}
