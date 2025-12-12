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

	var circuits []Set[Coord]

	for _, coord := range coords {
		c, _ := closest(coord, coords)
		circuits = insert([2]Coord{coord, c}, circuits)
	}

	circuitList := make([][]Coord, len(circuits))

	for i, circuit := range circuits {
		for coord := range circuit {
			circuitList[i] = append(circuitList[i], coord)
		}
	}

	slices.SortFunc(circuitList, func(a, b []Coord) int {
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

//func isAlreadyPresentInCircuits(coord Coord, circuits [][]Coord) bool {
//	for _, circuit := range circuits {
//		if slices.Contains(circuit, coord) {
//			return true
//		}
//	}
//	return false
//}

type Set[T comparable] map[T]struct{}

func (s Set[T]) insert(v T) Set[T] {
	s[v] = struct{}{}
	return s
}

func insert(coords [2]Coord, circuits []Set[Coord]) []Set[Coord] {
	var indicesContaining []int

	for i, circuit := range circuits {
		for _, coord := range coords {
			_, contains := circuit[coord]
			if contains {
				indicesContaining = append(indicesContaining, i)
			}
		}
	}

	switch len(indicesContaining) {
	case 0:
		return append(circuits, map[Coord]struct{}{coords[0]: {}})
	case 1:
		circuits[indicesContaining[0]].insert(coords[0])
		circuits[indicesContaining[0]].insert(coords[1])
		return circuits
	default:
		mergedCircuit := Set[Coord]{}

		for _, i := range indicesContaining {
			for c := range circuits[i] {
				mergedCircuit.insert(c)
			}
		}

		mergedCircuit.insert(coords[0])
		mergedCircuit.insert(coords[1])

		for _, i := range indicesContaining {
			circuits = append(circuits[:i], circuits[i+1:]...)
		}

		return append(circuits, mergedCircuit)
	}
}

func closest(coord Coord, coords []Coord) (Coord, int) {
	var closestC Coord
	var minDist int
	for _, c := range coords {
		if c == coord {
			continue
		}
		dist := distanceSquared(coord, c)
		if minDist == 0 || dist < minDist {
			minDist = dist
			closestC = c
		}
	}

	return closestC, minDist
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
