package day1

import (
	"fmt"
	"slices"
	"strings"
)

func PartB(data string) int {
	lines := strings.Split(data, "\n")

	list1, list2, _ := getNums(lines)

	slices.Sort(list1)
	slices.Sort(list2)

	fmt.Println(list1)
	fmt.Println(list2)

	ans := 0

	leftPointer, rightPointer := 0, 0

	for leftPointer < len(list1) {
		score1, score2 := 1, 0

		num := list1[leftPointer]

		for leftPointer+1 < len(list1) && num == list1[leftPointer+1] {
			score1++
			leftPointer++
		}

		// match the level
		for rightPointer < len(list2) && num > list2[rightPointer] {
			rightPointer++
		}

		for rightPointer < len(list2) && num == list2[rightPointer] {
			score2++
			rightPointer++
		}

		ans += score1 * score2 * num
		leftPointer++
	}

	return ans
}
