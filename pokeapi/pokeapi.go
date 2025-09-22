package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/anxhukumar/pokedexcli/internal/pokecache"
)

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

func (c *Client) GetLocationAreas(url string) (LocationResponse, error) {

	// check cache first
	if val, ok := c.cache.Get(url); ok {
		var location LocationResponse
		if err := json.Unmarshal(val, &location); err != nil {
			error := fmt.Errorf("failed to decode cache data")
			return LocationResponse{}, error
		}
		return location, nil
	}

	// create a new request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		error := fmt.Errorf("failed to make get request")
		return LocationResponse{}, error
	}

	// make the request
	client := c.httpClient
	res, err := client.Do(req)
	if err != nil {
		error := fmt.Errorf("failed to make get request")
		return LocationResponse{}, error
	}
	defer res.Body.Close()

	// decoding response
	var location LocationResponse
	val, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationResponse{}, err
	}
	if err := json.Unmarshal(val, &location); err != nil {
		error := fmt.Errorf("failed to decode api response data")
		return LocationResponse{}, error
	}

	//store data in cache
	c.cache.Add(url, val)

	return location, nil

}

func (c *Client) GetPokemonInArea(url string, areaName string) (LocationAreaResponse, error) {

	// check cache first
	if val, ok := c.cache.Get(url); ok {
		var areaRes LocationAreaResponse
		if err := json.Unmarshal(val, &areaRes); err != nil {
			error := fmt.Errorf("failed to decode cache data")
			return LocationAreaResponse{}, error
		}
	}

	url = url + areaName

	// create a new request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		error := fmt.Errorf("failed to make get request")
		return LocationAreaResponse{}, error
	}

	// make the request
	client := c.httpClient
	res, err := client.Do(req)
	if err != nil {
		error := fmt.Errorf("failed to make get request")
		return LocationAreaResponse{}, error
	}
	defer res.Body.Close()

	// decoding response
	var areaRes LocationAreaResponse
	val, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	if err := json.Unmarshal(val, &areaRes); err != nil {
		error := fmt.Errorf("failed to decode api response data")
		return LocationAreaResponse{}, error
	}

	//store data in cache
	c.cache.Add(url, val)

	return areaRes, nil
}
