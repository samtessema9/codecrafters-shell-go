package main

import (
	"bufio"
	"fmt"
	"os"

	"strings"

	"github.com/scylladb/go-set"
	"github.com/scylladb/go-set/strset"
)

var validCommands *strset.Set = set.NewStringSet(
	"echo",
	"exit",
	"type",
)

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

		switch (command) {
		case "exit":
			os.Exit(0)
		case "echo": 
			echoString := strings.Join(args, " ")
			fmt.Println(echoString)
		case "type":
			exists := validCommands.Has(args[0])
			if exists {
				fmt.Printf("%v is a shell builtin", args[0])
			} else {
				fmt.Printf("%v: not found", args[0])
			}
		default:
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