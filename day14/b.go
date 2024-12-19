package day14

import (
	"fmt"
)

func printGrid(array [][]byte) {
	for _, row := range array {
		for x := 0; x < len(row); x++ {
			if row[x] == '0' {
				fmt.Print(".")
			} else {
				fmt.Print(string(row[x]))
			}
		}
		fmt.Println()
	}
}

func copyGrid(grid [][]byte) [][]byte {
	out := make([][]byte, len(grid))
	for i, row := range grid {
		temp := make([]byte, len(row))
		for j, col := range row {
			temp[j] = col
		}
		out[i] = temp
	}
	return out
}

func PartB(data string) int {
	robots := parseRobots(data)

	grid := make([][]byte, LENGTH)

	for y := 0; y < LENGTH; y++ {
		temp := make([]byte, WIDTH)
		for x := 0; x < WIDTH; x++ {
			temp[x] = '0'
		}
		grid[y] = temp
	}

	for i := 0; i < 219549; i++ {
		temp := copyGrid(grid)

		uniques := 0

		for ri, robot := range robots {
			robots[ri].Pos = robot.Pos.Plus(robot.Step)
			temp[robots[ri].Pos.Y][robots[ri].Pos.X] += 1
		}

		for y := 0; y < LENGTH; y++ {
			for x := 0; x < WIDTH; x++ {
				if temp[y][x] == '1' {
					uniques++
				}
			}
		}

		if uniques == len(robots) {
			printGrid(temp)
			return i + 1
		}
	}

	return 0
}
