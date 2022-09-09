package day14

import (
	"strconv"
	"strings"
)

// SumOfMaskedMemFields returns the sum of all set mem fields (not 0).
// A field is set with a value that filters through a bit mask and is then stored.
func SumOfMaskedMemFields(input []string) int {
	var mask string
	mem := make(map[string]int)
	for _, line := range input {
		values := strings.Split(line, " = ")
		if values[0] == "mask" {
			mask = values[1]
			continue
		}
		val, _ := strconv.Atoi(values[1])
		mem[values[0]] = compute(mask, val)
	}
	return sum(mem)
}

func compute(mask string, value int) int {
	for i, v := range mask {
		if v == 'X' {
			continue
		}
		if v == '1' {
			value |= (1 << (35 - i))
		}
		if v == '0' {
			bmask := ^(1 << (35 - i))
			value &= bmask
		}
	}
	return value
}

func sum(a map[string]int) int {
	var total int
	for _, v := range a {
		total += v
	}
	return total
}
