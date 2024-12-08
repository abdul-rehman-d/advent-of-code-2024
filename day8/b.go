package day8

import (
	"advent-of-code-2024/utils"
	"strings"
)

func PartB(data string) int {
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

	for _, poss := range mapp {
		for _, refPos := range poss {
			for _, pos := range poss {
				if refPos.X == pos.X && refPos.Y == pos.Y {
					continue
				}
				diff := Coord{
					refPos.X - pos.X, refPos.Y - pos.Y,
				}
				antiPos := Coord{refPos.X, refPos.Y}

				for {
					if antiPos.X < 0 || antiPos.X >= len(lines) {
						break
					}
					if antiPos.Y < 0 || antiPos.Y >= len(lines[antiPos.X]) {
						break
					}

					lines[antiPos.X] = utils.ReplaceStringAtIndex(lines[antiPos.X], antiPos.Y, '#')

					antiPos = Coord{
						antiPos.X + diff.X, antiPos.Y + diff.Y,
					}
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
