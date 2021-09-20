package main

import (
	"fmt"
	"time"
)

type Fork struct {
	id        int
	inUse     bool
	timesUsed int

	requestReceiver  chan Request
	philosopherQueue chan *Philosopher
}

var forkId = 0

func NewFork() Fork {
	var fork Fork

	fork.id = forkId
	forkId++
	fork.inUse = false
	fork.timesUsed = 0
	fork.requestReceiver = make(chan Request, 2)
	fork.philosopherQueue = make(chan *Philosopher, 2)

	return fork
}

const SLEEP_TIME = 50

func (fork *Fork) ForkCycle() {
	for {
		select {
		case request := <-fork.requestReceiver:
			switch request.event {
			case enqueue:
				fork.philosopherQueue <- request.who
				if !fork.inUse {
					fork.giveUp()
				}
			case drop:
				if fork.inUse {
					fork.inUse = false
					time.Sleep(time.Millisecond * SLEEP_TIME)
					if len(fork.philosopherQueue) >= 1 && !fork.inUse {
						fork.giveUp()
					}
				}
			case print:
				fork.Print()
			}
		}
	}
}

func (fork *Fork) Print() {
	fmt.Printf("Fork with id %v has been used %v times. Is it currently in use? %t\n", fork.id, fork.timesUsed, fork.inUse)
}

func (fork *Fork) giveUp() {
	philosopher := <-fork.philosopherQueue
	if philosopher == nil {
		return
	}
	fork.inUse = true
	fork.timesUsed++
	philosopher.pickUp <- fork
}

type Request struct {
	who   *Philosopher
	event Event
}

type Event int

const (
	enqueue Event = iota
	drop
	print
)
