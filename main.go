package main

import (
	"time"

	"github.com/anxhukumar/pokedexcli/pokeapi"
)

func main() {
	repl(pokeapi.NewClient(5*time.Second, time.Minute*5))
}
