package main

import (
	"advent-of-code-2025/internal/day01"
	"flag"
	"fmt"
	"log"
	"os"
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
		result, err := day01.Solve(string(data))
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Answer: ", result)
	default:
		log.Fatalf("day %d not implemented", *day)
	}
}
