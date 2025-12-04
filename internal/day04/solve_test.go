package day04_test

import (
	"advent-of-code-2025/internal/day04"
	"testing"
)

func TestPartOne(t *testing.T) {
	got, err := day04.PartOne(input)
	if err != nil {
		t.Fatal(err)
	}

	want := 13

	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	got, err := day04.PartTwo(input)
	if err != nil {
		t.Fatal(err)
	}

	want := 43

	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

var input = `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.
`
