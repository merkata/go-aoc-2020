package day16

import (
	"strconv"
	"strings"
	"unicode"
)

type TicketField struct {
	LowVals  [2]int
	HighVals [2]int
}

// TicketErrorScanningRate scans nearby tickets for validity.
// The Error Scanning Rate is expressed as the sum of incorrect values
// for every invalid ticket.
func TicketErrorScanningRate(input []string) int {
	ticketfields := make(map[string]TicketField)
	var nearby []string
	for i, line := range input {
		if strings.Contains(line, "-") {
			fields := strings.Split(line, ": ")
			title := fields[0]
			values := strings.FieldsFunc(fields[1], func(c rune) bool { return !unicode.IsLetter(c) && !unicode.IsNumber(c) })
			ticketfields[title] = prepTicketFieldFrom(values)
		} else if strings.Contains(line, ",") {
			if input[i-1] != "your ticket:" {
				nearby = append(nearby, line)
			}
		}
	}
	var missing []int
	for _, ticket := range nearby {
		if v, ok := isValid(ticket, ticketfields); !ok {
			missing = append(missing, v)
		}
	}
	return sum(missing)
}

func isValid(a string, ticket map[string]TicketField) (int, bool) {
	for _, v := range strings.Split(a, ",") {
		n, _ := strconv.Atoi(v)
		anyValid := false
		for _, field := range ticket {
			if (n >= field.LowVals[0] && n <= field.LowVals[1]) || (n >= field.HighVals[0] && n <= field.HighVals[1]) {
				anyValid = true
				break
			}
		}
		if !anyValid {
			return n, false
		}
	}
	return 0, true
}

func sum(a []int) int {
	var total int
	for _, v := range a {
		total += v
	}
	return total
}

func prepTicketFieldFrom(fields []string) TicketField {
	v1, _ := strconv.Atoi(fields[0])
	v2, _ := strconv.Atoi(fields[1])
	v3, _ := strconv.Atoi(fields[3])
	v4, _ := strconv.Atoi(fields[4])
	return TicketField{LowVals: [2]int{v1, v2}, HighVals: [2]int{v3, v4}}
}
