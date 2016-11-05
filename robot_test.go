package main

import "testing"

func initRobot() *robot {
	table := table{point{0, 0}, point{4, 4}}
	return NewRobot(table)
}

type Placement struct {
	X, Y   int
	Facing Direction
}

func TestHasBeenPlaced(t *testing.T) {
	r := initRobot()
	if r.HasBeenPlaced() != false {
		t.Error("Expected HasBeenPlaced false got", r.HasBeenPlaced())
	}
	r.Place(-1, -1, NORTH)
	if r.HasBeenPlaced() != false {
		t.Error("Expected HasBeenPlaced false got", r.HasBeenPlaced())
	}
	r.Place(0, 0, NORTH)
	if r.HasBeenPlaced() != true {
		t.Error("Expected HasBeenPlaced false got", r.HasBeenPlaced())
	}
}

func TestPlace(t *testing.T) {
	r := initRobot()
	var tests = [...]Placement{
		{0, 0, NORTH},
		{4, 4, WEST},
		{1, 3, SOUTH},
		{2, 2, EAST},
	}
	for _, test := range tests {
		r.Place(test.X, test.Y, test.Facing)
		if r.direction != test.Facing || r.position.X != test.X || r.position.Y != test.Y {
			t.Error("For", test, "Failed")
		}
	}
}

type RotationTest struct {
	current, expected Direction
}

func TestRotateLeft(t *testing.T) {
	r := initRobot()
	var tests = [...]RotationTest{
		{NORTH, EAST},
		{EAST, SOUTH},
		{SOUTH, WEST},
		{WEST, NORTH},
	}
	for _, test := range tests {
		r.Place(0, 0, test.current)
		rotation := r.RotateRight()
		if rotation != test.expected {
			t.Error("For", test.current, "Expected", test.expected, "Got", rotation)
		}
	}
}

func TestRotateRight(t *testing.T) {
	r := initRobot()
	var tests = [...]RotationTest{
		{NORTH, WEST},
		{WEST, SOUTH},
		{SOUTH, EAST},
		{EAST, NORTH},
	}
	for _, test := range tests {
		r.Place(0, 0, test.current)
		rotation := r.RotateLeft()
		if rotation != test.expected {
			t.Error("For", test.current, "Expected", test.expected, "Got", rotation)
		}
	}
}

type ReportTest struct {
	position Placement
	expected string
}

func TestReport(t *testing.T) {
	r := initRobot()
	if r.Report() != "" {
		t.Error("For", "unplaced", "Expected", "", "Got", r.Report())
	}
	var tests = [...]ReportTest{
		{Placement{0, 0, NORTH}, "0,0,NORTH\n"},
		{Placement{4, 4, EAST}, "4,4,EAST\n"},
		{Placement{2, 2, SOUTH}, "2,2,SOUTH\n"},
		{Placement{3, 4, WEST}, "3,4,WEST\n"},
	}
	for _, test := range tests {
		r.Place(test.position.X, test.position.Y, test.position.Facing)
		if r.Report() != test.expected {
			t.Error("For", test.position, "Expected", test.expected, "Got", r.Report())
		}
	}
}

func TestMove(t *testing.T) {
	r := initRobot()
	r.Move()
	if r.Report() != "" {
		t.Error("Robot moved while not placed")
	}
	r.Place(0, 0, NORTH)
	r.Move()
	if r.Report() != "0,1,NORTH\n" {
		t.Error("Robot moved incorrectly")
	}
	r.RotateLeft()
	if r.Report() != "0,1,WEST\n" {
		t.Error("Robot moved incorrectly")
	}
	r.RotateRight()
	r.RotateRight()
	r.Move()
	r.Move()
	r.Move()
	if r.Report() != "3,1,EAST\n" {
		t.Error("Robot moved incorrectly")
	}
	r.Move()
	r.Move()
	r.Move()
	r.Move()
	if r.Report() != "4,1,EAST\n" {
		t.Error("Robot moved incorrectly")
	}
}
