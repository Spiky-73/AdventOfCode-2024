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
	scanner := bufio.NewScanner(os.Stdin)

	utils.Log("Parsing rules...")

	rules := make(map[string]map[string]struct{})
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		parts := strings.Split(line, "|")

		if _, ok := rules[parts[0]]; !ok {
			rules[parts[0]] = make(map[string]struct{})
		}
		after := rules[parts[0]]
		after[parts[1]] = struct{}{}
	}
	utils.Log("Rules:")
	for key, rs := range rules {
		utils.LogF("%s: ", key)
		for val := range rs {
			utils.LogF("%s ", val)
		}
		utils.Log("")
	}

	utils.Log("Checking updates...")
	middleNumbers := 0
	for scanner.Scan() {
		line := scanner.Text()
		utils.LogF("%s: ", line)
		var middle int
		parts := strings.Split(line, ",")
		for i := 0; i < len(parts); i++ {
			for j := 0; j < i; j++ {
				if _, after := rules[parts[i]][parts[j]]; after {
					utils.LogF("%s|%s broken\r\n", parts[i], parts[j])
					goto END_LOOP
				}
			}
		}
		middle, _ = strconv.Atoi(parts[int(len(parts)/2)])
		middleNumbers += middle
		utils.LogF("ordered with middle page %d\r\n", middle)
	END_LOOP:
	}
	fmt.Printf("Middle pages sum: %d\r\n", middleNumbers)
}

func part2() {
	scanner := bufio.NewScanner(os.Stdin)

	utils.Log("Parsing rules...")

	rules := make(map[string]map[string]struct{})
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		parts := strings.Split(line, "|")

		if _, ok := rules[parts[0]]; !ok {
			rules[parts[0]] = make(map[string]struct{})
		}
		after := rules[parts[0]]
		after[parts[1]] = struct{}{}
	}
	utils.Log("Rules:")
	for key, rs := range rules {
		utils.LogF("%s: ", key)
		for val := range rs {
			utils.LogF("%s ", val)
		}
		utils.Log("")
	}

	utils.Log("Checking updates...")
	middleNumbers := 0
	for scanner.Scan() {
		line := scanner.Text()
		utils.LogF("Checking %s\r\n", line)
		var middle int
		parts := strings.Split(line, ",")

		swapped := false
	RETRY_CHECK:
		for i := 0; i < len(parts); i++ {
			for j := 0; j < i; j++ {
				if _, after := rules[parts[i]][parts[j]]; after {
					utils.LogF("\tswapped %s and %s\r\n", parts[i], parts[j])
					swapped = true
					tmp := parts[i]
					parts[i] = parts[j]
					parts[j] = tmp
					goto RETRY_CHECK
				}
			}
		}
		middle, _ = strconv.Atoi(parts[int(len(parts)/2)])
		utils.LogF("\tordered with middle page %d\r\n", middle)
		if swapped {
			middleNumbers += middle
		}
	}
	fmt.Printf("Middle pages sum: %d\r\n", middleNumbers)
}
