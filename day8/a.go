package day8

import (
	"advent-of-code-2024/utils"
	"strings"
)

type Coord struct {
	X int
	Y int
}

func PartA(data string) int {
	lines := strings.Split(data, "\n")
	lines = utils.FilterEmptyLines(lines)

	mapp := map[byte][]Coord{}

	for x, line := range lines {
		for y := 0; y < len(line); y++ {
			if line[y] == '.' {
				continue
			}
			val, ok := mapp[line[y]]
			if ok {
				mapp[line[y]] = append(val, Coord{x, y})
			} else {
				mapp[line[y]] = []Coord{{x, y}}
			}
		}
	}

	for a, poss := range mapp {
		for _, refPos := range poss {
			for _, pos := range poss {
				if refPos.X == pos.X && refPos.Y == pos.Y {
					continue
				}
				diff := Coord{
					refPos.X - pos.X, refPos.Y - pos.Y,
				}
				antiPos := Coord{
					refPos.X + diff.X, refPos.Y + diff.Y,
				}

				if antiPos.X < 0 || antiPos.X >= len(lines) {
					continue
				}
				if antiPos.Y < 0 || antiPos.Y >= len(lines[antiPos.X]) {
					continue
				}

				if lines[antiPos.X][antiPos.Y] != a {
					lines[antiPos.X] = utils.ReplaceStringAtIndex(lines[antiPos.X], antiPos.Y, '#')
				}
			}
		}
	}

	count := 0
	for _, line := range lines {
		for y := 0; y < len(line); y++ {
			if line[y] == '#' {
				count++
			}
		}
	}

	return count
}
