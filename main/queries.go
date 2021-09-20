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
			<-forks[number].sender
			forks[number].reciever <- print
		} else {
			fmt.Println("Fork number out of bounds.")
		}

	case "inspectphilosopher":
		if number < count && number >= 0 {
			philosophers[number].reciever <- true
			status := <-philosophers[number].sender
			fmt.Printf("Philosopher with id %v has eaten %v times. Are they currently eating? %t\n", philosophers[number].id+1, status.timesEaten, status.isEating)
		} else {
			fmt.Println("Philosopher number out of bounds.")
		}

	case "inspectall":
		for i := 0; i < count; i++ {
			philosophers[i].reciever <- true
			status := <-philosophers[i].sender
			fmt.Printf("Philosopher with id %v has eaten %v times. Are they currently eating? %t\n", philosophers[i].id+1, status.timesEaten, philosophers[i].status.isEating)
		}
		for i := 0; i < count; i++ {
			<-forks[i].sender
			forks[i].reciever <- print
		}
	case "help":
		fmt.Println("--------------------------------------------------------------------------")
		fmt.Println("Available commands are:")
		fmt.Println("    - \"exit\"                   - To exit the program")
		fmt.Println("    - \"inspectPhilosopher <n>\" - To inspect the philosopher with index n")
		fmt.Println("    - \"inspectFork <n>\"        - To inspect the fork with index n")
		fmt.Println("    - \"inspectAll\"             - To inspect all forks and philosophers")
		fmt.Println("n must be in the range 1-5 inclusive")
		fmt.Println("--------------------------------------------------------------------------")
	default:
		fmt.Println("Unknown command. Type \"help\" for help.")
	}
}
