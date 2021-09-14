package implementations

//Hello this is a fork...
type Fork struct {
	used  int
	id    int
	inUse bool
}

<<<<<<< HEAD
func NewFork(id int) *Fork {
=======
func Create(id int) *Fork {
>>>>>>> 83c44878ed50dba49bb5fb3a388e5c7e7abd7d4e
	var fork Fork
	fork.used = 0
	fork.inUse = false
	fork.id = id

	var created *Fork = &fork

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
