package main

import "fmt"

var nullpoint = point{-1, -1}

type robot struct {
	position  point
	direction Direction
	surface   table
}

func NewRobot(surface table) *robot {
	return &robot{position: nullpoint, surface: surface}
}

func (r *robot) HasBeenPlaced() bool {
	return r.position != nullpoint
}

func (r *robot) Report() string {
	report := ""
	if r.HasBeenPlaced() {
		report = fmt.Sprintf("%d,%d,%s\n", r.position.X(), r.position.Y(), r.direction.String())
		fmt.Println(report)
	}
	return report
}

func (r *robot) Place(x int, y int, facing Direction) {
	newPosition := point{x, y}
	if r.surface.withinBounds(newPosition) {
		r.position = newPosition
		r.direction = facing
	}
}

func (r *robot) RotateLeft() Direction {
	if r.HasBeenPlaced() {
		r.direction = r.direction.RotateLeft()
	}
	return r.direction
}

func (r *robot) RotateRight() Direction {
	if r.HasBeenPlaced() {
		r.direction = r.direction.RotateRight()
	}
	return r.direction
}

func (r *robot) Move() {
	if r.HasBeenPlaced() {
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
}
