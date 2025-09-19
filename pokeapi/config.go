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

func Next() string {
	ApiState.Next = fmt.Sprintf("https://pokeapi.co/api/v2/location-area?offset=%d&limit=20", ApiState.Offset)
	ApiState.Offset += 20
	return ApiState.Next
}

func Previous() string {
	if ApiState.Offset < 40 {
		fmt.Println("This is first page")
		return ""
	}
	ApiState.Offset -= 40
	ApiState.Previous = fmt.Sprintf("https://pokeapi.co/api/v2/location-area?offset=%d&limit=20", ApiState.Offset)
	return ApiState.Previous
}
