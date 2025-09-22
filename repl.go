package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/anxhukumar/pokedexcli/command"
	"github.com/anxhukumar/pokedexcli/pokeapi"
)

func repl(client pokeapi.Client) {
	reader := bufio.NewScanner((os.Stdin))

	for {

		fmt.Print("Pokedex > ")
		reader.Scan()

		if len(reader.Text()) == 0 {
			continue
		}

		cleanString := cleanInput(reader.Text())
		cmd := cleanString[0]

		areaName := ""

		//set areaName for the commandExplore
		if cmd == "explore" {
			if len(cleanString) > 1 {
				areaName = cleanString[1]
			} else {
				fmt.Println("!Missing area input: explore ?")
				continue
			}
		}

		c, ok := command.GetCommands()[cmd]

		if ok {
			c.Callback(&client, &pokeapi.ApiState, areaName)
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	words := strings.Fields(text)
	for i := range words {
		words[i] = strings.ToLower(words[i])
	}
	return words
}
