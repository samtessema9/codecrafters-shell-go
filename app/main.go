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
			existsLocally := validCommands.Has(args[0])
			if existsLocally {
				fmt.Printf("%v is a shell builtin\n", args[0])
			} else {
				if path := investigatePath(args[0]); path != "" {
					fmt.Printf("%v is %v\n", args[0], path)
				} else {
					fmt.Printf("%v: not found\n", args[0])
				}
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

func parsePath() []string {
	path := os.Getenv("PATH")
	if path == "" {
		fmt.Errorf("PATH not found!")
		os.Exit(0)
	}

	paths := strings.Split(path, ":")
	
	return paths
}

func investigatePath(cmd string) string {
	paths := parsePath()

	for _, path := range paths {
		filename := fmt.Sprintf("%v/%v", path, cmd)

		_, err := os.Stat(filename)

		if os.IsNotExist(err) {
			continue
		} else {
			return filename
		}
	}

	return ""
}