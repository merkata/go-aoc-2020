package day13_test

import (
	"testing"

	"github.com/merkata/go-aoc-2020/day13"
)

type test struct {
	description string
	input       []string
	kind        string
	want        int
}

func TestGetIDTimesWaitTime(t *testing.T) {
	tests := []test{
		{
			description: "part 1 validation",
			input: []string{
				"939",
				"7,13,x,x,59,x,31,19",
			},
			kind: "part 1",
			want: 295,
		},
		{
			description: "part 1 solution",
			input: []string{
				"1000001",
				"29,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,41,x,x,x,x,x,x,x,x,x,577,x,x,x,x,x,x,x,x,x,x,x,x,13,17,x,x,x,x,19,x,x,x,23,x,x,x,x,x,x,x,601,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,37",
			},
			kind: "part 1",
			want: 174,
		},
	}
	for _, tc := range tests {
		got := day13.GetIDTimesWaitTime(tc.input)
		if tc.want != got {
			t.Errorf("failed test: %s, wanted %d got %d\n", tc.description, tc.want, got)
		}
	}
}
