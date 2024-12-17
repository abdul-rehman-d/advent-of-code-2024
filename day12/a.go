package day12

import (
	"advent-of-code-2024/utils"
	// "fmt"
	"strings"
)

type Coord = utils.Coord

type Sol struct {
	P int
	A int
}

func parseBoys(lines []string) ([][]int, map[int][]Coord) {
	areaMap := make([][]int, len(lines))
	// bigTemp := make([][]byte, len(lines))
	areaMapReverse := map[int][]Coord{}
	toCheck := map[Coord]bool{}

	for row, line := range lines {
		temp := make([]int, len(line))
		// temp2 := make([]byte, len(line))
		for col := 0; col < len(line); col++ {
			temp[col] = -1
			// temp2[col] = '.'
			toCheck[Coord{
				Row: row,
				Col: col,
			}] = true
		}
		areaMap[row] = temp
		// bigTemp[row] = temp2
	}

	var findAll func(Coord, int) bool
	findAll = func(current Coord, area int) bool {
		if _, has := toCheck[current]; !has {
			return false
		}
		currentCharacter := lines[current.Row][current.Col]
		if _, ok := areaMapReverse[area]; ok {
			areaMapReverse[area] = append(areaMapReverse[area], current)
		} else {
			areaMapReverse[area] = []Coord{current}
		}

		areaMap[current.Row][current.Col] = area
		// bigTemp[current.Row][current.Col] = currentCharacter
		delete(toCheck, current)

		for _, dir := range utils.GetDirections() {
			next := current.Plus(dir)
			if next.Row < 0 || next.Row >= len(lines) {
				continue
			}
			if next.Col < 0 || next.Col >= len(lines[next.Row]) {
				continue
			}

			nextCharacter := lines[next.Row][next.Col]
			if nextCharacter == currentCharacter {
				findAll(next, area)
			}
		}

		return true

	}

	area := 0
	for coord := range toCheck {
		if findAll(coord, area) {
			area++
		}
	}

	return areaMap, areaMapReverse
}

func PartA(data string) int {
	lines := strings.Split(data, "\n")
	lines = utils.FilterEmptyLines(lines)

	areaMap, areaMapReverse := parseBoys(lines)

	answer := 0
	solutions := make([]Sol, len(areaMapReverse))

	for area, coords := range areaMapReverse {
		if len(coords) == 1 {
			solutions[area] = Sol{4, 1}
		} else {
			solutions[area] = Sol{0, len(coords)}
			perimeter := 0

			for _, coord := range coords {
				sides := 4
				for _, dir := range utils.GetDirections() {
					neightbour := coord.Plus(dir)
					if neightbour.Row < 0 || neightbour.Row >= len(areaMap) {
						continue
					}
					if neightbour.Col < 0 || neightbour.Col >= len(areaMap[neightbour.Row]) {
						continue
					}
					if area == areaMap[neightbour.Row][neightbour.Col] {
						sides--
					}
				}
				perimeter += sides
			}
			solutions[area].P = perimeter
		}
		answer += (solutions[area].A * solutions[area].P)
	}

	// fmt.Println("DATA=>")
	// for _, line := range lines {
	// 	fmt.Println(strings.Split(line, ""))
	// }
	// fmt.Println()
	// fmt.Println("VERIFY DATA=>")
	// for _, line := range bigTemp {
	// 	fmt.Print("[ ")
	// 	for _, ch := range line {
	// 		fmt.Printf("%s ", string(ch))
	// 	}
	// 	fmt.Println("]")
	// }
	// fmt.Println()
	// fmt.Println("AREA MAP=>")
	// for _, line := range areaMap {
	// 	fmt.Println(line)
	// }
	// fmt.Println()
	// fmt.Println("AREA MAP REVERSE=>")
	// for _, line := range areaMapReverse {
	// 	character := lines[line[0].Row][line[0].Col]
	// 	fmt.Println(string(character), line)
	// 	// fmt.Println(string(character), area, ">>>>>>>>>>>>", solutions[area])
	// }

	// fmt.Println()
	// fmt.Println(answer)
	// fmt.Println()
	// fmt.Println()

	return answer
}
