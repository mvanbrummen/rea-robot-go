package main

import "testing"

type PointTest struct {
	current, expected point
	dx, dy            int
}

func TestTranslate(t *testing.T) {
	var tests = [...]PointTest{
		{point{0, 0}, point{0, 1}, 0, 1},
		{point{4, 4}, point{3, 4}, -1, 0},
		{point{320, -55}, point{270, -45}, -50, 10},
		{point{2, 3}, point{1, 2}, -1, -1},
	}
	for _, test := range tests {
		new := test.current.Translate(test.dx, test.dy)
		if new != test.expected {
			t.Error("For", test.current, test.dx, test.dy, "Expected", test.expected, "Got", new)
		}
	}
}
