package day06_test

import (
	"advent-of-code-2025/internal/day06"
	"testing"
)

func TestPartOne(t *testing.T) {
	got, err := day06.PartOne(input)
	if err != nil {
		t.Fatal(err)
	}

	want := 4277556

	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	got, err := day06.PartTwo(input)
	if err != nil {
		t.Fatal(err)
	}

	want := 0

	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

var input = `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `
