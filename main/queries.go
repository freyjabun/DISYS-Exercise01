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
	case "inspectfork":
		forks[number].requestReceiver <- Request{
			who:   nil,
			event: print,
		}
	case "inspectphilosopher":
		philosophers[number].status <- true
	case "help":
		fmt.Println("Available commands are \"exit\", \"inspectphilosopher n\", and \"inspectfork n\" where n is the id of a philosopher or fork in the range 1-5 inclusive")
	default:
		fmt.Println("Unknown command. Type \"help\" for help.")
	}
}
