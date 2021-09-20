package main

import "fmt"

const count = 5

var forks = make([]*Fork, count)
var philosophers = make([]*Philosopher, count)

func main() {
	// Create forks and starts their go routine..
	for i := 0; i < count; i++ {
		forks[i] = NewFork(i)
		go forks[i].forkCycle()
	}

	// Create philosophers
	for i := 0; i < count; i++ {
		philosophers[i] = NewPhilosopher(i, forks[i], forks[(i+1)%count])
		go philosophers[i].philosopherCycle()
	}

	fmt.Println("Program started. Type \"help\" for available commands")

	//starts query
	for {
		query()
	}
}
