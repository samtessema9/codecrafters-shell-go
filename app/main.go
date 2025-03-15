package main

import (
	"bufio"
	"fmt"
	"os"

	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	for {
		// Prompt User
		fmt.Fprint(os.Stdout, "$ ")
	
		// Wait for user input
		userInput, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Errorf("Error reading input: %v", err)
		}
		command, args := parseInput(userInput)

		if command == "echo" {
			echoString := strings.Join(args, " ")
			fmt.Println(echoString)
		} else if command == "exit" {
			os.Exit(0)
		} else {
			fmt.Printf("%v: command not found\n", strings.TrimSpace(command))
		} 
	}
}

func parseInput(input string) (string, []string) {
	strs := strings.Split(input, " ")
	
	// Strip the \n from the end of the last arg
	strs[len(strs) - 1] = strings.TrimSpace(strs[len(strs) - 1])

	if len(strs) > 1 {
		return strs[0], strs[1:]
	}

	return strs[0], []string{}
}