package main

type point struct {
	x, y int
}

func (p point) X() int {
	return p.x
}

func (p point) Y() int {
	return p.y
}

func (p point) Translate(dx, dy int) point {
	return point{p.x + dx, p.y + dy}
}
