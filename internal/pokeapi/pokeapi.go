package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationArea struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type Config struct {
	next string
	prev string
}

func (c *Config) GetLocationArea() (LocationArea, error) {
	var locationArea LocationArea
	var res *http.Response
	var err error

	if c.next == "" {
		res, err = http.Get("https://pokeapi.co/api/v2/location-area")
	} else {
		res, err = http.Get(c.next)
	}
	if err != nil {
		return locationArea, err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return locationArea, fmt.Errorf("status error %v", res.StatusCode)
	}
	if err != nil {
		return locationArea, err
	}

	err = json.Unmarshal(body, &locationArea)
	if err != nil {
		fmt.Println("Error unmarshaling json")
		return locationArea, err
	}

	c.next = locationArea.Next
	c.prev = locationArea.Previous
	return locationArea, nil
}
