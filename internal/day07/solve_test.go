package day07_test

import (
	"advent-of-code-2025/internal/day07"
	"testing"
)

func TestPartOne(t *testing.T) {
	got, err := day07.PartOne(input)
	if err != nil {
		t.Fatal(err)
	}

	want := 21

	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	got, err := day07.PartTwo(input)
	if err != nil {
		t.Fatal(err)
	}

	want := 0

	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

var input = `.......S.......
...............
.......^.......
...............
......^.^......
...............
.....^.^.^.....
...............
....^.^...^....
...............
...^.^...^.^...
...............
..^...^.....^..
...............
.^.^.^.^.^...^.
...............`
