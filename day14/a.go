package day14

import (
	"advent-of-code-2024/utils"
	"strconv"
	"strings"
)

const (
	LENGTH int = 103
	WIDTH  int = 101
	RUNS   int = 100
)

type Coord struct {
	X int
	Y int
}

type Robot struct {
	Pos  Coord
	Step Coord
}

func (c Coord) Plus(s Coord) Coord {
	nextX := c.X + s.X
	nextY := c.Y + s.Y

	if nextX >= WIDTH {
		nextX = nextX % WIDTH
	} else if nextX < 0 {
		nextX = WIDTH + nextX
	}
	if nextY >= LENGTH {
		nextY = nextY % LENGTH
	} else if nextY < 0 {
		nextY = LENGTH + nextY
	}

	return Coord{nextX, nextY}
}

func parseRobots(data string) []Robot {
	lines := strings.Split(data, "\n")
	lines = utils.FilterEmptyLines(lines)

	robots := make([]Robot, len(lines))

	for idx, line := range lines {
		temp := strings.Split(line, " ")
		if len(temp) < 2 {
			panic("why no space")
		}
		p, v := temp[0], temp[1]
		p, _ = strings.CutPrefix(p, "p=")
		v, _ = strings.CutPrefix(v, "v=")

		pa := strings.Split(p, ",")
		va := strings.Split(v, ",")

		if len(pa) < 2 || len(va) < 2 {
			panic("why no comma")
		}

		pxr, pyr := pa[0], pa[1]
		vxr, vyr := va[0], va[1]

		px, _ := strconv.Atoi(pxr)
		py, _ := strconv.Atoi(pyr)
		vx, _ := strconv.Atoi(vxr)
		vy, _ := strconv.Atoi(vyr)

		robots[idx] = Robot{
			Coord{px, py},
			Coord{vx, vy},
		}
	}

	return robots
}

func PartA(data string) int {
	robots := parseRobots(data)

	grid := make([][]byte, LENGTH)

	for y := 0; y < LENGTH; y++ {
		temp := make([]byte, WIDTH)
		for x := 0; x < WIDTH; x++ {
			temp[x] = '0'
		}
		grid[y] = temp
	}

	for i := 0; i < RUNS; i++ {
		for ri, robot := range robots {
			robots[ri].Pos = robot.Pos.Plus(robot.Step)
		}
	}

	for _, robot := range robots {
		grid[robot.Pos.Y][robot.Pos.X] += 1
	}

	answer := 1

	for q := 0; q < 4; q++ {
		num := 0

		yS := 0
		yE := LENGTH / 2
		if q >= 2 {
			yS = (LENGTH / 2) + 1
			yE = LENGTH
		}
		xS := (q % 2) * ((WIDTH / 2) + 1)
		xE := ((q % 2) + 1) * WIDTH / 2

		for y := yS; y < yE; y++ {
			for x := xS; x < xE; x++ {
				num += int(grid[y][x] - '0')
			}
		}

		if num != 0 {
			answer *= num
		}
	}

	return answer
}
