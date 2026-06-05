package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
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
			err := commands[cleanedInput[0]].callback()
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
