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

//dsfsdf

func runCommand(command string, n int) {
	number := n - 1
	command = strings.ToLower(command)

	switch command {
	case "exit":
		fmt.Println("Program is exiting")
		os.Exit(0)
	case "inspectfork":
		status := <-forks[number].sender
		fmt.Println("test %v sdfsdf", status.timesUsed)
		forks[number].reciever <- print
	case "inspectphilosopher":
		status := <-philosophers[number].sender
		fmt.Printf("Philosopher with id %v has eaten %v times. Are they currently eating? %t\n", philosophers[number].id, status.timesEaten, status.isEating)
	case "help":
		fmt.Println("Available commands are \"exit\", \"inspectphilosopher n\", and \"inspectfork n\" where n is the id of a philosopher or fork in the range 1-5 inclusive")
	default:
		fmt.Println("Unknown command. Type \"help\" for help.")
	}
}
