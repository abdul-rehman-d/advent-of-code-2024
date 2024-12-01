package day1

import (
	"testing"
)

func TestPartA(t *testing.T) {
	data := `3   4
4   3
2   5
1   3
3   9
3   3`

	expected := 11
	result := PartA(data)

	if expected != result {
		t.Fatalf("\nExpected = %d\nResult = %d\n", expected, result)
	}
}

func TestPartB(t *testing.T) {
	data := ""

	expected := 0
	result := PartB(data)

	if expected != result {
		t.Fatalf("\nExpected = %d\nResult = %d\n", expected, result)
	}
}
