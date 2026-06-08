package main

import (
	"fmt"
	"os"
	"strconv"

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
	nextId, err := strconv.Atoi(client.Config.Next[40:])
	if err != nil {
		return err
	}
	for i := 0; i < 20; i++ {
		data, err := client.Get(client.Config.Next[:40] + fmt.Sprint(i+nextId))
		if err != nil {
			fmt.Printf("Error: %v", err)
			return err
		}
		name := data["name"]
		fmt.Println(name)
	}
	if nextId < 20 {
		client.Config.Previous = client.Config.Previous[:40] + fmt.Sprint(1)
	} else {
		client.Config.Previous = client.Config.Next[:40] + fmt.Sprint(nextId-20)
	}
	client.Config.Next = client.Config.Next[:40] + fmt.Sprint(20+nextId)
	return nil

}

func commandMapb(client *pokeapi.Client) error {
	previousId, err := strconv.Atoi(client.Config.Previous[40:])
	if err != nil {
		return err
	}
	for i := 0; i < 20; i++ {
		data, err := client.Get(client.Config.Next[:40] + fmt.Sprint(i+previousId))
		if err != nil {
			fmt.Printf("Error: %v", err)
			return err
		}
		name := data["name"]
		fmt.Println(name)
	}

	client.Config.Next = client.Config.Previous
	if previousId < 20 {
		client.Config.Previous = client.Config.Previous[:40] + fmt.Sprint(1)
	} else {
		client.Config.Previous = client.Config.Previous[:40] + fmt.Sprint(previousId-20)
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
