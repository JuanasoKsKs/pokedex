package main

import (
	"fmt"
	"strings"
	"os"
	"bufio"
	"github.com/JuanasoKsKs/pokedex/internal/pokeapi"
)

func startRepl(conf *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}
		commandName := words[0]
		command, ok := getCommands()[commandName]
		if ok {
			if len(words) > 1 {
				conf.argument = &words[1]
			}
			err := command.callback(conf)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
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
			name: "map",
			description: "Display the next 20 Locations",
			callback:	commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Desplay the previous 20 locations",
			callback: commandMapb,
		},
		"explore":{
			name: "explore",
			description: "It takes the name of a location and returns the pokemos found there",
			callback: commandExplore,
		},
		"catch":{
			name: "catch",
			description: "It takes the name of a pokemon and prints if it was able to be caught or not",
			callback: commandCatch,
		},
		"inspect":{
			name: "inspect",
			description: "It shows you the information about the pokemon selected",
			callback: commandInspect,
		},
		"pokedex":{
			name: "pokedex",
			description: "It shows you the pokemons in your Inventory",
			callback: commandPokedex,
		},
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	myclient pokeapi.Client
	next *string
	previous *string
	argument *string
}