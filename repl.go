package main

import "strings"

func cleanInput(text string) []string {
	formattedString := strings.ToLower(text)

	output := []string{}
	ok := true
	before := ""
	for {
		formattedString = strings.Trim(formattedString, " ")
		before, formattedString, ok = strings.Cut(formattedString, " ")
		if !ok {
			output = append(output, before)
			break
		}
		output = append(output, before)
	}

	return output
}
