package day4

import (
	"advent-of-code-2024/utils"
	"fmt"
	"strings"
)

const word = "XMAS"

type Coord struct {
	X int
	Y int
}

type SingleSolution struct {
	row int
	col int
	dir Coord
}

type Solution struct {
	Solutions []SingleSolution
}

func NewSolution() *Solution {
	return &Solution{
		Solutions: []SingleSolution{},
	}
}

func (s *Solution) print(m [][]string) {
	toPrint := make([][]string, len(m), len(m))
	for i := 0; i < len(m); i++ {
		toAppend := make([]string, len(m[i]))
		for j := 0; j < len(m[i]); j++ {
			toAppend[j] = "."
		}
		toPrint[i] = toAppend
	}

	for _, sol := range s.Solutions {
		toPrint[sol.row][sol.col] = m[sol.row][sol.col]
		multiplier := 1
		for multiplier < len(word) {
			i, j := sol.row+(sol.dir.X*multiplier), sol.col+(sol.dir.Y*multiplier)
			toPrint[i][j] = m[i][j]
			multiplier++
		}
	}

	for _, row := range toPrint {
		fmt.Println(row)
	}
}

func check(m [][]string, row, col, idx int, dir Coord) bool {
	i, j := row+dir.X, col+dir.Y
	if i >= len(m) || i < 0 || j < 0 || j >= len(m[i]) {
		return false
	}

	ch := m[i][j]
	if ch != string(word[idx]) {
		return false
	}
	if idx == len(word)-1 {
		return true
	}
	return check(m, i, j, idx+1, dir)
}

func (s *Solution) seeAround(m [][]string, row, col int) {
	dirs := []Coord{
		{-1, 0},  // up
		{1, 0},   // down
		{0, -1},  // left
		{0, 1},   // right
		{-1, -1}, // up left
		{-1, 1},  // up right
		{1, -1},  // down left
		{1, 1},   // down right
	}

	for _, dir := range dirs {
		if check(m, row, col, 1, dir) {
			s.Solutions = append(s.Solutions, SingleSolution{
				row,
				col,
				dir,
			})
		}
	}
}

func PartA(data string) int {
	solutiion := NewSolution()

	lines := strings.Split(data, "\n")
	lines = utils.FilterEmptyLines(lines)
	matrix := make([][]string, len(lines))
	for idx, line := range lines {
		matrix[idx] = strings.Split(line, "")
	}

	row, col := 0, 0

	for row < len(lines) {
		col = 0

		for col < len(lines) {
			if matrix[row][col] == string(word[0]) {
				solutiion.seeAround(matrix, row, col)
			}

			col++
		}

		row++
	}

	// solutiion.print(matrix)

	return len(solutiion.Solutions)
}
