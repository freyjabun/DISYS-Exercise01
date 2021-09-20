package main

import "fmt"

var AMOUNT_OF_FORKS = 5
var AMOUNT_OF_PHILOSOPHERS = 5
var forks = make([]Fork, AMOUNT_OF_FORKS)
var philosophers = make([]Philosopher, AMOUNT_OF_PHILOSOPHERS)

func main() {
	for i := 0; i < AMOUNT_OF_FORKS; i++ {
		forks[i] = NewFork()
		go forks[i].ForkCycle()
	}

	for i := 0; i < AMOUNT_OF_PHILOSOPHERS; i++ {
		var leftForkId = i - 1
		if leftForkId < 0 {
			leftForkId = AMOUNT_OF_FORKS - 1
		}
		var rightForkId = i
		philosophers[i] = NewPhilosopher(&forks[leftForkId], &forks[rightForkId])
		go philosophers[i].PhilosopherCycle()
	}

	fmt.Println("Program started. Type \"help\" for available commands")

	for {
		query()
	}
}
