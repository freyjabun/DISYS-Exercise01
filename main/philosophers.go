package main

import (
	"sync"
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

var arbiter sync.Mutex

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

		//reseting some variables and creating action
		p.isEating = false
		action := "nothing"
		
		//these two channels will tell the philosopher whether the forks are 
		//available or not... i think?
		isLeftForkInUse := <-p.leftFork.sender
		isRightForkInUse := <-p.rightFork.sender

		willPickUp := !isLeftForkInUse && !isRightForkInUse

		if willPickUp{
			action = "pick up"
		}

		p.leftFork.reciever <- action
		p.leftFork.reciever <- action

		if willPickUp {
			p.isEating = true
			p.timesEaten++
		}

		arbiter.Unlock()

		time.Sleep(2*time.Second)

		if p.isEating {
			arbiter.Lock()
			<- p.leftFork.sender
			<- p.rightFork.sender
			p.leftFork.reciever <- "put down"
			p.leftFork.reciever <- "put down"
			arbiter.Unlock()
		}
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

func GetTimesEaten(p Philosopher) int {
	return p.timesEaten
}

func IsEating(p Philosopher) bool {
	return p.isEating
}
