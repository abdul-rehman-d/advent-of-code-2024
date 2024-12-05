package day5

import (
	"strconv"
	"strings"
)

func PartB(data string) int {
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
		// newNums := make([]int, len(numsRaw))

		for j, numRaw := range numsRaw {
			nums[j], _ = strconv.Atoi(numRaw)
		}

		isValid := true
		for {
			isValidInner := true
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
					isValidInner = false
					nums[beforeIdx], nums[afterIdx] = nums[afterIdx], nums[beforeIdx]
				} else {
				}
			}
			if isValid {
				break
			}
			if isValidInner {
				break
			}
		}
		if !isValid {
			ans += getMid(nums)
		}
	}

	return ans
}
