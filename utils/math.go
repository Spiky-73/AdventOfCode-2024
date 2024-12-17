package utils

import "math"

type Point struct {
	X int
	Y int
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
