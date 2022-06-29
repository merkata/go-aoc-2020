package day11

import "fmt"

//OccupiedSeats runs a simulation until it stabilizes
//and returns the number of occupied seats when no
//changes are done in a subsequent round accourdingly.
func OccupiedSeats(input []string) int {
	layout := NewLayoutFrom(input)
	result := layout.CalculateOccupied()
	for {
		diff := layout.CalculateOccupied()
		fmt.Printf("Iterating: previous %d current %d\n", result, diff)
		if diff != result {
			result = diff
		} else {
			break
		}
	}
	return result
}

//Layout maps to a dimension of seats.
type Layout struct {
	Seats [][]Seat
}

//RowCol is a mapping of a seat on the plane axes.
type RowCol [2]int

//CalculateOccupied runs through the layout and
//returns how many seats are occupied in that run.
func (l *Layout) CalculateOccupied() int {
	occupy := []RowCol{}
	free := []RowCol{}
	for i, col := range l.Seats {
		for j, row := range col {
			occupied := l.OccupiedSeats(i, j)
			switch row {
			case 'L':
				if occupied == 0 {
					occupy = append(occupy, RowCol{i, j})
				}
			case '#':
				if occupied >= 4 {
					free = append(free, RowCol{i, j})
				}
			}
		}
	}
	for _, seat := range occupy {
		l.Seats[seat[0]][seat[1]] = '#'
	}
	for _, seat := range free {
		l.Seats[seat[0]][seat[1]] = 'L'
	}
	occupied := 0
	for _, isle := range l.Seats {
		for _, seat := range isle {
			if seat == '#' {
				occupied++
			}
		}
	}
	return occupied
}

//OccupiedSeats counts how many occupied seats
//there are adjecent to the seat at Row{i, j}.
func (l *Layout) OccupiedSeats(i, j int) int {
	occupied := 0
	if i >= 1 {
		if j >= 1 {
			if l.Seats[i-1][j-1] == '#' {
				occupied++
			}
		}
		if l.Seats[i-1][j] == '#' {
			occupied++
		}
		if j+1 < len(l.Seats[i]) {
			if l.Seats[i-1][j+1] == '#' {
				occupied++
			}
		}
	}
	if i+1 < len(l.Seats) {
		if j >= 1 {
			if l.Seats[i+1][j-1] == '#' {
				occupied++
			}
		}
		if l.Seats[i+1][j] == '#' {
			occupied++
		}
		if j+1 < len(l.Seats[i]) {
			if l.Seats[i+1][j+1] == '#' {
				occupied++
			}
		}
	}
	if j >= 1 {
		if l.Seats[i][j-1] == '#' {
			occupied++
		}
	}
	if j+1 < len(l.Seats[i]) {
		if l.Seats[i][j+1] == '#' {
			occupied++
		}
	}
	return occupied
}

//Seat is either:
//.: floor, L: empty, #: taken.
type Seat rune

//NewLayoutFrom takes a slice of strings and places
//them inside a Layout struct.
func NewLayoutFrom(input []string) *Layout {
	allSeats := make([][]Seat, len(input))
	for i, line := range input {
		seatsLine := make([]Seat, len(line))
		for j, ch := range line {
			seatsLine[j] = Seat(ch)
		}
		allSeats[i] = seatsLine
	}
	return &Layout{Seats: allSeats}
}
