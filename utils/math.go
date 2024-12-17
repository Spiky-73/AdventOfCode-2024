package utils

import "math"

type Point struct {
	X int
	Y int
}

func (a *Point) Add(b *Point) Point {
	return Point{X: a.X + b.X, Y: a.Y + b.Y}
}

func (a *Point) Neg() Point {
	return Point{X: -a.X, Y: -a.Y}
}
func (a *Point) Sub(b *Point) Point {
	c := b.Neg()
	return a.Add(&c)
}

func Abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}

func Max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func IsInside(x int, y int, w int, h int) bool {
	return 0 <= x && x < w && 0 <= y && y < h
}

func Pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
