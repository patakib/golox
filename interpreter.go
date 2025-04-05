package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func RunFile(path string) {
	// read all bytes from file
	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	str := string(file)
	run(str)
}

func RunPrompt() {
	promptReader := bufio.NewReader(os.Stdin)
	fmt.Print(">> ") // Print prompt
	for {
		input, err := promptReader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		input = input[:len(input)-1] // Remove newline character
		if strings.Contains(input, "exit()") {
			fmt.Println("Exiting...")
			break
		}
		run(input)
		fmt.Print("GoŁoX >>> ") // Print prompt again
	}
}

func run(source string) {
	// tokenizing

	// for now just print the source code
	fmt.Println("Source code:", source)
}
