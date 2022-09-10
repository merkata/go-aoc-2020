package day15

// SeenIndex maintains when a number has been seen.
type SeenIndex struct {
	Last             int
	MostRecentToLast int
}

// MemoryGame returns the number spoken on the nth turn, rules below.
// In this game, the players take turns saying numbers. They begin by taking turns reading from a list of starting numbers (your puzzle input). Then, each turn consists of considering the most recently spoken number:
// If that was the first time the number has been spoken, the current player says 0.
// Otherwise, the number had been spoken before; the current player announces how many turns apart the number is from when it was previously spoken.
func MemoryGame(input []int, turn int) int {
	seen := make(map[int]SeenIndex)
	var counter int
	var last int
	for _, v := range input {
		seen[v] = SeenIndex{Last: counter, MostRecentToLast: -1}
		last = v
		counter++
	}
	for counter < turn {
		var current int
		if seen[last].MostRecentToLast >= 0 {
			current = seen[last].Last - seen[last].MostRecentToLast
		}
		if _, ok := seen[current]; !ok {
			seen[current] = SeenIndex{Last: counter, MostRecentToLast: -1}
		} else {
			seen[current] = SeenIndex{MostRecentToLast: seen[current].Last, Last: counter}
		}
		last = current
		counter++
	}
	return last
}
