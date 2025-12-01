package day01_test

import (
	"advent-of-code-2025/internal/day01"
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	solve, err := day01.Solve(input)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("Answer: ", solve)
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
