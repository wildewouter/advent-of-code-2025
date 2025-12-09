package day07

import "strings"

func PartOne(input string) (int, error) {
	lines := strings.Split(input, "\n")

	var grid = make(map[Coord]rune)

	var start = map[Coord]struct{}{}

	for y, line := range lines {
		for x, char := range []rune(line) {
			coord := Coord{x, y}
			if char == 'S' {
				start[coord] = struct{}{}
			}
			grid[coord] = char
		}
	}

	return down(start, grid, 0), nil
}

func PartTwo(input string) (int, error) {
	return 0, nil
}

type Coord struct {
	x int
	y int
}

func down(coords map[Coord]struct{}, grid map[Coord]rune, splits int) int {
	nextCoords := map[Coord]struct{}{}

	for coord, _ := range coords {
		c := Coord{coord.x, coord.y + 1}
		next, ok := grid[c]

		if !ok {
			return splits
		}

		switch next {
		case '.':
			nextCoords[c] = struct{}{}
		case '^':
			splits++
			nextCoords[Coord{c.x - 1, c.y}] = struct{}{}
			nextCoords[Coord{c.x + 1, c.y}] = struct{}{}
		}
	}

	return down(nextCoords, grid, splits)
}
