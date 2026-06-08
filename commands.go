package main

import (
	"fmt"
	"os"

	"github.com/KriKri98/pokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*pokeapi.Client) error
}

func commandExit(client *pokeapi.Client) error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp(client *pokeapi.Client) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	commands := getCommands()
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandMap(client *pokeapi.Client) error {
	if client.Config.Next == "" {
		return fmt.Errorf("no next maps")
	}
	data, err := client.Get(client.Config.Next)
	if err != nil {
		return err
	}
	if next, ok := data["next"].(string); ok {
		client.Config.Next = next
	} else {
		client.Config.Next = "" // or however your config represents "no next page"
	}
	if previous, ok := data["previous"].(string); ok {
		client.Config.Previous = previous
	} else {
		client.Config.Previous = "" // or however your config represents "no next page"
	}

	locations := data["results"].([]any)
	for _, l := range locations {
		loc := l.(map[string]any)
		fmt.Println(loc["name"].(string))
	}
	return nil

}

func commandMapb(client *pokeapi.Client) error {
	if client.Config.Previous == "" {
		return fmt.Errorf("no previous maps")
	}
	data, err := client.Get(client.Config.Previous)
	if err != nil {
		return err
	}
	if next, ok := data["next"].(string); ok {
		client.Config.Next = next
	} else {
		client.Config.Next = "" // or however your config represents "no next page"
	}
	if previous, ok := data["previous"].(string); ok {
		client.Config.Previous = previous
	} else {
		client.Config.Previous = "" // or however your config represents "no next page"
	}

	locations := data["results"].([]any)
	for _, l := range locations {
		loc := l.(map[string]any)
		fmt.Println(loc["name"].(string))
	}
	return nil

}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 map locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the last 20 map locations",
			callback:    commandMapb,
		},
	}
}
