package day11

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Coord struct {
	Stone int
	N     int
}

type Solver struct {
	Cache map[Coord]int
}

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

	solver := &Solver{
		Cache: map[Coord]int{},
	}

	count := 0

	for stoneIdx, ogStone := range stones {
		fmt.Printf("Dealing with stone # %d:\t\t%d\n", stoneIdx+1, ogStone)
		stone := Coord{Stone: ogStone, N: num_of_blinks}
		count += solver.do(stone)
	}

	return count
}

func (solver *Solver) do(stone Coord) int {
	if count, has := solver.Cache[stone]; has {
		return count
	}
	if stone.N == 0 {
		solver.Cache[stone] = 1
		return 1
	}
	if stone.Stone == 0 {
		solver.Cache[stone] = solver.do(Coord{
			Stone: 1,
			N:     stone.N - 1,
		})
		return solver.Cache[stone]
	}
	if numOfDigits := getNumOfDigits(stone.Stone); numOfDigits%2 == 0 {
		num1, num2 := SplitNumIntoTwo(stone.Stone)

		solver.Cache[stone] = solver.do(Coord{
			Stone: num1,
			N:     stone.N - 1,
		}) + solver.do(Coord{
			Stone: num2,
			N:     stone.N - 1,
		})
		return solver.Cache[stone]
	}
	solver.Cache[stone] = solver.do(Coord{
		Stone: stone.Stone * 2024,
		N:     stone.N - 1,
	})
	return solver.Cache[stone]
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
