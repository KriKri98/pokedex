package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	configuration := config{
		Next:     "https://pokeapi.co/api/v2/location-area/1",
		Previous: "https://pokeapi.co/api/v2/location-area/1",
	}
	getCommands()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleanedInput := cleanInput(input)
		commands := getCommands()
		_, exists := commands[cleanedInput[0]]
		if exists {
			err := commands[cleanedInput[0]].callback(&configuration)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
