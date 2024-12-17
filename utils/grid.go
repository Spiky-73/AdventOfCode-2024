package utils

func LogGrid(grid []string) {
	for _, line := range grid {
		Log(line)
	}
}

func IsInsideGrid(pos Point, grid []string) bool {
	return IsInside(pos.X, pos.Y, len(grid[0]), len(grid))
}
