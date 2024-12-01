package main

import (
	"aoc2024/utils"
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
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

func part1() {
	var group1 []int
	var group2 []int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text() // Get the current line of text
		parts := strings.Split(line, "   ")
		i, _ := strconv.Atoi(parts[0])
		group1 = append(group1, i)
		i, _ = strconv.Atoi(parts[1])
		group2 = append(group2, i)
	}

	slices.Sort(group1)
	slices.Sort(group2)

	var distance int = 0
	for i := 0; i < len(group1); i++ {
		utils.Log(group1[i], group2[i])
		distance += int(math.Abs(float64(group1[i] - group2[i])))
	}

	fmt.Println(distance)
}

func part2() {
	var group1 []string
	counts := make(map[string]int)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text() // Get the current line of text
		parts := strings.Split(line, "   ")
		group1 = append(group1, parts[0])
		c := counts[parts[1]]
		counts[parts[1]] = c + 1
	}

	var similarity int = 0
	for i := 0; i < len(group1); i++ {
		c := counts[group1[i]]
		i, _ := strconv.Atoi(group1[i])
		utils.Log(i, c)
		similarity += i * c
	}

	fmt.Println(similarity)
}
