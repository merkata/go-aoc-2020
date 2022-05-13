package day5

import (
	"sort"
)

// PlaneSeat describes a seat in terms of low and high values,
// converging to one specific value in the end.
type PlaneSeat struct {
	Low  int
	High int
}

// FindSeatID determines seat IDs based on binary space partitioning.
// Once the seat is found (row and column), the ID is computed as
// row * 8 + column. FindSeatID can find the highest ID and a missing one.
func FindSeatID(input []string, kind string) int {
	var seats []int
	if len(input) <= 0 {
		return 0
	}
	for _, seat := range input {
		seats = append(seats, compute(seat))
	}
	sort.Ints(seats)
	switch kind {
	case "missing":
		for i := range seats {
			if seats[i+1]-seats[i] != 1 {
				return seats[i] + 1
			}
		}
		return 0
	default:
		return seats[len(seats)-1]
	}
}

func compute(a string) int {
	row := PlaneSeat{High: 127}
	column := PlaneSeat{High: 7}
	for _, pos := range a {
		switch pos {
		case 'B':
			row.Low += (row.High-row.Low)/2 + 1
		case 'F':
			row.High -= (row.High-row.Low)/2 + 1
		case 'R':
			column.Low += (column.High-column.Low)/2 + 1
		case 'L':
			column.High -= (column.High-column.Low)/2 + 1
		}
	}
	return row.High*8 + column.High
}
