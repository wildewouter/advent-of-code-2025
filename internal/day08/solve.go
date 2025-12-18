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
	closestPairs := getPairs(coords)

	answer := getAnswer(closestPairs[:1000])

	return answer, nil
}

func getAnswer(closestPairs PairingList) int {
	var circuits []Set[Coord]

	for _, pair := range closestPairs {
		circuits = insert(pair, circuits)
	}

	slices.SortFunc(circuits, func(a, b Set[Coord]) int {
		return len(b) - len(a)
	})

	answer := 1

	for _, circuit := range circuits[:3] {
		answer *= len(circuit)
	}

	return answer
}

func getPairs(coords []Coord) PairingList {
	var literallyAllPairs PairingList

	for _, coord := range coords {
		for _, c := range coords {
			contains := literallyAllPairs.contains(Pairing{coord, c, 0})
			if c == coord || contains {
				continue
			}
			dist := distanceSquared(coord, c)

			literallyAllPairs = append(literallyAllPairs, Pairing{coord, c, dist})
		}
	}
	literallyAllPairs.sortByDistance()
	return literallyAllPairs
}

func PartTwo(input string) (int, error) {
	return 0, nil
}

type Pairing struct {
	a        Coord
	b        Coord
	distance int
}

type PairingList []Pairing

func (l PairingList) sortByDistance() PairingList {
	slices.SortFunc(l, func(a, b Pairing) int { return a.distance - b.distance })
	return l
}

type Set[T comparable] map[T]struct{}

func (s Set[T]) insert(v T) Set[T] {
	s[v] = struct{}{}
	return s
}
func (s Set[T]) contains(v T) bool {
	_, ok := s[v]
	return ok
}

func (s Set[T]) items() []T {
	var items []T
	for item := range s {
		items = append(items, item)
	}
	return items
}

func insert(pair Pairing, circuits []Set[Coord]) []Set[Coord] {
	indicesContaining := Set[int]{}

	for i, circuit := range circuits {
		for _, coord := range [2]Coord{pair.a, pair.b} {
			if circuit.contains(coord) {
				indicesContaining.insert(i)
			}
		}
	}

	switch len(indicesContaining) {
	case 0:
		return append(circuits, Set[Coord]{pair.a: {}, pair.b: {}})
	case 1:
		for i := range indicesContaining {
			circuits[i].insert(pair.a)
			circuits[i].insert(pair.b)
		}

		return circuits
	default:
		mergedCircuit := Set[Coord]{}

		for i := range indicesContaining {
			for c := range circuits[i] {
				mergedCircuit.insert(c)
			}
		}

		indices := indicesContaining.items()

		slices.SortFunc(indices, func(a, b int) int { return b - a })

		for _, i := range indices {
			if len(circuits)-1 <= i {
				circuits = circuits[:i]
				continue
			}

			circuits = append(circuits[:i], circuits[i+1:]...)
		}

		return append(circuits, mergedCircuit)
	}
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

func (l PairingList) contains(p Pairing) bool {
	contains := false

	for _, pairing := range l {
		contains = pairing.a == p.a && pairing.b == p.b || pairing.a == p.b && pairing.b == p.a

		if contains {
			return contains
		}
	}

	return contains
}
