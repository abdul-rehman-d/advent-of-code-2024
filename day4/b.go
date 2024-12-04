package day4

import (
	"advent-of-code-2024/utils"
	"strings"
)

func checkCross(m [][]string, row, col int) bool {
	upLeft := Coord{-1, -1}
	upRight := Coord{-1, 1}
	downLeft := Coord{1, -1}
	downRight := Coord{1, 1}

	ulx, uly := row+upLeft.X, col+upLeft.Y
	if ulx >= len(m) || ulx < 0 || uly < 0 || uly >= len(m[ulx]) {
		return false
	}
	urx, ury := row+upRight.X, col+upRight.Y
	if urx >= len(m) || urx < 0 || ury < 0 || ury >= len(m[urx]) {
		return false
	}
	dlx, dly := row+downLeft.X, col+downLeft.Y
	if dlx >= len(m) || dlx < 0 || dly < 0 || dly >= len(m[dlx]) {
		return false
	}
	drx, dry := row+downRight.X, col+downRight.Y
	if drx >= len(m) || drx < 0 || dry < 0 || dry >= len(m[drx]) {
		return false
	}

	ulCh, urCh, dlCh, drCh := m[ulx][uly], m[urx][ury], m[dlx][dly], m[drx][dry]

	validRight := (ulCh == "M" && drCh == "S") || (ulCh == "S" && drCh == "M")
	validLeft := (urCh == "M" && dlCh == "S") || (urCh == "S" && dlCh == "M")

	return validLeft && validRight
}

func PartB(data string) int {
	solutions := []Coord{}

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
			if matrix[row][col] == "A" {
				if checkCross(matrix, row, col) {
					solutions = append(solutions, Coord{row, col})
				}
			}

			col++
		}

		row++
	}

	return len(solutions)
}
