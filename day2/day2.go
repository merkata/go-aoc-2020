package day2

import (
	"strconv"
	"strings"
)

// ValidPasswords takes a list of password policies and passwords and
// determines how many passwords are valid from the supplied list.
func ValidPasswords(input []string, check string) int {
	var valid int
	for _, line := range input {
		chunks := strings.Split(line, " ")
		marker := strings.Split(chunks[0], "-")
		left, _ := strconv.Atoi(marker[0])
		right, _ := strconv.Atoi(marker[1])
		letter := string(chunks[1][0])
		switch check {
		case "boundary":
			policy := strings.Count(chunks[2], letter)
			if policy >= left && policy <= right {
				valid++
			}
		case "index":
			policyLeft := string(chunks[2][left-1]) == letter
			policyRight := string(chunks[2][right-1]) == letter
			if policyLeft != policyRight {
				valid++
			}
		}
	}
	return valid
}
