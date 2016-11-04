package main

import "fmt"

type robot struct {
	position  point
	direction Direction
	surface   table
}

func NewRobot() *robot {
	return &robot{
		point{0, 0},
		NORTH,
		table{
			point{0, 0},
			point{4, 4},
		},
	}
}

func (r *robot) Report() {
	fmt.Printf("%d,%d,%s\n", r.position.X(), r.position.Y(), r.direction.String())
}

func (r *robot) Move() {
	dx, dy := 0, 0
	switch r.direction {
	case NORTH:
		dy = 1
	case SOUTH:
		dy = -1
	case EAST:
		dx = 1
	case WEST:
		dx = -1
	}
	newPosition := r.position.Translate(dx, dy)
	if r.surface.withinBounds(newPosition) {
		r.position = newPosition
	}
}
