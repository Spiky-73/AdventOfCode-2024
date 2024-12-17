package main

import (
	"aoc2024/utils"
	"bufio"
	"fmt"
	"os"
)

func main() {
	utils.ParseArgs(os.Args)
	if utils.Part == 1 {
		part1()
	} else {
		part2()
	}
}

func parseAntennas() ([]string, map[byte][]utils.Point) {
	utils.Log("Parsing antennas...")
	antennas := make(map[byte][]utils.Point)
	scanner := bufio.NewScanner(os.Stdin)
	var grid []string
	for scanner.Scan() {
		line := scanner.Text()
		for x := 0; x < len(line); x++ {
			if line[x] != '.' {
				antennas[line[x]] = append(antennas[line[x]], utils.Point{X: x, Y: len(grid)})
			}
		}
		grid = append(grid, line)
	}
	utils.LogF("Grid size: %dx%d\r\n", len(grid[0]), len(grid))
	for freq, points := range antennas {
		utils.LogF("%c: %v\r\n", freq, points)

	}
	return grid, antennas
}

func part1() {
	grid, antennas := parseAntennas()

	utils.Log("Computing antinodes...")
	locations := make(map[utils.Point]struct{})
	for freq, points := range antennas {
		utils.LogF("%c: ", freq)
		for i := 0; i < len(points); i++ {
			for j := 0; j < i; j++ {
				ij := points[j].Sub(&points[i])
				utils.LogF("%v ", points[i].Sub(&ij))
				utils.LogF("%v ", points[j].Add(&ij))
				if utils.IsInsideGrid(points[i].Sub(&ij), grid) {
					locations[points[i].Sub(&ij)] = struct{}{}
				}
				if utils.IsInsideGrid(points[j].Add(&ij), grid) {
					locations[points[j].Add(&ij)] = struct{}{}
				}
			}
		}
		utils.Log()
	}

	fmt.Printf("Unique locations in bounds: %d\r\n", len(locations))
}

func part2() {
	grid, antennas := parseAntennas()

	utils.Log("Computing antinodes...")
	locations := make(map[utils.Point]struct{})
	for freq, points := range antennas {
		utils.LogF("%c: ", freq)
		for i := 0; i < len(points); i++ {
			for j := 0; j < i; j++ {
				ij := points[j].Sub(&points[i])
				for pos := points[i]; utils.IsInsideGrid(pos, grid); pos = pos.Sub(&ij) {
					locations[pos] = struct{}{}
					utils.LogF("%v ", pos)
				}
				for pos := points[j]; utils.IsInsideGrid(pos, grid); pos = pos.Add(&ij) {
					locations[pos] = struct{}{}
					utils.LogF("%v ", pos)
				}
			}
		}
		utils.Log()
	}

	fmt.Printf("Unique locations in bounds: %d\r\n", len(locations))
}
