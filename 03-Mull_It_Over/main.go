package main

import (
	"aoc2024/utils"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	utils.ParseArgs(os.Args)
	if utils.Part == 1 {
		part1()
	} else {
		part2()
	}
}

func part1() {
	regex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	total := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		matches := regex.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			left, _ := strconv.Atoi(match[1])
			right, _ := strconv.Atoi(match[2])
			total += left * right
		}
		utils.Log(matches)
	}
	fmt.Printf("%d\r\n", total)
}

func part2() {
	regex := regexp.MustCompile(`(do)\(\)|(don't)\(\)|(mul)\((\d{1,3}),(\d{1,3})\)`)
	total := 0
	scanner := bufio.NewScanner(os.Stdin)
	enabled := true
	for scanner.Scan() {
		line := scanner.Text()
		matches := regex.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			utils.Log(match)
			if match[1] == "do" {
				enabled = true
			} else if match[2] == "don't" {
				enabled = false

			} else if enabled {
				left, _ := strconv.Atoi(match[4])
				right, _ := strconv.Atoi(match[5])
				total += left * right
			}
		}
	}
	fmt.Printf("%d\r\n", total)
}
