package day10

import (
	"testing"
)

const (
	data = `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732
`
)

func TestPartA(t *testing.T) {
	expected := 36
	result := PartA(data)

	if expected != result {
		t.Fatalf("\nExpected = %d\nResult = %d\n", expected, result)
	}
}

func TestPartB(t *testing.T) {
	expected := 81
	result := PartB(data)

	if expected != result {
		t.Fatalf("\nExpected = %d\nResult = %d\n", expected, result)
	}
}
