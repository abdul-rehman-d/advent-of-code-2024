package day2

import (
	"advent-of-code-2024/utils"
	"math"
	"strconv"
	"strings"
)

func parse(lines []string) [][]int {
	parsed := make([][]int, len(lines))
	for idx, line := range lines {
		temp := strings.Split(strings.TrimSpace(line), " ")
		parsedLine := make([]int, 0, len(temp))
		for _, numStr := range temp {
			if numStr == "" {
				continue
			}
			num, err := strconv.Atoi(numStr)
			if err != nil {
				panic(err)
			}
			parsedLine = append(parsedLine, num)
		}
		parsed[idx] = parsedLine
	}
	return parsed
}

func check(report []int) bool {
	dir := "inc"
	if report[1] < report[0] {
		dir = "dec"
	}
	for i := 0; i < len(report)-1; i++ {
		diff := report[i+1] - report[i]
		// increasing or decreasing
		cond1 := diff > 0 && dir == "dec"
		cond2 := diff < 0 && dir == "inc"
		// Any two adjacent levels differ by at least one and at most three.
		cond3 := math.Abs(float64(diff)) < 1 || math.Abs(float64(diff)) > 3

		if cond1 || cond2 || cond3 {
			return false
		}
	}
	return true
}

func PartA(data string) int {
	lines := strings.Split(data, "\n")
	lines = utils.FilterEmptyLines(lines)

	safe := 0

	for _, report := range parse(lines) {
		if len(report) <= 1 {
			panic("why only one value")
		}

		if check(report) {
			safe++
		}
	}

	return safe
}
