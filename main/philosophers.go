package main

import (
	"fmt"
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

	status PhilosopherStatus

	sender   chan PhilosopherStatus
	reciever chan bool
}

type PhilosopherStatus struct {
	timesEaten int
	isEating   bool
}

var arbiter sync.Mutex

func (p Philosopher) philosopherCycle() {
	for {
		select {
		// if query want to know about a philosopher we return philosopher info
		case <-p.reciever:
			p.sender <- p.status
		}

		//This locks the entire table s.t. only this philosopher can acces the
		//forks.
		arbiter.Lock()

		//creating action
		action := doNothing

		//these two channels will tell the philosopher whether the forks are
		//available or not... i think?--
		LeftForkStatus := <-p.leftFork.sender
		RightForkStatus := <-p.rightFork.sender

		willPickUp := !LeftForkStatus.inUse && !RightForkStatus.inUse

		if willPickUp {
			action = pickUp
		}

		p.leftFork.reciever <- action
		p.leftFork.reciever <- action

		if willPickUp {
			p.status.isEating = true
			p.status.timesEaten++
		}

		fmt.Println("i am eating nom nom")

		arbiter.Unlock()

		time.Sleep(2 * time.Second)

		if p.status.isEating {
			arbiter.Lock()
			<-p.leftFork.sender
			<-p.rightFork.sender
			p.leftFork.reciever <- putDown
			p.leftFork.reciever <- putDown
			p.status.isEating = false
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

	p.status.timesEaten = 0
	p.status.isEating = false

	p.sender = make(chan PhilosopherStatus, 2)
	p.reciever = make(chan bool)

	var newPhilosopher *Philosopher = &p

	return newPhilosopher
}
