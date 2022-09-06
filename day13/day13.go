package day13

import (
	"fmt"
	"strconv"
	"strings"
)

// GetIDTimesWaitTime calculates the earliest Bus our hero can take.
// Substracting the time of the earliest Bus from the earlies leave time
// gets you how many minutes our hero will wait.
// Multiply that by the Bus ID and you get the result wanted.
func GetIDTimesWaitTime(input []string) int {
	earliestTime, _ := strconv.Atoi(input[0])
	waitTime := earliestTime
	fmt.Printf("earliest time %d\n", earliestTime)
	var bus int
	for _, busID := range strings.Split(input[1], ",") {
		fmt.Printf("looking at Bus %s\n", busID)
		if busID == "x" {
			continue
		}
		id, _ := strconv.Atoi(busID)
		distance := ((earliestTime / id) * id) + id
		fmt.Printf("calculated time %d for Bus %s\n", distance, busID)
		if distance-earliestTime <= waitTime {
			waitTime = distance - earliestTime
			bus = id
		}
	}

	return bus * waitTime
}
