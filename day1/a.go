package day1

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func insertRightly(nums []int, num int) []int {
	return append(nums, num)
}

func getNums(lines []string) ([]int, []int, int) {
	count := len(lines)
	list1 := make([]int, 0, len(lines))
	list2 := make([]int, 0, len(lines))
	for _, line := range lines {
		temp := strings.Split(strings.TrimSpace(line), " ")
		if len(temp) < 2 {
			count--
			continue
		}
		left, err := strconv.Atoi(temp[0])
		if err != nil {
			panic(err)
		}
		right, err := strconv.Atoi(temp[len(temp)-1])
		if err != nil {
			panic(err)
		}
		list1 = insertRightly(list1, left)
		list2 = insertRightly(list2, right)
	}
	return list1, list2, count
}

func PartA(data string) int {
	lines := strings.Split(data, "\n")

	list1, list2, count := getNums(lines)

	slices.Sort(list1)
	slices.Sort(list2)

	sum := 0

	for i := 0; i < count; i++ {
		diff := int(math.Abs(float64(list1[i]) - float64(list2[i])))
		fmt.Println(list1[i], list2[i], diff)
		sum += diff
	}

	return sum
}
