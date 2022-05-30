package day9

import "strconv"

// XMASSumFail slides through the input with a preable window
// and checks for the property that every next number can be
// expressed as the sum of any two numbers in the window.
func XMASSumFail(input []string, preamble int, kind string) int {
	present := make(map[int]bool)
	window := make([]int, len(input))
	for idx, val := range input {
		num, _ := strconv.Atoi(val)
		window[idx] = num
	}
	for i := 0; i < preamble; i++ {
		present[window[i]] = true
	}
	forget := 0
	for j := preamble; j < len(window); j++ {
		sums := false
		for k := forget + 1; k < j; k++ {
			if present[window[j]-window[k]] {
				sums = true
			}
		}
		if !sums {
			if kind == "failed" {
				return window[j]
			}
			return minmaxInContiguousSumFor(window, window[j])
		}
		present[window[forget]] = false
		forget++
		present[window[j]] = true
	}
	return 0
}

func minmaxInContiguousSumFor(window []int, num int) int {
	for i := 0; i < len(window)-1; i++ {
		sum := window[i]
		smallest, largest := window[i], window[i]
		for j := i + 1; j < len(window); j++ {
			sum += window[j]
			if window[j] < smallest {
				smallest = window[j]
			}
			if window[j] > largest {
				largest = window[j]
			}
			if sum == num {
				return smallest + largest
			}
			if sum > num {
				break
			}
		}
	}
	return 0
}
