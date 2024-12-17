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

func part1() {
	utils.Log("Parsing equations...")
	scanner := bufio.NewScanner(os.Stdin)

	calibration := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		target, _ := strconv.Atoi(parts[0][:len(parts[0])-1])
		numbers := make([]int, len(parts)-1)
		for i := 0; i < len(parts)-1; i++ {
			numbers[i], _ = strconv.Atoi(parts[i+1])
		}
		utils.LogF("%d with %v: ", target, numbers)
		for operators := 0; operators < (1 << (len(numbers) - 1)); operators++ {
			current := numbers[0]
			for o := 0; o < len(numbers)-1; o++ {
				if operators&(1<<o) == 0 {
					current += numbers[o+1]
				} else {
					current *= numbers[o+1]
				}

			}
			if current == target {
				utils.LogF(fmt.Sprintf("%%0%db\r\n", len(numbers)-1), operators)
				calibration += target
				goto SUCCESS
			}
		}
		utils.Log("no solution")
	SUCCESS:
	}
	fmt.Printf("Calibration result: %d\r\n", calibration)
}

func part2() {
	utils.Log("Parsing equations...")
	scanner := bufio.NewScanner(os.Stdin)
	operators := "+*|"

	calibration := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		target, _ := strconv.Atoi(parts[0][:len(parts[0])-1])
		numbers := make([]int, len(parts)-1)
		for i := 0; i < len(parts)-1; i++ {
			numbers[i], _ = strconv.Atoi(parts[i+1])
		}
		utils.LogF("%d with %v: ", target, numbers)

		for id := 0; id < utils.Pow(len(operators), (len(numbers)-1)); id++ {
			rawOps := fmt.Sprintf(fmt.Sprintf("%%0%ds", len(numbers)-1), strconv.FormatInt(int64(id), len(operators)))
			ops := strings.Map(func(r rune) rune { return rune(operators[r-'0']) }, rawOps)

			current := numbers[0]
			for o := 0; o < len(numbers)-1; o++ {
				switch ops[o] {
				case '+':
					current += numbers[o+1]
				case '*':
					current *= numbers[o+1]
				case '|':
					current, _ = strconv.Atoi(strconv.Itoa(current) + strconv.Itoa(numbers[o+1]))
				}
			}
			if current == target {
				utils.LogF(fmt.Sprintf("%%0%ds\r\n", len(numbers)-1), ops)
				calibration += target
				goto SUCCESS
			}
		}
		utils.Log("no solution")
	SUCCESS:
	}
	fmt.Printf("Calibration result: %d\r\n", calibration)
}
