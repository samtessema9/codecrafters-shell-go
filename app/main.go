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
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Errorf("Error reading input: %v", err)
		}
	
		fmt.Printf("%v: command not found\n", strings.TrimSpace(command))
	}
}
