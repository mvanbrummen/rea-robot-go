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
