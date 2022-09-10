package day15_test

import (
	"testing"

	"github.com/merkata/go-aoc-2020/day15"
)

type test struct {
	description string
	input       []int
	turn        int
	want        int
}

func TestSumOfMaskedMemFields(t *testing.T) {
	tests := []test{
		{
			description: "part 1 validation",
			input:       []int{0, 3, 6},
			turn:        2020,
			want:        436,
		},
		{
			description: "part 1 solution",
			input:       []int{0, 1, 4, 13, 15, 12, 16},
			turn:        2020,
			want:        1665,
		},
		{
			description: "part 2 validation",
			input:       []int{0, 3, 6},
			turn:        30000000,
			want:        175594,
		},
		{
			description: "part 2 solution",
			input:       []int{0, 1, 4, 13, 15, 12, 16},
			turn:        30000000,
			want:        16439,
		},
	}
	for _, tc := range tests {
		got := day15.MemoryGame(tc.input, tc.turn)
		if tc.want != got {
			t.Errorf("failed test: %s, wanted %d got %d\n", tc.description, tc.want, got)
		}
	}
}
