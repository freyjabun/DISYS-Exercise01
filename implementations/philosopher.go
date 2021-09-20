package main

import (
	"fmt"
)

type Philosopher struct {
	id         int
	leftFork   *Fork
	rightFork  *Fork
	timesEaten int
	isEating   bool
	pickUp     chan *Fork
	status     chan bool
}

var philosopherId = 0

func NewPhilosopher(leftFork *Fork, rightFork *Fork) Philosopher {
	var philosopher Philosopher

	philosopher.id = philosopherId
	philosopherId++
	philosopher.leftFork = leftFork
	philosopher.rightFork = rightFork
	philosopher.timesEaten = 0
	philosopher.isEating = false
	philosopher.pickUp = make(chan *Fork, 2)
	philosopher.status = make(chan bool)

	return philosopher
}

func (philosopher *Philosopher) PhilosopherCycle() {
	for {
		select {
		case <-philosopher.status:
			fmt.Printf("Philosopher with id %v has eaten %v times. Are they currently eating? %t\n", philosopher.id, philosopher.timesEaten, philosopher.isEating)
		case fork := <-philosopher.pickUp:
			var id = fork.id
			//pick up logic
		}
	}
}

func (philosopher *Philosopher) eat() {
	if philosopher.isEating {
		return
	}
	philosopher.timesEaten++
	philosopher.isEating = true
	philosopher.think()
}

func (philosopher *Philosopher) think() {
	if !philosopher.isEating {
		return
	}
	philosopher.isEating = false

	philosopher.leftFork.requestReceiver <- Request{
		who:   philosopher,
		event: drop,
	}

	philosopher.rightFork.requestReceiver <- Request{
		who:   philosopher,
		event: drop,
	}

}

func (philosopher *Philosopher) reserveForks() {
	if philosopher.isEating {
		return
	}
	if !philosopher.leftFork.inHand {
		philosopher.leftFork.fork.input <- Request{
			who:   philosopher,
			event: enqueue,
		}
	}
	if !philosopher.rightFork.inHand {
		philosopher.rightFork.fork.input <- Request{
			who:   philosopher,
			event: enqueue,
		}
	}
}
