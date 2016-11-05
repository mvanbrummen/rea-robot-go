package main

import "fmt"

type point struct {
	X, Y int
}

func (p point) Translate(dx, dy int) point {
	return point{p.X + dx, p.Y + dy}
}

func (p point) String() string {
	return fmt.Sprintf("%d,%d", p.X, p.Y)
}
