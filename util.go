package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// GetUserInput prompts the user with a message and returns the input as a string
func GetUserInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt + "\n> ")
	input, _ := reader.ReadString('\n')
	input = strings.ToUpper(strings.TrimSpace(input))
	return input
}

// PrintStrings takes a slice of strings and prints each string on a new line
func PrintStrings(strings []string) {
	for _, str := range strings {
		fmt.Println(str)
	}
}
