package day5

import (
	"fmt"
	"strconv"
	"strings"
)

type Rule struct {
	Before int
	After  int
}

func PartA(data string) int {
	seperated := strings.Split(data, "\n\n")
	if len(seperated) < 2 {
		panic("what")
	}

	rules := parseRules(seperated[0])

	ans := 0

	for _, row := range strings.Split(seperated[1], "\n") {
		if len(row) == 0 {
			continue
		}
		numsRaw := strings.Split(row, ",")

		nums := make([]int, len(numsRaw))

		for j, numRaw := range numsRaw {
			nums[j], _ = strconv.Atoi(numRaw)
		}

		isValid := true

		for _, rule := range rules {
			beforeIdx := -1
			afterIdx := -1
			for idx, num := range nums {
				if num == rule.Before {
					beforeIdx = idx
					continue
				}
				if num == rule.After {
					afterIdx = idx
					if beforeIdx > -1 {
						break
					}
				}
			}
			if afterIdx != -1 && beforeIdx != -1 && beforeIdx > afterIdx {
				isValid = false
				break
			}
		}
		if isValid {
			ans += getMid(nums)
		}
	}

	return ans

}

func parseRules(data string) []Rule {
	rulesRaw := strings.Split(data, "\n")

	rules := make([]Rule, len(rulesRaw))

	for idx, rule := range rulesRaw {
		splitted := strings.Split(rule, "|")
		if len(splitted) < 2 {
			panic(fmt.Sprintf("rule line %d: why less than two?", idx))
		}
		before, _ := strconv.Atoi(splitted[0])
		after, _ := strconv.Atoi(splitted[1])

		rules[idx] = Rule{
			Before: before,
			After:  after,
		}
	}

	return rules
}

func getMid(nums []int) int {
	midIdx := len(nums) / 2
	return nums[midIdx]
}
