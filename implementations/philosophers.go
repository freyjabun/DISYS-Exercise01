package implementations

// each philosopher must include two channels (one for input and one for output,
// both usable from the outside) through which it is possible to make queries on
// the state of the philosopher (number of times eaten, eating or thinking)

type Philosopher struct {
	id          int
	leftForkId  int
	rightForkId int

	timesEaten int
	isEating   bool
}

func Think() {
	//Afspil think af aretha franklin or some shit
}

func (p Philosopher) Eat() {
	//Spis din weirdass mad, som skal spises med 2 gafler???
	p.timesEaten++
}
