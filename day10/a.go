package day10

import (
	"advent-of-code-2024/utils"
	"fmt"
)

type Coord = utils.Coord

func getAll(data [][]int, digit int) []Coord {
	out := []Coord{}
	for y := 0; y < len(data); y++ {
		for x := 0; x < len(data); x++ {
			if data[y][x] == digit {
				out = append(out, Coord{Row: y, Col: x})
			}
		}
	}
	return out
}

func PartA(data string) int {
	mapp := utils.MakeIntMatrix(data)

	fmt.Println(mapp)

	allStartingPoints := getAll(mapp, 0)
	// maxScoreForEach := len(getAll(mapp, 9))

	ans := 0

	for _, sp := range allStartingPoints {
		stack := []Coord{sp}
		score := map[Coord]int{}
		for len(stack) != 0 {
			fmt.Println("stack>", stack)
			fmt.Println("score>", score)
			curr := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			currHeight := mapp[curr.Row][curr.Col]
			nextHeight := currHeight + 1
			dirs := utils.GetDirections()

			for _, dir := range dirs {
				nextCoord := curr.Plus(dir)
				if nextCoord.Row < 0 || nextCoord.Row >= len(mapp) {
					continue
				}
				if nextCoord.Col < 0 || nextCoord.Col >= len(mapp[nextCoord.Row]) {
					continue
				}

				if nextHeight == mapp[nextCoord.Row][nextCoord.Col] {
					if nextHeight == 9 {
						if _, ok := score[nextCoord]; !ok {
							score[nextCoord] = 1
						}
						continue
					}
					// happy case

					stack = append(stack, nextCoord)
				}
			}
		}

		ans += len(score)
	}

	return ans
}
