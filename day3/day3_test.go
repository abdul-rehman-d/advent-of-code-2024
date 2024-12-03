package day3

import (
	"testing"
)

const (
	dataA = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	dataB = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
)

func TestPartA(t *testing.T) {
	expected := 161
	result := PartA(dataA)

	if expected != result {
		t.Fatalf("\nExpected = %d\nResult = %d\n", expected, result)
	}
}

func TestPartB(t *testing.T) {
	expected := 48
	result := PartB(dataB)

	if expected != result {
		t.Fatalf("\nExpected = %d\nResult = %d\n", expected, result)
	}
}
