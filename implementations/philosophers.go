package implementations

import (
	"fmt"
	"time"
)

// each philosopher must include two channels (one for input and one for output,
// both usable from the outside) through which it is possible to make queries on
// the state of the philosopher (number of times eaten, eating or thinking)

type Philosopher struct {
	id        int
	leftFork  *Fork
	rightFork *Fork

	timesEaten int
	isEating   bool

	reciever chan bool
	sender   chan bool
}

// Creates a new philosopher and returns a pointer to the address
// uses an Id, a pointer to the left fork and a pointer to the right fork

func NewPhilosopher(id int, leftFork *Fork, rightFork *Fork) *Philosopher {
	var p Philosopher
	p.id = id
	p.leftFork = leftFork
	p.rightFork = rightFork

	p.timesEaten = 0
	p.isEating = false

	p.reciever = make(chan bool, 2)
	p.sender = make(chan bool, 2)

	var newPhilosopher *Philosopher = &p

	return newPhilosopher
}

func Think() {
	//Afspil think af aretha franklin or some shit
}

func Eat(p Philosopher) {
	fmt.Println("He eatin' now")
	p.isEating = true
	time.Sleep(1 * time.Second)
	p.isEating = false
	p.timesEaten++
	fmt.Println("He no eat no mo")

}

func GetTimesEaten(p Philosopher) int {
	return p.timesEaten
}

func IsEating(p Philosopher) bool {
	return p.isEating
}
