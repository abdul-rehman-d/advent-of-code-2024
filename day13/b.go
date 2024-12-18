package day13

import (
	"advent-of-code-2024/utils"
	"strings"
)

func parseMachinesB(data string) []Machine {
	lines := strings.Split(data, "\n")
	lines = utils.FilterEmptyLines(lines)

	machines := []Machine{}

	for i := 0; i < len(lines)-2; i += 3 {
		ra, _ := strings.CutPrefix(lines[i+0], "Button A: ")
		rb, _ := strings.CutPrefix(lines[i+1], "Button B: ")
		rp, _ := strings.CutPrefix(lines[i+2], "Prize: ")

		p := parseMovement(rp)
		machines = append(machines, Machine{
			A: parseMovement(ra),
			B: parseMovement(rb),
			Prize: Move{
				// X: p.X,
				// Y: p.Y,
				X: p.X + 10000000000000,
				Y: p.Y + 10000000000000,
			},
		})
	}

	return machines
}

func solveB(data string, parseMachines func(string) []Machine) int {
	machines := parseMachines(data)

	answer := 0

	for _, m := range machines {
		// x1a + x2b = c1
		// y1a + y2b = c2
		// =================================
		// | x1 x2 || a | = | c1 |
		// | y1 y2 || a | = | c2 |
		// =================================
		// a = y2c1 - x2c2
		//     -----------
		//     x1y2 - x2y1
		// =================================
		// b = x1c2 - y1c1
		//     -----------
		//     x1y2 - x2y1

		x1 := m.A.X
		x2 := m.B.X
		y1 := m.A.Y
		y2 := m.B.Y
		c1 := m.Prize.X
		c2 := m.Prize.Y

		d := (x1 * y2) - (x2 * y1)

		a := ((y2 * c1) - (x2 * c2)) / d
		b := ((x1 * c2) - (y1 * c1)) / d

		if (x1*a)+(x2*b) == c1 && (y1*a)+(y2*b) == c2 {
			tokens := (a * 3) + b
			answer += tokens
		}
	}

	return answer
}

func PartB(data string) int {
	return solveB(data, parseMachinesB)
}
