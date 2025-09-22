package command

import (
	"fmt"
	"math/rand"

	"github.com/anxhukumar/pokedexcli/pokeapi"
)

func commandCatch(client *pokeapi.Client, config *pokeapi.Config, areaName string, pokemon string) error {
	url := "https://pokeapi.co/api/v2/pokemon/"
	pokemonData, err := client.GetPokemonData(url, pokemon)
	if err != nil {
		return err
	}

	prob := rand.Intn(pokemonData.BaseExperience)

	fmt.Println("Throwing a Pokeball at " + pokemon + "...")

	if prob > 40 {
		fmt.Println(pokemon + " escaped!")
		return nil
	}

	fmt.Println(pokemon + " was caught!")
	pokeapi.Pokemoncollection[pokemon] = pokemonData
	return nil
}
