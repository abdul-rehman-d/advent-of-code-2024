package day12

import (
	"advent-of-code-2024/utils"
	"fmt"
	"strings"
)

type Coord = utils.Coord

type Sol struct {
	P int
	A int
}

func PartA(data string) int {
	lines := strings.Split(data, "\n")
	lines = utils.FilterEmptyLines(lines)

	// mapp := map[byte][]Coord{}
	// seen := map[Coord]bool{}
	// allCharacters := map[byte]Coord{}

	// startingPoint := Coord{
	// 	Row: 0,
	// 	Col: 0,
	// }
	// current := lines[startingPoint.Row][startingPoint.Col]

	areaMap := make([][]int, len(lines))
	areaMapReverse := map[int][]Coord{}
	toCheck := map[Coord]bool{}

	for row, line := range lines {
		for col := 0; col < len(line); col++ {
			toCheck[Coord{
				Row: row,
				Col: col,
			}] = true
		}
	}

	// id := 0
	//
	// getAreaId := func() int {
	// 	id++
	// 	return id
	// }
	//
	var findAll func(Coord, int)
	findAll = func(current Coord, area int) {
		if _, has := toCheck[current]; !has {
			return
		}
		currentCharacter := lines[current.Row][current.Col]
		if _, ok := areaMapReverse[area]; ok {
			areaMapReverse[area] = append(areaMapReverse[area], current)
		} else {
			areaMapReverse[area] = []Coord{current}
		}
		delete(toCheck, current)
		for _, dir := range utils.GetDirections() {
			next := current.Plus(dir)
			if next.Row < 0 || next.Row >= len(areaMap) {
				continue
			}
			if next.Col < 0 || next.Col >= len(areaMap[next.Row]) {
				continue
			}

			nextCharacter := lines[next.Row][next.Col]
			if nextCharacter == currentCharacter {
				findAll(next, area)
			} else {
				findAll(next, area+1)
			}
		}

	}
	for coord := range toCheck {
		findAll(coord, 0)
	}
	// for row, line := range lines {
	// 	temp := make([]int, len(line))
	// 	areaMap[row] = temp
	//
	// 	col := 0
	//
	// 	for {
	// 		current := line[col]
	//
	// 		// up
	// 		if row != 0 && lines[row-1][col] == current {
	// 			temp[col] = areaMap[row-1][col]
	// 			continue
	// 		}
	//
	// 		// left
	// 		if col != 0 && lines[row][col-1] == current {
	// 			temp[col] = temp[col-1]
	// 			continue
	// 		}
	//
	// 		// new area
	// 		temp[col] = id
	// 		id++
	// 	}
	// }

	// for row := range areaMap {
	// 	for col := range areaMap[row] {
	// 		area := areaMap[row][col]
	// 		if _, ok := areaMapReverse[area]; ok {
	// 			areaMapReverse[area] = append(areaMapReverse[area], Coord{
	// 				Row: row,
	// 				Col: col,
	// 			})
	// 		} else {
	// 			areaMapReverse[area] = []Coord{{
	// 				Row: row,
	// 				Col: col,
	// 			}}
	// 		}
	// 	}
	// }
	//
	// answer := 0
	// solutions := make([]Sol, len(areaMapReverse))
	//
	// for area, coords := range areaMapReverse {
	// 	if len(coords) == 1 {
	// 		solutions[area] = Sol{4, 1}
	// 	} else {
	// 		solutions[area] = Sol{0, len(coords)}
	// 		perimeter := 0
	//
	// 		for _, coord := range coords {
	// 			sides := 4
	// 			for _, dir := range utils.GetDirections() {
	// 				neightbour := coord.Plus(dir)
	// 				if neightbour.Row < 0 || neightbour.Row >= len(areaMap) {
	// 					continue
	// 				}
	// 				if neightbour.Col < 0 || neightbour.Col >= len(areaMap[neightbour.Row]) {
	// 					continue
	// 				}
	// 				if area == areaMap[neightbour.Row][neightbour.Col] {
	// 					sides--
	// 				}
	// 			}
	// 			perimeter += sides
	// 		}
	// 		solutions[area].P = perimeter
	// 	}
	// 	answer += (solutions[area].A * solutions[area].P)
	// }
	//
	for _, line := range lines {
		fmt.Println(strings.Split(line, ""))
	}
	fmt.Println()
	for _, line := range areaMap {
		fmt.Println(line)
	}
	fmt.Println()
	for _, line := range areaMapReverse {
		character := lines[line[0].Row][line[0].Col]
		fmt.Println(string(character), line)
		// fmt.Println(string(character), area, ">>>>>>>>>>>>", solutions[area])
	}

	fmt.Println()
	fmt.Println()

	return 0
}
