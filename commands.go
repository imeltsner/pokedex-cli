package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/imeltsner/pokedex-cli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(c *pokeapi.Config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Shows the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Shows the previous 20 locations",
			callback:    commandMapB,
		},
	}
}

func commandHelp(c *pokeapi.Config) error {
	for _, cmd := range getCommands() {
		fmt.Printf("%s - %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandExit(c *pokeapi.Config) error {
	os.Exit(0)
	return nil
}

func commandMap(c *pokeapi.Config) error {
	var link string
	if c.Next == "" {
		link = "https://pokeapi.co/api/v2/location-area"
	} else {
		link = c.Next
	}

	res, err := pokeapi.GetLocationArea(c, link)
	if err != nil {
		return err
	}

	for _, v := range res.Results {
		fmt.Println(v.Name)
	}

	return nil
}

func commandMapB(c *pokeapi.Config) error {
	if c.Prev == "" {
		return errors.New("no previous data")
	}

	link := c.Prev
	res, err := pokeapi.GetLocationArea(c, link)
	if err != nil {
		return err
	}
	for _, v := range res.Results {
		fmt.Println(v.Name)
	}
	return nil
}
