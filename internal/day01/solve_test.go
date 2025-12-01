package day01_test

import (
	"advent-of-code-2025/internal/day01"
	"fmt"
	"testing"
)

func TestPartOne(t *testing.T) {
	solve, err := day01.PartOne(input)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("Answer: ", solve)
}

func TestPartTwo(t *testing.T) {
	got, err := day01.PartTwo(input)
	if err != nil {
		t.Fatal(err)
	}

	want := 6

	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

var input = `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`

var input2 = `L68
L300
R90
R356
R1010`
