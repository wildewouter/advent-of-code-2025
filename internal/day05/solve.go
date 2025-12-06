package day05

import (
	"strconv"
	"strings"
)

func PartOne(input string) (int, error) {
	parts := strings.SplitN(input, "\n\n", 2)

	ranges := strings.Split(parts[0], "\n")
	ingredients := strings.Split(parts[1], "\n")

	fresh := 0

	for _, ingredientString := range ingredients {
		ingredient, _ := strconv.Atoi(ingredientString)
		for _, freshRange := range ranges {
			startWithEnd := strings.SplitN(freshRange, "-", 2)

			start, _ := strconv.Atoi(startWithEnd[0])
			end, _ := strconv.Atoi(startWithEnd[1])

			if ingredient >= start && ingredient <= end {
				fresh++
				break
			}
		}
	}

	return fresh, nil
}

func PartTwo(input string) (int, error) {
	parts := strings.SplitN(input, "\n\n", 2)

	ranges := strings.Split(parts[0], "\n")

	var freshMap = make(map[Range]int)

rangeLoop:
	for _, freshRange := range ranges {
		startWithEnd := strings.SplitN(freshRange, "-", 2)
		start, _ := strconv.Atoi(startWithEnd[0])
		end, _ := strconv.Atoi(startWithEnd[1])

		for newRange := range freshMap {
			if newRange.start <= start && end <= newRange.end {
				continue rangeLoop
			}

			if start < newRange.start && newRange.end < end {
				delete(freshMap, newRange)
			}

			if start < newRange.start && end <= newRange.end && end >= newRange.start {
				end = newRange.start - 1
			}

			if start >= newRange.start && start <= newRange.end && end > newRange.end {
				start = newRange.end + 1
			}
		}

		freshMap[Range{start: start, end: end}] = end - start + 1
	}

	fresh := 0

	for _, amount := range freshMap {
		fresh += amount
	}

	return fresh, nil
}

type Range struct {
	start int
	end   int
}
