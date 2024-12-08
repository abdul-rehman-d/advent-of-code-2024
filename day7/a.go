package day7

import (
	"advent-of-code-2024/utils"
	"fmt"
	"strconv"
	"strings"
)

type Row struct {
	Ans  int
	Nums []int
}

func parse(lines []string) []Row {
	rows := make([]Row, len(lines))

	for idx, line := range lines {
		temp := strings.Split(line, ": ")
		if len(temp) < 2 {
			panic("what")
		}

		sum, _ := strconv.Atoi(temp[0])

		nums := []int{}
		for _, num := range strings.Split(temp[1], " ") {
			nump, _ := strconv.Atoi(num)
			nums = append(nums, nump)
		}

		rows[idx] = Row{
			Ans:  sum,
			Nums: nums,
		}

	}

	return rows
}

func getAllCombinations(n int) []string {
	characters := []rune{'+', '*'}
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

func PartA(data string) int {
	lines := strings.Split(data, "\n")
	lines = utils.FilterEmptyLines(lines)

	rows := parse(lines)

	answer := 0

	for _, row := range rows {
		allCombinations := getAllCombinations(len(row.Nums) - 1)

		for _, combo := range allCombinations {
			ans := row.Nums[0]
			for i := 1; i < len(row.Nums); i++ {
				switch combo[i-1] {
				case '+':
					ans += row.Nums[i]
				case '*':
					ans *= row.Nums[i]
				}
			}
			if ans == row.Ans {
				answer += row.Ans
				break
			}
		}
		fmt.Println(row)
	}

	return answer
}
