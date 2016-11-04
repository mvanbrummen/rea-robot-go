package main

type Direction int

const (
	NORTH Direction = iota
	EAST
	SOUTH
	WEST
)

var directions = [...]string{
	"NORTH",
	"EAST",
	"SOUTH",
	"WEST",
}

func (d Direction) String() string {
	return directions[d]
}

func (d Direction) RotateRight() Direction {
	return (d + 1) % 4
}

func (d Direction) RotateLeft() Direction {
	return (d + 3) % 4
}
