package command

import (
	"fmt"

	"github.com/anxhukumar/pokedexcli/pokeapi"
)

func commandMap(config *pokeapi.Config) error {
	mapSlice, err := pokeapi.ApiCall(pokeapi.Next())
	if err != nil {
		return err
	}
	for _, v := range mapSlice.Results {
		fmt.Println(v.Name)
		fmt.Println(v.Url)
	}
	return nil
}

func commandMapB(config *pokeapi.Config) error {
	mapSlice, err := pokeapi.ApiCall(pokeapi.Previous())
	if err != nil {
		return err
	}
	for _, v := range mapSlice.Results {
		fmt.Println(v.Name)
		fmt.Println(v.Url)
	}
	return nil
}
