package day10

import (
	"sort"
	"strconv"
)

// AdapterJolts tries to form a chain of all adapters present
// and return the result of multiplying the 1-difference to
// the 3-difference jolts.
func AdapterJolts(input []string, kind string) int {
	nums := make([]int, len(input))
	for i, strnum := range input {
		num, _ := strconv.Atoi(strnum)
		nums[i] = num
	}
	sort.Ints(nums)
	current := 0
	ones := 0
	threes := 0
	for _, n := range nums {
		switch n - current {
		case 1:
			ones++
		case 3:
			threes++
		}

		current = n
	}
	threes++ // your adapter
	if kind == "jolts" {
		return ones * threes
	}
	return 0
}
