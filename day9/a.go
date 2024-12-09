package day9

import (
	"strings"
)

func PartA(data string) int {
	data = strings.TrimSpace(data)
	memory := []int{}

	id := 0
	for i := 0; i < len(data); i++ {
		num := int(data[i] - '0')
		if i%2 == 0 {
			// file
			for j := 0; j < num; j++ {
				memory = append(memory, id)
			}
			id++
		} else {
			// free space
			for j := 0; j < num; j++ {
				memory = append(memory, -1)
			}
		}
	}

	i, j := 0, len(memory)-1
	for i < j {
		if memory[i] != -1 {
			i++
			continue
		}
		if memory[j] == -1 {
			j--
			continue
		}
		memory[i] = memory[j]
		memory[j] = -1
		i++
		j--
	}

	answer := 0

	for i := 0; i < len(memory); i++ {
		if memory[i] == -1 {
			break
		}

		answer += memory[i] * i
	}

	return answer
}
