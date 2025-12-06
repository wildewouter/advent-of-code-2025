package day05_test

import (
	"advent-of-code-2025/internal/day05"
	"testing"
)

func TestPartOne(t *testing.T) {
	got, err := day05.PartOne(input)
	if err != nil {
		t.Fatal(err)
	}

	want := 3

	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	got, err := day05.PartTwo(input)
	if err != nil {
		t.Fatal(err)
	}

	want := 14

	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

var input = `3-5
10-14
16-20
12-18

1
5
8
11
17
32`
