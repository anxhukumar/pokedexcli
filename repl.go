package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/anxhukumar/pokedexcli/command"
	"github.com/anxhukumar/pokedexcli/pokeapi"
)

func repl() {
	reader := bufio.NewScanner((os.Stdin))

	for {

		fmt.Print("Pokedex > ")
		reader.Scan()

		if len(reader.Text()) == 0 {
			continue
		}

		cleanString := cleanInput(reader.Text())
		cmd := cleanString[0]

		c, ok := command.GetCommands()[cmd]

		if ok {
			c.Callback(&pokeapi.ApiState)
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	res := []string{}

	trimmedStr := strings.TrimSpace(text)
	splitStrings := strings.Split(trimmedStr, " ")

	for i := range splitStrings {
		res = append(res, strings.ToLower(splitStrings[i]))
	}

	return res
}
