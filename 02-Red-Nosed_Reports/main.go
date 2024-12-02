package main

import (
	"aoc2024/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func isSafe(report []int, skipped int) bool {
	previous := -1
	setup := false
	increasing := true
	for i := 0; i < len(report); i++ {
		if skipped == i {
			continue
		}
		n := report[i]
		if previous != -1 {
			if !setup {
				increasing = n > previous
				setup = true
			}
			if n == previous || (n > previous) != increasing || utils.Abs(n-previous) > 3 {
				return false
			}
		}
		previous = n
	}
	return true
}

func part1() {
	safeReport := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		report := make([]int, len(parts))
		for i := 0; i < len(parts); i++ {
			report[i], _ = strconv.Atoi(parts[i])
		}
		if isSafe(report, -1) {
			safeReport += 1
		}
	}
	fmt.Println(safeReport)

}

func part2() {
	safeReport := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		report := make([]int, len(parts))
		for i := 0; i < len(parts); i++ {
			report[i], _ = strconv.Atoi(parts[i])
		}
		for i := -1; i < len(report); i++ {
			if isSafe(report, i) {
				utils.LogF("%s: safe (skipped %d)\r\n", line, i)
				safeReport += 1
				break
			}
		}
	}
	fmt.Println(safeReport)
}
