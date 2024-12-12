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

func countXMAS(s string) int {
	return strings.Count(s, "XMAS") + strings.Count(s, "SAMX")
}

func part1() {
	count := 0
	scanner := bufio.NewScanner(os.Stdin)
	var grid []string

	utils.Log("----- HORIZONTAL -----")
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, line)
		c := countXMAS(line)
		utils.Log(line, c)
		count += c
	}

	utils.Log("----- VERTICAL -----")
	for i := 0; i < len(grid); i++ {
		var chars []byte
		for _, line := range grid {
			chars = append(chars, line[i])
		}
		line := string(chars)
		c := countXMAS(line)
		utils.Log(line, c)
		count += c
	}

	utils.Log("----- DOWN+RIGHT -----")
	for i := -len(grid) + 1; i < len(grid); i++ {
		var chars []byte
		for j := utils.Max(-i, 0); j < len(grid) && i+j < len(grid); j++ {
			chars = append(chars, grid[j][j+i])
		}
		line := string(chars)
		c := countXMAS(line)
		utils.Log(line, c)
		count += c
	}
	utils.Log("----- UP+RIGHT -----")
	for i := -len(grid) + 1; i < len(grid); i++ {
		var chars []byte
		for j := utils.Max(-i, 0); j < len(grid) && i+j < len(grid); j++ {
			chars = append(chars, grid[len(grid)-1-j][j+i])
		}
		line := string(chars)
		c := countXMAS(line)
		utils.Log(line, c)
		count += c
	}
	fmt.Printf("%d\r\n", count)
}

func isMAS(s string) bool {
	return s == "SM" || s == "MS"
}

func part2() {
	count := 0
	scanner := bufio.NewScanner(os.Stdin)
	var grid []string

	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, line)
	}
	for j := 1; j < len(grid)-1; j++ {
		for i := 1; i < len(grid)-1; i++ {
			if grid[j][i] == 'A' {
				down := string(grid[j-1][i-1]) + string(grid[j+1][i+1])
				up := string(grid[j+1][i-1]) + string(grid[j-1][i+1])
				if isMAS(down) && isMAS(up) {
					count += 1
					utils.Log(i, j)
				}
			}
		}
	}
	fmt.Printf("%d\r\n", count)

}
