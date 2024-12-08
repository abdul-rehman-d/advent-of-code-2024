package day6

import (
	"advent-of-code-2024/utils"
	"strings"
)

func PartB(data string) int {
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

	og := make([]string, len(lines))
	copy(og, lines)

	startingPosition := Coord{
		Row:       currentPosition.Row,
		Col:       currentPosition.Col,
		LookingAt: currentPosition.LookingAt,
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

	count1 := 0
	count := 0
	for row, line := range lines {
		for i := 0; i < len(line); i++ {
			if line[i] == 'X' {
				if row == startingPosition.Row && i == startingPosition.Col {
					continue
				}
				count1++
				temp := make([]string, len(og))
				copy(temp, og)
				temp[row] = replaceStringAtIndex(temp[row], i, 'O')

				currentPosition = Coord{
					Row:       startingPosition.Row,
					Col:       startingPosition.Col,
					LookingAt: startingPosition.LookingAt,
				}
				seen := map[Coord]int{}
				seen[currentPosition] = 1

				next := currentPosition.getNext()
				for {
					if next.Row < 0 {
						break
					}
					if next.Row >= len(temp) {
						break
					}
					if next.Col < 0 {
						break
					}
					if next.Col >= len(temp[next.Row]) {
						break
					}

					if temp[next.Row][next.Col] == '#' || temp[next.Row][next.Col] == 'O' {
						switch next.LookingAt {
						case "up":
							next.LookingAt = "right"
						case "right":
							next.LookingAt = "down"
						case "down":
							next.LookingAt = "left"
						case "left":
							next.LookingAt = "up"
						}
						next.Row = currentPosition.Row
						next.Col = currentPosition.Col
					}

					if _, ok := seen[next]; ok {
						count++
						break
					} else {
						seen[next] = 1
					}

					currentPosition = Coord{
						Row:       next.Row,
						Col:       next.Col,
						LookingAt: next.LookingAt,
					}
					temp[currentPosition.Row] = replaceStringAtIndex(
						temp[currentPosition.Row],
						currentPosition.Col,
						'X',
					)
					next = currentPosition.getNext()
				}
			}
		}
	}
	return count
}
