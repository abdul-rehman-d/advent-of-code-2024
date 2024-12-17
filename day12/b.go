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
			solutions[area] = Sol{0, len(coords)}
			perimeter := 0

			type A struct {
				Coord Coord
				Dir   int
			}
			perimeters := map[A]bool{}

			for _, coord := range coords {
				sides := 0
				for idx, dir := range utils.GetDirections() {
					neightbour := coord.Plus(dir)
					if neightbour.Row >= 0 && neightbour.Row < len(areaMap) &&
						neightbour.Col >= 0 && neightbour.Col < len(areaMap[neightbour.Row]) &&
						area == areaMap[neightbour.Row][neightbour.Col] {
						continue
					}
					sides++
					perimeters[A{coord, idx}] = true
				}
				perimeter += sides
			}

			fmt.Println()
			fmt.Println(area)
			fmt.Println(perimeter)
			fmt.Println(perimeters)
			seen := map[A]bool{}
			for p1 := range perimeters {
				if _, has := seen[p1]; has {
					continue
				}
				fmt.Println(p1)
				toFind := A{
					Coord: p1.Coord,
					Dir:   p1.Dir,
				}
				switch p1.Dir {
				case 0:
					toFind.Coord.Col += 1
				case 1:
					toFind.Coord.Col += 1
				case 2:
					toFind.Coord.Row += 1
				case 3:
					toFind.Coord.Row += 1
				}
				fmt.Println(">", toFind)
				if _, has := perimeters[toFind]; has {
					if _, has := seen[toFind]; !has {
						fmt.Println(">> minus")
						perimeter--
						seen[toFind] = true
					}
				}
				toFind = A{
					Coord: p1.Coord,
					Dir:   p1.Dir,
				}
				switch p1.Dir {
				case 0:
					toFind.Coord.Col -= 1
				case 1:
					toFind.Coord.Col -= 1
				case 2:
					toFind.Coord.Row -= 1
				case 3:
					toFind.Coord.Row -= 1
				}
				fmt.Println(">", toFind)
				if _, has := perimeters[toFind]; has {
					if _, has := seen[toFind]; !has {
						fmt.Println(">> minus")
						perimeter--
						seen[toFind] = true
					}
				}

				fmt.Println(perimeter)
			}
			fmt.Println()

			solutions[area].P = perimeter
		}
		answer += (solutions[area].A * solutions[area].P)
	}

	return answer
}
