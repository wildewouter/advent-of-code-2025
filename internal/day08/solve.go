package day08

import (
	"fmt"
	"slices"
	"strings"
)

func PartOne(input string) (int, error) {
	lines := strings.Split(input, "\n")

	coords, err := getCoords(lines)
	if err != nil {
		return 0, err
	}

	var circuits [][]Coord

	for _, coord := range coords {
		if isAlreadyPresentInCircuits(coord, circuits) {
			continue
		}
		minimal := closest(coord, coords)
		circuits = insert([2]Coord{minimal, coord}, circuits)
	}

	slices.SortFunc(circuits, func(a, b []Coord) int {
		return len(b) - len(a)
	})

	answer := 1

	for _, circuit := range circuits[:3] {
		answer *= len(circuit)
	}

	return answer, nil
}

func PartTwo(input string) (int, error) {
	return 0, nil
}

func isAlreadyPresentInCircuits(coord Coord, circuits [][]Coord) bool {
	for _, circuit := range circuits {
		if slices.Contains(circuit, coord) {
			return true
		}
	}
	return false
}

func insert(coords [2]Coord, circuits [][]Coord) [][]Coord {
	for i, circuit := range circuits {
		if slices.Contains(circuit, coords[0]) {
			circuits[i] = append(circuits[i], coords[1])
			return circuits
		}
		if slices.Contains(circuit, coords[1]) {
			circuits[i] = append(circuits[i], coords[0])
			return circuits
		}
	}

	return append(circuits, []Coord{coords[0], coords[1]})
}

func closest(coord Coord, coords []Coord) Coord {
	distances := make(map[Coord]int)

	for _, c := range coords {
		if c == coord {
			continue
		}
		distances[c] = distanceSquared(coord, c)
	}

	var closestC Coord
	var minDist int
	for c, d := range distances {
		if minDist == 0 || d < minDist {
			minDist = d
			closestC = c
		}
	}

	return closestC
}

func distanceSquared(a, b Coord) int {
	dx := a[0] - b[0]
	dy := a[1] - b[1]
	dz := a[2] - b[2]
	return dx*dx + dy*dy + dz*dz
}

type Coord [3]int

func getCoords(lines []string) ([]Coord, error) {
	var coords []Coord

	for _, line := range lines {
		var x, y, z int
		_, err := fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
		if err != nil {
			return nil, err
		}
		coords = append(coords, [3]int{x, y, z})
	}
	return coords, nil
}
