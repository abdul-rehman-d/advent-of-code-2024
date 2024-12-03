package day2

import (
	"advent-of-code-2024/utils"
	"slices"
	"strings"
)

func deleteIndex(array []int, index int) []int {
	return slices.Concat(array[:index], array[index+1:])
}

func PartB(data string) int {
	lines := strings.Split(data, "\n")
	lines = utils.FilterEmptyLines(lines)

	safe := 0

	for _, report := range parse(lines) {
		if len(report) <= 1 {
			panic("why only one value")
		}

		if check(report) {
			safe++
			continue
		}

		for i := 0; i < len(report); i++ {
			if check(deleteIndex(report, i)) {
				safe++
				break
			}
		}
	}

	return safe
}
