package main

import (
	"aoc2024/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	utils.ParseArgs(os.Args)
	if utils.Part == 1 {
		part1()
	} else {
		part2()
	}
}

func parseGrid() ([]string, utils.Point) {
	utils.Log("Parsing grid...")
	pos := utils.Point{}
	var grid []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		if start := strings.Index(line, "^"); start != -1 {
			pos = utils.Point{X: start, Y: len(grid)}
		}
		grid = append(grid, line)
	}
	utils.LogGrid(grid)
	utils.LogF("Guard position: %o\r\n", pos)

	return grid, pos
}

func part1() {
	grid, pos := parseGrid()

	utils.Log("Simulating guard movement...")
	var visited [][]bool
	for i := 0; i < len(grid); i++ {
		visited = append(visited, make([]bool, len(grid[i])))
	}
	visited[pos.Y][pos.X] = true
	numVisited := 1

	direction := 0
	dirChar := "^>v<"
	deltas := []utils.Point{{X: 0, Y: -1}, {X: 1, Y: 0}, {X: 0, Y: 1}, {X: -1, Y: 0}}
	for {
		nPos := utils.Point{X: pos.X + deltas[direction].X, Y: pos.Y + deltas[direction].Y}
		if !utils.IsInsideGrid(nPos, grid) {
			break
		}
		if grid[nPos.Y][nPos.X] != '#' {
			pos = nPos
			bytes := []byte(grid[pos.Y])
			bytes[pos.X] = dirChar[direction]
			grid[pos.Y] = string(bytes)
			if !visited[pos.Y][pos.X] {
				visited[pos.Y][pos.X] = true
				numVisited += 1
			}
			continue
		}
		direction = (direction + 1) % len(deltas)
		utils.LogF("Turn on %o\r\n", pos)
	}
	utils.LogGrid(grid)

	fmt.Printf("Visited cells: %d\r\n", numVisited)
}

func part2() {
}
