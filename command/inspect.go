package command

import (
	"fmt"

	"github.com/anxhukumar/pokedexcli/pokeapi"
)

func commandInspect(client *pokeapi.Client, config *pokeapi.Config, areaName string, pokemon string) error {
	v, ok := pokeapi.PokemonCollection[pokemon]

	if !ok {
		fmt.Println("You have not caught " + pokemon + " yet")
		return nil
	}

	fmt.Printf("Name: %s\n", v.Name)
	fmt.Printf("Height: %d\n", v.Height)
	fmt.Printf("Weight: %d\n", v.Weight)
	fmt.Println("Stats:")

	// Map the stat names to the labels you want
	statMap := map[string]string{
		"hp":              "hp",
		"attack":          "attack",
		"defense":         "defense",
		"special-attack":  "special-attack",
		"special-defense": "special-defense",
		"speed":           "speed",
	}

	for _, s := range v.Stats {
		if label, ok := statMap[s.Stat.Name]; ok {
			fmt.Printf(" -%s: %d\n", label, s.BaseStat)
		}
	}

	fmt.Println("Types:")
	for _, t := range v.Types {
		fmt.Printf(" - %s\n", t.Type.Name)
	}

	return nil
}
