package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/KriKri98/pokedex/internal/pokeapi"
)

func main() {
	client := pokeapi.NewClient(5 * time.Second)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleanedInput := cleanInput(input)
		commands := getCommands()
		_, exists := commands[cleanedInput[0]]
		if exists {
			err := commands[cleanedInput[0]].callback(&client)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
		scanErr := scanner.Err()
		if scanErr != nil {
			fmt.Println(scanErr)
		}
	}
}
