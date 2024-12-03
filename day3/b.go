package day3

import (
	"advent-of-code-2024/utils"
	"strings"
	"unicode"
)

func PartB(data string) int {
	result := 0

	lines := strings.Split(data, "\n")
	lines = utils.FilterEmptyLines(lines)

	state := "idle" // idle | first | second | disabled
	firstNum := 0
	secondNum := 0
	digitCount := 0

	for _, line := range lines {
		i := 0
		for i < len(line) {
			ch := line[i]
			if state == "disabled" {
				if ch == 'd' {
					if line[i:i+4] == "do()" {
						state = "idle"
						i += 4
						continue
					}
				}
				i++
			} else {
				if ch == 'd' && line[i:i+7] == "don't()" {
					state = "disabled"
					i += 7
					firstNum = 0
					secondNum = 0
					digitCount = 0
					continue
				}
			}
			if state == "idle" {
				if ch == 'm' {
					if line[i:i+4] == "mul(" {
						state = "first"
						i += 4
						continue
					}
				}
				i++
			} else if state == "first" {
				if digitCount > 0 && ch == ',' {
					state = "second"
					digitCount = 0
					i++
					continue
				}
				if digitCount < 3 && unicode.IsDigit(rune(ch)) {
					firstNum *= 10
					firstNum += int(ch - '0')
					digitCount++
					i++
					continue
				}
				state = "idle"
				firstNum = 0
				digitCount = 0
				i++
			} else if state == "second" {
				if digitCount > 0 && ch == ')' {
					state = "idle"
					digitCount = 0
					result += (firstNum * secondNum)
					firstNum = 0
					secondNum = 0
					i++
					continue
				}
				if digitCount < 3 && unicode.IsDigit(rune(ch)) {
					secondNum *= 10
					secondNum += int(ch - '0')
					digitCount++
					i++
					continue
				}
				state = "idle"
				firstNum = 0
				secondNum = 0
				digitCount = 0
				i++
			}
		}
	}

	return result
}
