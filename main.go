package main

import (
	"advent-of-code-2025/internal/day01"
	"advent-of-code-2025/internal/day02"
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
		start := time.Now()
		one, err := day01.PartOne(string(data))
		partOne := time.Since(start)
		start = time.Now()
		two, err := day01.PartTwo(string(data))
		partTwo := time.Since(start)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Part One: ", one)
		fmt.Println("Time (Part One): ", partOne)
		fmt.Println("Part Two: ", two)
		fmt.Println("Time (Part Two): ", partTwo)
	case 2:
		start := time.Now()
		one, err := day02.PartOne(string(data))
		partOne := time.Since(start)
		start = time.Now()
		two, err := day02.PartTwo(string(data))
		partTwo := time.Since(start)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Part One: ", one)
		fmt.Println("Time (Part One): ", partOne)
		fmt.Println("Part Two: ", two)
		fmt.Println("Time (Part Two): ", partTwo)
	default:
		log.Fatalf("day %d not implemented", *day)
	}
}
