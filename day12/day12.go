package day12

import (
	"strconv"
)

//Ship encapsulates the movement over a coordinate map of a ship.
type Ship struct {
	Facing    string
	Distances map[string]int
	Coords    map[string]int
	InvCoords map[int]string
}

//ManhattanDistance calculates east/west and north/south sum of a ship.
func ManhattanDistance(input []string) int {
	coords := map[string]int{"E": 0, "N": 3, "W": 2, "S": 1}
	invcoords := map[int]string{0: "E", 3: "N", 2: "W", 1: "S"}
	distances := make(map[string]int)
	ship := &Ship{Facing: "E", Coords: coords, InvCoords: invcoords, Distances: distances}
	for _, line := range input {
		move := string(line[0])
		distance, _ := strconv.Atoi(line[1:])
		ship.Advance(move, distance)
	}
	return ship.CalculateDistance()
}

//Advance updates the ship on the coordinates map.
func (s *Ship) Advance(move string, distance int) {
	switch move {
	case "F":
		move = s.Facing
		fallthrough
	case "N", "S", "E", "W":
		if distance-s.Distances[s.InvCoords[(s.Coords[move]+2)%4]] > 0 {
			s.Distances[move] += distance - s.Distances[s.InvCoords[(s.Coords[move]+2)%4]]
			s.Distances[s.InvCoords[(s.Coords[move]+2)%4]] = 0
		} else {
			s.Distances[s.InvCoords[(s.Coords[move]+2)%4]] -= distance
		}
	case "L":
		s.Facing = s.InvCoords[(s.Coords[s.Facing]-distance/90+4)%4]
	case "R":
		s.Facing = s.InvCoords[(s.Coords[s.Facing]+(distance/90))%4]
	}
}

//CalculateDistance returns the Manhattan distance according to the coords.
func (s *Ship) CalculateDistance() int {
	return s.Distances["E"] + s.Distances["W"] + s.Distances["N"] + s.Distances["S"]
}
