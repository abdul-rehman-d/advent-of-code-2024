package day11

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

const NUM_OF_BLINKS = 25

func PartA(data string) int {
	lines := strings.Split(data, "\n")
	if len(lines) < 1 {
		panic("why")
	}

	stones := []int{}
	for _, numRaw := range strings.Split(lines[0], " ") {
		if len(numRaw) == 0 {
			continue
		}
		num, _ := strconv.Atoi(numRaw)
		stones = append(stones, num)
	}

	fmt.Println(stones)

	for i := 0; i < NUM_OF_BLINKS; i++ {
		newStones := make([]int, len(stones))
		copy(newStones, stones)
		idx := 0
		for _, stone := range stones {
			if stone == 0 {
				newStones[idx] = 1
				idx++
				continue
			}

			if numOfDigits := getNumOfDigits(stone); numOfDigits%2 == 0 {
				num1, num2 := SplitNumIntoTwo(stone)
				newStones = slices.Concat(
					newStones[:idx],
					[]int{num1, num2},
					newStones[idx+1:],
				)
				idx += 2
				continue
			}

			newStones[idx] = stone * 2024
			idx++
		}
		stones = newStones
		fmt.Printf("Done with round %d\n", i+1)
	}

	return len(stones)
}

func getNumOfDigits(num int) int {
	return int(math.Floor(math.Log10(float64(num)))) + 1
}

func SplitNumIntoTwo(num int) (int, int) {
	numOfDigits := getNumOfDigits(num)
	middle := numOfDigits / 2
	numStr := strconv.Itoa(num)
	num1Str := numStr[:middle]
	num2Str := numStr[middle:]

	num1, _ := strconv.Atoi(num1Str)
	num2, _ := strconv.Atoi(num2Str)

	return num1, num2
}
