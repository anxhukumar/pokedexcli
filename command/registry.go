package command

import (
	"github.com/anxhukumar/pokedexcli/pokeapi"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func(client *pokeapi.Client, config *pokeapi.Config, areaName string, pokemon string) error
}

func GetCommands() map[string]CliCommand {

	return map[string]CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
		"map": {
			Name:        "map",
			Description: "Displays the names of 20 location areas in the Pokemon world",
			Callback:    commandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays the names of previous 20 location areas in the Pokemon world",
			Callback:    commandMapB,
		},
		"explore": {
			Name:        "explore",
			Description: "Displays the list of all the Pokemon in a particular location",
			Callback:    commandExplore,
		},
		"catch": {
			Name:        "catch",
			Description: "Catches Pokemon and adds it to the user's Pokedex",
			Callback:    commandCatch,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Shows the details about a Pokemon if its in the Pokedex",
			Callback:    commandInspect,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "Shows the list of Pokemon in the Pokedex",
			Callback:    commandPokedex,
		},
	}
}
