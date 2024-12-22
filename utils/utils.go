package utils

import (
	"fmt"
	"slices"
	"strings"

	"golang.org/x/exp/constraints"
)

func FilterEmptyLines(data []string) []string {
	out := make([]string, 0, len(data))
	for _, line := range data {
		if line != "" {
			out = append(out, line)
		}
	}
	return out
}

func DeleteIndex[T any](array []T, idx int) []T {
	return slices.Concat(array[:idx], array[idx+1:])
}

func ReplaceStringAtIndex(s string, idx int, ch byte) string {
	out := ""
	for i := 0; i < len(s); i++ {
		if i == idx {
			out += string(ch)
		} else {
			out += string(s[i])
		}
	}
	return out
}

func MakeIntMatrix(data string) [][]int {
	lines := strings.Split(data, "\n")
	lines = FilterEmptyLines(lines)

	out := make([][]int, len(lines))

	for i, line := range lines {
		temp := make([]int, len(line))
		for i := 0; i < len(line); i++ {
			temp[i] = int(line[i] - '0')
		}
		out[i] = temp
	}
	return out
}

func PrintGrid[T constraints.Integer](array [][]T, opts ...bool) {
	for _, row := range array {
		for x := 0; x < len(row); x++ {
			fmt.Print(string(row[x]))
			if x != len(row)-1 && len(opts) == 0 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
