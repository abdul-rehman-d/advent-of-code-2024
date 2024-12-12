package day12

import (
	"testing"
)

const (
	data1 = `OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`
	data2 = `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`
)

func TestPartA(t *testing.T) {
	expected1 := 772
	expected2 := 1930
	result1 := PartA(data1)
	result2 := PartA(data2)

	if expected1 != result1 {
		t.Fatalf("\nExpected = %d\nResult = %d\n", expected1, result1)
	}
	if expected2 != result2 {
		t.Fatalf("\nExpected = %d\nResult = %d\n", expected2, result2)
	}
}

func TestPartB(t *testing.T) {
	expected1 := 0
	expected2 := 0
	result1 := PartB(data1)
	result2 := PartB(data2)

	if expected1 != result1 {
		t.Fatalf("\nExpected = %d\nResult = %d\n", expected1, result1)
	}
	if expected2 != result2 {
		t.Fatalf("\nExpected = %d\nResult = %d\n", expected2, result2)
	}
}
