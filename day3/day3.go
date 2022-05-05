package day3

// Tree denotes the character used for a tree in the input.
const Tree = '#'

// Slope describes the path taken or a slope through a track.
type Slope struct {
	Right int
	Down  int
}

// SlopeTrees counts the trees encountered on a slope.
func SlopeTrees(input []string, slopes []Slope) int {
	trees := make([]int, len(slopes))
	for i, line := range input {
		for j, slope := range slopes {
			if i%slope.Down == 0 {
				if line[((i*slope.Right)/slope.Down)%len(line)] == Tree {
					trees[j]++
				}
			}
		}
	}
	return multiply(trees)
}

func multiply(a []int) int {
	product := 1
	for _, b := range a {
		product *= b
	}
	return product
}
