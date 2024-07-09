package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/imeltsner/pokedex-cli/internal/pokeapi"
)

func cleanInput(s string) []string {
	clean := strings.ToLower(s)
	return strings.Fields(clean)
}

func startRepl() {
	fmt.Println("Welcome to the Pokedex")
	fmt.Println("Usage:")
	config := pokeapi.Config{}
	commandHelp(&config)

	for {
		fmt.Print("pokedex > ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		cmds := getCommands()
		if cmd, ok := cmds[words[0]]; !ok {
			fmt.Println("invalid command")
			continue
		} else {
			err := cmd.callback(&config)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
