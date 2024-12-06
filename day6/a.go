package day6

import (
	"advent-of-code-2024/utils"
	"strings"
)

type Coord struct {
	Row       int
	Col       int
	LookingAt string
}

func (current Coord) getNext() Coord {
	next := current
	switch current.LookingAt {
	case "up":
		next.Row -= 1
	case "down":
		next.Row += 1
	case "left":
		next.Col -= 1
	case "right":
		next.Col += 1
	}
	return next
}

func replaceStringAtIndex(s string, idx int, ch byte) string {
	out := ""
	for i := 0; i < len(s); i++ {
		if i == idx {
			out += string(ch)
		} else {
			out += string(s[i])
		}
	}
	return out
}

func PartA(data string) int {
	lines := strings.Split(data, "\n")
	lines = utils.FilterEmptyLines(lines)

	currentPosition := Coord{-1, -1, "up"}

	for i, line := range lines {
		pos := strings.Index(line, "^")
		if pos > -1 {
			lines[i] = replaceStringAtIndex(line, pos, 'X')
			currentPosition.Row = i
			currentPosition.Col = pos
			break
		}
	}

	for {
		next := currentPosition.getNext()
		if next.Row < 0 {
			break
		}
		if next.Row >= len(lines) {
			break
		}
		if next.Col < 0 {
			break
		}
		if next.Col >= len(lines[next.Row]) {
			break
		}

		if lines[next.Row][next.Col] == '#' {
			switch currentPosition.LookingAt {
			case "up":
				currentPosition.LookingAt = "right"
			case "right":
				currentPosition.LookingAt = "down"
			case "down":
				currentPosition.LookingAt = "left"
			case "left":
				currentPosition.LookingAt = "up"
			}
			continue
		}
		currentPosition = next
		lines[currentPosition.Row] = replaceStringAtIndex(
			lines[currentPosition.Row],
			currentPosition.Col,
			'X',
		)

	}

	count := 0
	for _, line := range lines {
		for i := 0; i < len(line); i++ {
			if line[i] == 'X' {
				count++
			}
		}
	}
	return count
}
