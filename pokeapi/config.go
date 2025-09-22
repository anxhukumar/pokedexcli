// This is the state of the API to manage pagination.
package pokeapi

import (
	"fmt"
)

type Config struct {
	Offset   int
	Previous string
	Next     string
}

var ApiState = Config{
	Offset: 0,
}

var PokemonCollection = map[string]PokemonData{}

func Next() string {
	ApiState.Offset += 20
	ApiState.Next = fmt.Sprintf(
		"https://pokeapi.co/api/v2/location-area?offset=%d&limit=20",
		ApiState.Offset-20,
	)
	return ApiState.Next
}

func Previous() string {
	if ApiState.Offset <= 20 {
		fmt.Println("This is first page")
		return ""
	}
	ApiState.Offset -= 20
	ApiState.Previous = fmt.Sprintf(
		"https://pokeapi.co/api/v2/location-area?offset=%d&limit=20",
		ApiState.Offset-20,
	)
	return ApiState.Previous
}
