package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:] // Get command line arguments
	if len(args) > 1 {
		fmt.Println("Usage: go run main.go [args]")
		return
	}
	if len(args) == 1 {
		RunFile(args[0])
	} else {
		RunPrompt()
	}
}
