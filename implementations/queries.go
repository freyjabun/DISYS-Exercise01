package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func query() {
	var command string
	var number int

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%s %d", &command, &number)

	runCommand(command, number)
}

func runCommand(command string, number int) {
	command = strings.ToLower(command)

	switch command {
	case "exit":
		fmt.Println("Program is exiting")
		os.Exit(0)
	case "yo":
		fmt.Println("yo whatup")
	default:
		fmt.Println("Unknown command. Type \"help\" for help.")
	}
}

func main() {
	for {
		query()
	}
}
