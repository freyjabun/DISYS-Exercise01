package implementations

func main() {
	// Number of Philosophers and forks
	count := 5

	// Create forks
	forks := make([]*Fork, count)
	for i := 0; i < count; i++ {
		forks[i] = Create(i)
	}

	// Create philosophers
	philosophers := make([]*Philosopher, count)
	for i := 0; i < count; i++ {
		philosophers[i] = NewPhilosopher(i, forks[i], forks[(i+1)%count])
	}
}
