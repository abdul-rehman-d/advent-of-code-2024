package main

import (
	"advent-of-code-2024/day1"
	"fmt"
	"log"
	"os"
	"strconv"
)

func getChallengeInput(i int) string {
	arrByte, err := os.ReadFile(fmt.Sprintf("./day%d/input.txt", i))
	if err != nil {
		panic(err)
	}
	return string(arrByte)
}

func getMainFileContents(day int, part string) string {
	return fmt.Sprintf(`package day%d

func Part%s(data string) int {
	return 0
}
`, day, part)
}

func getTestFileContents(day int) string {
	return fmt.Sprintf(`package day%d

import (
	"testing"
)

const (
	data = ""
)

func TestPartA(t *testing.T) {
	expected := 0
	result := PartA(data)

	if expected != result {
		t.Fatalf("\nExpected = %%d\nResult = %%d\n", expected, result)
	}
}


func TestPartB(t *testing.T) {
	expected := 0
	result := PartB(data)

	if expected != result {
		t.Fatalf("\nExpected = %%d\nResult = %%d\n", expected, result)
	}
}
`, day)
}

func generate(day int) error {
	folderName := fmt.Sprintf("./day%d", day)

	if s, err := os.Stat(folderName); !os.IsNotExist(err) && s != nil {
		return fmt.Errorf("already exists")
	}

	err := os.Mkdir(folderName, 0755)
	if err != nil {
		return err
	}

	var f *os.File

	f, err = os.Create(fmt.Sprintf("%s/a.go", folderName))
	if err != nil {
		return err
	}
	f.WriteString(getMainFileContents(day, "A"))

	f, err = os.Create(fmt.Sprintf("%s/b.go", folderName))
	if err != nil {
		return err
	}
	f.WriteString(getMainFileContents(day, "B"))

	_, err = os.Create(fmt.Sprintf("%s/input.txt", folderName))
	if err != nil {
		return err
	}

	_, err = os.Create(fmt.Sprintf("%s/questionA.md", folderName))
	if err != nil {
		return err
	}
	_, err = os.Create(fmt.Sprintf("%s/questionB.md", folderName))
	if err != nil {
		return err
	}

	f, err = os.Create(fmt.Sprintf("%s/day%d_test.go", folderName, day))
	if err != nil {
		return err
	}
	f.WriteString(getTestFileContents(day))

	return nil
}

func solve(day int) {

	partA := []func(string) int{day1.PartA}
	partB := []func(string) int{day1.PartB}

	if day > len(partA) {
		log.Fatalf("Day cannot be more than %d\n", len(partA))
	}

	fmt.Printf("Part A Answer => %d\n", partA[day-1](getChallengeInput(day)))
	fmt.Printf("Part B Answer => %d\n", partB[day-1](getChallengeInput(day)))
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println(`Usage ./aoc [cmd] [day]
Available commands:-
- solve
- generate`)
		os.Exit(1)
	}

	day, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("Error parsing day %s\n", os.Args[2])
		os.Exit(1)
	}

	if day <= 0 {
		fmt.Println("Day cannot be less than 1")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "solve":
		solve(day)
	case "generate":
		err := generate(day)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	default:
		fmt.Println("Invalid command.")
		os.Exit(1)
	}
}
