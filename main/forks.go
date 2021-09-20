package main

//Hello this is a fork...
type Fork struct {
	used  int
	id    int
	inUse bool

	reciever        chan string
	sender          chan bool
	requestReceiver chan Request
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

func NewFork(id int) *Fork {
	var fork Fork
	fork.used = 0
	fork.inUse = false
	fork.id = id

	var created *Fork = &fork

	fork.reciever = make(chan string, 2)
	fork.sender = make(chan bool)
	fork.requestReceiver = make(chan Request)

	return created
}

func JustUsed(fork Fork) {
	fork.used++
}

func IsInUse(fork Fork) bool {
	return fork.inUse
}

func TimesUsed(fork Fork) int {
	return fork.used
}

func (f *Fork) ForkCycle() {
	for {
		// The fork sends a message through the channel. This locks the fork until
		// a philosopher recieves the message. The philosopher will then write back
		// with an action, describing what the philosopher will do with the fork
		f.sender <- f.inUse
		action := <-f.reciever

		switch {

		}

		if action == "pick up" {
			f.inUse = true
			f.used++
		} else if action == "put down" {
			f.inUse = false
		}

	}
}
