package main

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

	reciever chan string
	sender   chan string
}

func (p Philosopher) philosopherCycle() {
	for {
		select {
		// In case that a request is recieved from a fork, p writes a response
		case <-p.reciever:
			p.sender <- "Send som info i guess?"
		// In case that there is nothing send to the reciever, it does nothing
		default:
		}

		//This locks the entire table s.t. only this philosopher can acces the
		//forks. Therefore we can stop a deadlock??? i think?
		arbiter.Lock()

		leftfork := <-leftFork.sender
		rightFork := <-rightFork.sender

	}
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

	p.reciever = make(chan string, 2)
	p.sender = make(chan string)

	var newPhilosopher *Philosopher = &p

	return newPhilosopher
}

func Eat(p Philosopher) {
	p.sender <- "I wanna eat"
	<-p.reciever
	<-p.reciever
	p.isEating = true
	fmt.Println("Philosopher %d is eating", p.id)
	time.Sleep(2 * time.Second)
	p.isEating = false
	p.timesEaten++
	fmt.Println("Philosopher %d has stopped eating and has eaten %d", p.id, p.timesEaten)
}

func GetTimesEaten(p Philosopher) int {
	return p.timesEaten
}

func IsEating(p Philosopher) bool {
	return p.isEating
}
