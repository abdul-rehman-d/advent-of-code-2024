package day11

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func PartA(data string) int {
	const NUM_OF_BLINKS = 25
	return solve(data, NUM_OF_BLINKS)
}

func solve(data string, num_of_blinks int) int {
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

	count := 0

	for stoneIdx, ogStone := range stones {
		fmt.Printf("Dealing with stone # %d:\t\t%d\n", stoneIdx+1, ogStone)
		stones := []int{ogStone}
		for i := 0; i < num_of_blinks; i++ {
			idx := 0
			maxx := len(stones)
			for idx < maxx {
				stone := stones[idx]
				if stone == 0 {
					stones[idx] = 1
					idx++
					continue
				}

				if numOfDigits := getNumOfDigits(stone); numOfDigits%2 == 0 {
					num1, num2 := SplitNumIntoTwo(stone)
					stones = slices.Concat(
						stones[:idx],
						[]int{num1, num2},
						stones[idx+1:],
					)
					idx += 2
					maxx += 1
					continue
				}

				stones[idx] = stone * 2024
				idx++
			}
		}
		count += len(stones)
	}

	return count
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
