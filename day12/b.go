package day12

import (
	"advent-of-code-2024/utils"
	"fmt"
	"strings"
)

func PartB(data string) int {
	lines := strings.Split(data, "\n")
	lines = utils.FilterEmptyLines(lines)

	areaMap, areaMapReverse := parseBoys(lines)
	fmt.Println("AREA MAP=>")
	for _, line := range areaMap {
		fmt.Println(line)
	}
	fmt.Println()

	answer := 0
	solutions := make([]Sol, len(areaMapReverse))

	for area, coords := range areaMapReverse {
		if len(coords) == 1 {
			solutions[area] = Sol{4, 1}
		} else {
			fmt.Println("area =>", area)
			fmt.Println("char =>", string(lines[coords[0].Row][coords[0].Col]))
			solutions[area] = Sol{0, len(coords)}
			// perimeter := 0

			type A struct {
				Val int
				Dir string
			}

			sides := map[A]bool{}

			for _, coord := range coords {
				fmt.Println(coord)
				for idx, dir := range utils.GetDirections() {
					neightbour := coord.Plus(dir)
					if neightbour.Row >= 0 && neightbour.Row < len(areaMap) &&
						neightbour.Col >= 0 && neightbour.Col < len(areaMap[neightbour.Row]) &&
						area == areaMap[neightbour.Row][neightbour.Col] {
						continue
					}
					fmt.Println(idx, neightbour)

					switch idx {
					case 0: // up
						sides[A{
							Val: coord.Row,
							Dir: "up",
						}] = true
					case 1: // down
						sides[A{
							Val: neightbour.Row,
							Dir: "up",
						}] = true
					case 2: // right
						sides[A{
							Val: coord.Col,
							Dir: "right",
						}] = true
					case 3: // left
						sides[A{
							Val: neightbour.Col,
							Dir: "right",
						}] = true
					}
				}
				fmt.Println(sides)
			}
			fmt.Println()
			fmt.Println(sides)
			fmt.Println("perimeter =>", len(sides))
			// fmt.Println(perimeter)
			solutions[area].P = len(sides)
		}
		answer += (solutions[area].A * solutions[area].P)
		fmt.Println()
	}

	return answer
}
