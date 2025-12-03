package day03_test

import (
	"advent-of-code-2025/internal/day03"
	"testing"
)

func TestPartOne(t *testing.T) {
	got, err := day03.PartOne(input)
	if err != nil {
		t.Fatal(err)
	}

	want := 357

	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	got, err := day03.PartTwo(input)
	if err != nil {
		t.Fatal(err)
	}

	want := 3121910778619

	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

var input = `987654321111111
811111111111119
234234234234278
818181911112111`
