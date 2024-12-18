package day13

import (
	"advent-of-code-2024/utils"
	"math"
	"strconv"
	"strings"
)

type Move struct {
	X int
	Y int
}

type Machine struct {
	A     Move
	B     Move
	Prize Move
}

func parseMovement(str string) Move {
	str = strings.Replace(str, "X+", "", -1)
	str = strings.Replace(str, "X=", "", -1)
	str = strings.Replace(str, "Y+", "", -1)
	str = strings.Replace(str, "Y=", "", -1)
	str = strings.Replace(str, " ", "", -1)

	rx, ry := strings.Split(str, ",")[0], strings.Split(str, ",")[1]

	x, _ := strconv.Atoi(rx)
	y, _ := strconv.Atoi(ry)

	return Move{X: x, Y: y}
}

func parseMachinesA(data string) []Machine {
	lines := strings.Split(data, "\n")
	lines = utils.FilterEmptyLines(lines)

	machines := []Machine{}

	for i := 0; i < len(lines)-2; i += 3 {
		ra, _ := strings.CutPrefix(lines[i+0], "Button A: ")
		rb, _ := strings.CutPrefix(lines[i+1], "Button B: ")
		rp, _ := strings.CutPrefix(lines[i+2], "Prize: ")

		machines = append(machines, Machine{
			A:     parseMovement(ra),
			B:     parseMovement(rb),
			Prize: parseMovement(rp),
		})
	}

	return machines
}

func solve(data string, parseMachines func(string) []Machine) int {
	machines := parseMachines(data)

	answer := 0

	for _, m := range machines {
		tokens := math.MaxInt64
		for numB := 1; numB < m.Prize.X/m.B.X; numB++ {
			bX := numB * m.B.X
			if bX > m.Prize.X {
				break
			}
			remainingX := m.Prize.X - bX
			if remainingX%m.A.X != 0 {
				// possible
				continue
			}
			numA := remainingX / m.A.X

			// works with y too
			if m.Prize.Y == (numB*m.B.Y)+(numA*m.A.Y) {
				newTokens := (numA * 3) + (numB * 1)
				if newTokens < tokens {
					tokens = newTokens
				}
			}
		}
		if tokens != math.MaxInt64 {
			answer += tokens
		}
	}

	return answer
}

func PartA(data string) int {
	return solve(data, parseMachinesA)
}
