package day1

import (
	"testing"
)

const (
	data = `3   4
4   3
2   5
1   3
3   9
3   3`
)

func TestPartA(t *testing.T) {
	expected := 11
	result := PartA(data)

	if expected != result {
		t.Fatalf("\nExpected = %d\nResult = %d\n", expected, result)
	}
}

func TestPartB(t *testing.T) {
	expected := 31
	result := PartB(data)

	if expected != result {
		t.Fatalf("\nExpected = %d\nResult = %d\n", expected, result)
	}
}
