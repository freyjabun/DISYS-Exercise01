package main

import (
	"fmt"
)

type Fork struct {
	id       int
	status   ForkStatus
	reciever chan Event
	sender   chan ForkStatus
}

type ForkStatus struct {
	inUse     bool
	timesUsed int
}

func NewFork(id int) *Fork {
	var fork Fork
	fork.status = ForkStatus{inUse: false, timesUsed: 0}
	fork.id = id

	var created *Fork = &fork

	fork.reciever = make(chan Event, 2)
	fork.sender = make(chan ForkStatus)

	return created
}

func (f *Fork) forkCycle() {
	for {
		// The fork sends a message through the channel. This locks the fork until
		// a philosopher recieves the message. The philosopher will then write back
		// with an action, describing what the philosopher will do with the fork--
		f.sender <- f.status
		event := <-f.reciever

		switch event {
		case pickUp:
			f.status.inUse = true
			f.status.timesUsed++
		case putDown:
			f.status.inUse = false
		case print:
			f.Print()
		}

	}
}

func (fork *Fork) Print() {
	fmt.Printf("Fork with id %v has been used %v times. Is it currently in use? %t\n", fork.id, fork.status.timesUsed, fork.status.inUse)
}

type Event int

const (
	pickUp Event = iota
	putDown
	doNothing
	print
)
