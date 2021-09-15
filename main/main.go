package main

func main() {
	// Number of Philosophers and forks
	count := 5

	// Create forks
	forks := make([]*Fork, count)
	for i := 0; i < count; i++ {
		forks[i] = NewFork(i)
	}

	// Create philosophers
	philosophers := make([]*Philosopher, count)
	for i := 0; i < count; i++ {
		philosophers[i] = NewPhilosopher(i, forks[i], forks[(i+1)%count])

	}

	// Create goroutines for each philosopher and fork i guess?
	for i := 0; i < count; i++ {
		go forks[i].ForkCycle()
	}

	for i := 0; i < count; i++ {
		go philosophers[i].philosopherCycle()
	}

}

/*
func fork(chIn, chOut chan string) {
	for {
		<-chIn
		chOut <- "Eat with me"
		<-chIn
	}
}

func phil(chIn, chOut chan string) {
	counter := 0
	for {
		chOut <- "I wanna eat"
		<-chIn
		counter++
		fmt.Print("I ate %d",counter)
		chOut <- "I'm done"
	}
}
*/
