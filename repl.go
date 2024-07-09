package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(s string) []string {
	clean := strings.ToLower(s)
	return strings.Fields(clean)
}

func startRepl() {
	fmt.Println("Welcome to the Pokedex")
	fmt.Println("Usage:")
	fmt.Println("help - displays a help message")
	fmt.Println("exit - exits the program")

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
			cmd.callback()
		}
	}
}
