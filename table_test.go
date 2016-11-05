package main

import "testing"

type TableTest struct {
	x, y     int
	expected bool
}

func TestWithinBounds(t *testing.T) {
	max := point{4, 4}
	min := point{0, 0}
	table := table{min, max}
	var tests = [...]TableTest{
		{0, 0, true},
		{4, 4, true},
		{-1, 0, false},
		{0, -1, false},
		{4, 5, false},
		{5, 4, false},
	}
	for _, test := range tests {
		actual := table.withinBounds(point{test.x, test.y})
		if actual != test.expected {
			t.Error("For", test, "Expected", test.expected, "Got", actual)
		}
	}
}
