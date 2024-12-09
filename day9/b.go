package day9

import (
	"advent-of-code-2024/utils"
	"strings"
)

type Unit struct {
	Start   int
	HowMany int
}

func PartB(data string) int {
	data = strings.TrimSpace(data)
	memory := []int{}
	spaces := []Unit{}
	files := []Unit{}

	id := 0
	for i := 0; i < len(data); i++ {
		num := int(data[i] - '0')
		idx := len(memory)
		if num == 0 {
			continue
		}
		if i%2 == 0 {
			// file
			files = append(files, Unit{idx, num})
			for j := 0; j < num; j++ {
				memory = append(memory, id)
			}
			id++
		} else {
			// free space
			spaces = append(spaces, Unit{idx, num})
			for j := 0; j < num; j++ {
				memory = append(memory, -1)
			}

		}
	}

	right := len(files) - 1
	left := 0
	for right > 0 {
		if left == len(spaces) {
			right--
			left = 0
			continue
		}

		file := files[right]
		space := spaces[left]

		if space.Start >= file.Start {
			right--
			left = 0
			continue
		}

		if file.HowMany > space.HowMany {
			left++
			continue
		}

		for i := space.Start; i < space.Start+file.HowMany; i++ {
			memory[i] = right
		}
		for i := file.Start; i < file.Start+file.HowMany; i++ {
			memory[i] = -1
		}
		files[right].Start = space.Start
		if file.HowMany == space.HowMany {
			// spaces = slices.Concat(spaces[:left], spaces[left+1:])
			spaces = utils.DeleteIndex(spaces, left)
		} else {
			spaces[left].HowMany -= file.HowMany
			spaces[left].Start += file.HowMany
		}
		right--
		left = 0
	}

	answer := 0

	for i := 0; i < len(memory); i++ {
		if memory[i] == -1 {
			continue
		}

		answer += memory[i] * i
	}

	return answer

}
