package day7

import (
	"advent-of-code-2024/utils"
	"math"
	"strings"
)

func getAllCombinationsB(n int) []string {
	characters := []rune{'+', '*', '|'}
	var results []string

	var combine func(current string, depth int)
	combine = func(current string, depth int) {
		if depth == n {
			results = append(results, current)
			return
		}
		for _, char := range characters {
			combine(current+string(char), depth+1)
		}
	}

	combine("", 0)
	return results
}

func PartB(data string) int {
	lines := strings.Split(data, "\n")
	lines = utils.FilterEmptyLines(lines)

	rows := parse(lines)

	answer := 0

	for _, row := range rows {
		allCombinations := getAllCombinationsB(len(row.Nums) - 1)

		for _, combo := range allCombinations {
			ans := row.Nums[0]
			for i := 1; i < len(row.Nums); i++ {
				switch combo[i-1] {
				case '+':
					ans += row.Nums[i]
				case '*':
					ans *= row.Nums[i]
				case '|':
					numOfDigit := math.Floor(math.Log10(float64(row.Nums[i]))) + 1
					for i := 0; i < int(numOfDigit); i++ {
						ans *= 10
					}
					ans += row.Nums[i]
				}
			}
			if ans == row.Ans {
				answer += row.Ans
				break
			}
		}
	}

	return answer
}
