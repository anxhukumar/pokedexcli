package command

import (
	"fmt"
	"os"

	"github.com/anxhukumar/pokedexcli/pokeapi"
)

func commandExit(client *pokeapi.Client, config *pokeapi.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
