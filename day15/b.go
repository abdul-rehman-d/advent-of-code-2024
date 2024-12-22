package day15

import (
	"advent-of-code-2024/utils"
	"slices"
	"strings"
)

func check(next Coord, grid [][]byte, isVertical bool, dir Coord) ([]Coord, bool) {
	n := []Coord{}
	ch := grid[next.Row][next.Col]

	if ch == '#' {
		return nil, true
	} else if ch == '[' || ch == ']' {
		n = append(n, next)
		dir2 := Coord{Row: 0, Col: 0}
		if isVertical {
			if ch == '[' {
				dir2.Col = 1
			} else {
				dir2.Col = -1
			}

			opp := next.Plus(dir2)
			n = append(n, next.Plus(dir2))

			stack, nextIsWall := check(
				opp.Plus(dir),
				grid,
				isVertical,
				dir,
			)
			if nextIsWall {
				return nil, true
			}
			n = slices.Concat(n, stack)
		}
		stack, nextIsWall := check(
			next.Plus(dir),
			grid,
			isVertical,
			dir,
		)

		if nextIsWall {
			return nil, true
		}

		return slices.Concat(n, stack), false
	} else {
		return n, false
	}
}

func PartB(data string) int {
	splitted := strings.Split(data, "\n\n")
	if len(splitted) < 2 {
		panic("why no two line breaks")
	}

	gridR := strings.Split(splitted[0], "\n")
	grid := make([][]byte, len(gridR))

	robot := Coord{}

	for rowIdx, row := range gridR {
		temp := make([]byte, len(row)*2)
		for col, col2 := 0, 0; col < len(row); col, col2 = col+1, col2+2 {
			switch row[col] {
			case '@':
				robot = Coord{
					Row: rowIdx,
					Col: col2,
				}
				temp[col2] = '@'
				temp[col2+1] = '.'
			case '#':
				temp[col2] = '#'
				temp[col2+1] = '#'
			case '.':
				temp[col2] = '.'
				temp[col2+1] = '.'
			case 'O':
				temp[col2] = '['
				temp[col2+1] = ']'
			}
		}
		grid[rowIdx] = temp
	}

	steps := strings.ReplaceAll(splitted[1], "\n", "")

	for _, step := range steps {
		dir := Coord{Row: 0, Col: 0}
		isVertical := false
		switch step {
		case '^':
			isVertical = true
			dir.Row = -1
		case 'v':
			isVertical = true
			dir.Row = 1
		case '>':
			dir.Col = 1
		case '<':
			dir.Col = -1
		}

		stack, nextIsWall := check(robot.Plus(dir), grid, isVertical, dir)

		stackSet := map[Coord]bool{}

		if nextIsWall {
			continue
		}

		stack = slices.Concat([]Coord{robot}, stack)

		for i := len(stack) - 1; i >= 0; i-- {
			toSwap := stack[i]
			toSwapCh := grid[toSwap.Row][toSwap.Col]

			if _, ok := stackSet[toSwap]; ok {
				continue
			}
			stackSet[toSwap] = true

			toSwapWith := stack[i].Plus(dir)

			grid[toSwapWith.Row][toSwapWith.Col], grid[toSwap.Row][toSwap.Col] = grid[toSwap.Row][toSwap.Col], grid[toSwapWith.Row][toSwapWith.Col]

			if toSwapCh == '@' {
				robot = toSwapWith
			}
		}
	}
	utils.PrintGrid(grid, false)

	answer := 0
	for row, line := range grid {
		col := 0
		for col < len(line) {
			ch := line[col]
			if ch == '[' {
				answer += (row * 100) + col
				col++
			}
			col++
		}
	}

	return answer
}
