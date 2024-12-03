package day2

import (
	"advent-of-code-2024/utils"
	"fmt"
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

func PartA(data string) int {
	lines := strings.Split(data, "\n")
	lines = utils.FilterEmptyLines(lines)

	safe := 0

	for _, report := range parse(lines) {
		if len(report) <= 1 {
			panic("why only one value")
		}

		flag := false
		dir := "inc"
		if report[1] < report[0] {
			dir = "dec"
		}

		for i := 0; i < len(report)-1; i++ {
			diff := report[i+1] - report[i]
			// increasing or decreasing
			if diff > 0 && dir == "dec" {
				flag = true
				break
			}
			if diff < 0 && dir == "inc" {
				flag = true
				break
			}
			// Any two adjacent levels differ by at least one and at most three.
			if math.Abs(float64(diff)) < 1 || math.Abs(float64(diff)) > 3 {
				flag = true
				break
			}
		}

		if !flag {
			fmt.Println(report)
			safe++
		}
	}

	return safe
}
