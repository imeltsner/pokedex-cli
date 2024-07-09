package main

import (
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
	res, err := c.GetLocationArea(true)
	if err != nil {
		return err
	}
	for _, v := range res.Results {
		fmt.Println(v.Name)
	}
	return nil
}

func commandMapB(c *pokeapi.Config) error {
	res, err := c.GetLocationArea(false)
	if err != nil {
		return err
	}
	for _, v := range res.Results {
		fmt.Println(v.Name)
	}
	return nil
}
