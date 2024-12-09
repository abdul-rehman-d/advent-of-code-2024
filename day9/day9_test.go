package day9

import (
	"testing"
)

const (
	data = "2333133121414131402"
)

func TestPartA(t *testing.T) {
	expected := 1928
	result := PartA(data)

	if expected != result {
		t.Fatalf("\nExpected = %d\nResult = %d\n", expected, result)
	}
}

func TestPartB(t *testing.T) {
	expected := 0
	result := PartB(data)

	if expected != result {
		t.Fatalf("\nExpected = %d\nResult = %d\n", expected, result)
	}
}
