package command

import (
	"github.com/anxhukumar/pokedexcli/pokeapi"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func(client *pokeapi.Client, config *pokeapi.Config) error
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
	}
}
