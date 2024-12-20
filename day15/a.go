package day15

import (
	"advent-of-code-2024/utils"
	"strings"
)

type Coord = utils.Coord

func PartA(data string) int {
	splitted := strings.Split(data, "\n\n")
	if len(splitted) < 2 {
		panic("why no two line breaks")
	}

	gridR := strings.Split(splitted[0], "\n")
	grid := make([][]byte, len(gridR))

	robot := Coord{}

	for rowIdx, row := range gridR {
		temp := make([]byte, len(row))
		for col := 0; col < len(row); col++ {
			if row[col] == '@' {
				robot = Coord{
					Row: rowIdx,
					Col: col,
				}
			}
			temp[col] = row[col]
		}
		grid[rowIdx] = temp
	}

	steps := strings.ReplaceAll(splitted[1], "\n", "")

	for _, step := range steps {
		dir := Coord{Row: 0, Col: 0}
		switch step {
		case '^':
			dir.Row = -1
		case 'v':
			dir.Row = 1
		case '>':
			dir.Col = 1
		case '<':
			dir.Col = -1
		}

		obsToPush := 0
		nextIsWall := false

		next := robot.Plus(dir)
		for {
			if grid[next.Row][next.Col] == '#' {
				nextIsWall = true
				break
			} else if grid[next.Row][next.Col] == 'O' {
				obsToPush++
			} else {
				break
			}
			next = next.Plus(dir)
		}

		if nextIsWall {
			continue
		}

		if obsToPush > 0 {
			negativeDir := Coord{
				Row: dir.Row * -1,
				Col: dir.Col * -1,
			}
			// next would be empty space
			for i := 1; i <= obsToPush; i++ {
				temp := grid[next.Row][next.Col]
				nextNext := next.Plus(negativeDir)
				grid[nextNext.Row][nextNext.Col] = temp
				grid[next.Row][next.Col] = 'O'

				next = nextNext
			}
		}

		next = robot.Plus(dir)
		grid[robot.Row][robot.Col] = '.'
		grid[next.Row][next.Col] = '@'
		robot = next
	}
	// utils.PrintGrid(grid)

	answer := 0
	for row, line := range grid {
		for col, ch := range line {
			if ch == 'O' {
				answer += (row * 100) + col
			}
		}
	}

	return answer
}
