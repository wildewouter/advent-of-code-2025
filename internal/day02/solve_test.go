package day02_test

import (
	"advent-of-code-2025/internal/day02"
	"testing"
)

func TestPartOne(t *testing.T) {
	got, err := day02.PartOne(input)
	if err != nil {
		t.Fatal(err)
	}

	want := 1227775554

	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	got, err := day02.PartTwo(input)
	if err != nil {
		t.Fatal(err)
	}

	want := 4174379265

	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

var input = `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`
