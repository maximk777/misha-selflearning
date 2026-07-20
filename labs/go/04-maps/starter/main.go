package starter

func Lookup(scores map[string]int, name string) (int, bool) {
	value, ok := scores[name]
	return value, ok
}

func NewScores() map[string]int {
	return make(map[string]int)
}
