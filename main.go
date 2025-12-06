package main

import (
	"advent-of-code-2025/internal/day01"
	"advent-of-code-2025/internal/day02"
	"advent-of-code-2025/internal/day03"
	"advent-of-code-2025/internal/day04"
	"advent-of-code-2025/internal/day05"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	day := flag.Int("day", 1, "Which day to run")
	flag.Parse()

	data, err := os.ReadFile(fmt.Sprintf("input/day%d.txt", *day))
	if err != nil {
		log.Fatal(err)
	}

	switch *day {
	case 1:
		RunWithTimer(day01.PartOne, day01.PartTwo, string(data))
	case 2:
		RunWithTimer(day02.PartOne, day02.PartTwo, string(data))
	case 3:
		RunWithTimer(day03.PartOne, day03.PartTwo, string(data))
	case 4:
		RunWithTimer(day04.PartOne, day04.PartTwo, string(data))
	case 5:
		RunWithTimer(day05.PartOne, day05.PartTwo, string(data))
	default:
		log.Fatalf("day %d not implemented", *day)
	}
}

func RunWithTimer(partOneFn func(string) (int, error), partTwoFn func(string) (int, error), data string) {
	start := time.Now()
	one, err := partOneFn(data)
	partOne := time.Since(start)
	start = time.Now()
	two, err := partTwoFn(data)
	partTwo := time.Since(start)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part One: ", one)
	fmt.Println("Time (Part One): ", partOne)
	fmt.Println("Part Two: ", two)
	fmt.Println("Time (Part Two): ", partTwo)
}
