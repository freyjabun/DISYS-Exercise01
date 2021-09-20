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
		if number < count && number >= 0 {
			status := <-forks[number].sender
			fmt.Println("test %v sdfsdf", status.timesUsed)
			forks[number].reciever <- print
		} else {
			fmt.Println("Fork number out of bounds.")
		}

	case "inspectphilosopher":
		if number < count && number >= 0 {
			philosophers[number].reciever <- true
			status := <-philosophers[number].sender
			fmt.Printf("Philosopher with id %v has eaten %v times. Are they currently eating? %t\n", philosophers[number].id, status.timesEaten, status.isEating)
		} else {
			fmt.Println("Philosopher number out of bounds.")
		}
	case "help":
		fmt.Println("Available commands are \"exit\", \"inspectphilosopher n\", and \"inspectfork n\" where n is the id of a philosopher or fork in the range 1-5 inclusive")
	default:
		fmt.Println("Unknown command. Type \"help\" for help.")
	}
}
