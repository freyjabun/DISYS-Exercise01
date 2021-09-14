package implementations

//Hello this is a fork...
type Fork struct {
	used  int
	id    int
	inUse bool
}

func Create(id int) *Fork {
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
