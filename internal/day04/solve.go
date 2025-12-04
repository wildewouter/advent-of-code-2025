package day04

import "strings"

func PartOne(input string) (int, error) {
	lines := strings.Split(input, "\n")

	var grid = make(map[Coord]bool)

	for y, line := range lines {
		for x, char := range strings.Split(line, "") {
			grid[Coord{x, y}] = char == "@"
		}
	}

	rolls := 0

	for coord, isRoll := range grid {
		if !isRoll {
			continue
		}
		adjacent := 0

		directions := []Coord{
			{coord.x, coord.y + 1},
			{coord.x + 1, coord.y + 1},
			{coord.x + 1, coord.y - 1},
			{coord.x + 1, coord.y},
			{coord.x - 1, coord.y + 1},
			{coord.x, coord.y - 1},
			{coord.x - 1, coord.y - 1},
			{coord.x - 1, coord.y},
		}

		for _, direction := range directions {
			isRoll, ok := grid[direction]

			if ok && isRoll {
				adjacent++
			}
		}

		if adjacent < 4 {
			rolls++
		}
	}

	return rolls, nil
}

func PartTwo(input string) (int, error) {
	lines := strings.Split(input, "\n")

	var grid = make(map[Coord]bool)

	for y, line := range lines {
		for x, char := range strings.Split(line, "") {
			grid[Coord{x, y}] = char == "@"
		}
	}

	return electricBoogaloo(grid, 0), nil
}

func electricBoogaloo(grid map[Coord]bool, rolls int) int {
	prevRollCount := rolls

	for coord, isRoll := range grid {
		if !isRoll {
			continue
		}
		adjacent := 0

		directions := []Coord{
			{coord.x, coord.y + 1},
			{coord.x + 1, coord.y + 1},
			{coord.x + 1, coord.y - 1},
			{coord.x + 1, coord.y},
			{coord.x - 1, coord.y + 1},
			{coord.x, coord.y - 1},
			{coord.x - 1, coord.y - 1},
			{coord.x - 1, coord.y},
		}

		for _, direction := range directions {
			isRoll, ok := grid[direction]

			if ok && isRoll {
				adjacent++
			}
		}

		if adjacent < 4 {
			rolls++
			grid[coord] = false
		}
	}

	if prevRollCount == rolls {
		return rolls
	}

	return electricBoogaloo(grid, rolls)
}

type Coord struct {
	x int
	y int
}
