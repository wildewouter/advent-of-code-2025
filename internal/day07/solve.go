package day07

import "strings"

func PartOne(input string) (int, error) {
	grid, start := getGridAndStart(input)

	return down(map[Coord]struct{}{start: {}}, grid, 0), nil
}

func PartTwo(input string) (int, error) {
	grid, start := getGridAndStart(input)

	return downManyWorlds(map[Coord]struct{}{start: {}}, grid, make(map[Coord]int)), nil
}

func getGridAndStart(input string) (map[Coord]rune, Coord) {
	lines := strings.Split(input, "\n")

	var grid = make(map[Coord]rune)

	var start Coord

	for y, line := range lines {
		for x, char := range []rune(line) {
			coord := Coord{x, y}
			if char == 'S' {
				start = coord
			}
			grid[coord] = char
		}
	}
	return grid, start
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

func downManyWorlds(coords map[Coord]struct{}, grid map[Coord]rune, hits map[Coord]int) int {
	nextCoords := map[Coord]struct{}{}

	for coord := range coords {
		c := Coord{coord.x, coord.y + 1}
		next, ok := grid[c]

		if !ok {
			total := 0
			for hitC, w := range hits {
				if hitC.y == coord.y {
					total += w
				}
			}
			return total
		}

		previousCount := hits[coord]
		if previousCount == 0 {
			previousCount = 1
		}

		switch next {
		case '.':
			nextCoords[c] = struct{}{}
			hits[c] = previousCount + hits[c]
		case '^':
			next1 := Coord{c.x - 1, c.y}
			next2 := Coord{c.x + 1, c.y}

			hits[next1] = hits[next1] + previousCount
			hits[next2] = hits[next2] + previousCount

			nextCoords[next1] = struct{}{}
			nextCoords[next2] = struct{}{}
		}
	}

	return downManyWorlds(nextCoords, grid, hits)
}
