package day1

import "strconv"

// CalculateExpense finds the entries that sum up to 2020
// and multiplies them, returning the result.
func CalculateExpense(input []string, entries int) int {
	for i, n := range input {
		for j := i; j < len(input); j++ {
			num1, err := strconv.Atoi(n)
			num2, err := strconv.Atoi(input[j])
			if err != nil {
				return 0
			}
			if entries > 2 {
				for k := j; k < len(input); k++ {
					num3, err := strconv.Atoi(input[k])
					if err != nil {
						return 0
					}
					if 2020-num1-num2 == num3 {
						return num1 * num2 * num3
					}
				}
			} else if 2020-num1 == num2 {
				return num1 * num2
			}
		}
	}
	return 0
}
