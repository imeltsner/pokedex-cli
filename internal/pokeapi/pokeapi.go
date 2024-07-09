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
	Next string
	Prev string
}

func GetLocationArea(c *Config, link string) (LocationArea, error) {
	var locationArea LocationArea

	res, err := http.Get(link)
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

	c.Next = locationArea.Next
	c.Prev = locationArea.Previous
	return locationArea, nil
}
