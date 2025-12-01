package day01

import (
	"strconv"
	"strings"
)

func PartOne(input string) (int, error) {
	moves, err := getMoves(input)
	if err != nil {
		return 0, err
	}

	dial := 50
	zeroCount := 0

	for _, m := range moves {
		r := DialState{dial, 0}
		r = moveCountEndOnZero(DialState{dial, zeroCount}, m)
		dial = r.dial
		zeroCount = r.zeroCount
	}

	return zeroCount, nil
}

func PartTwo(input string) (int, error) {
	moves, err := getMoves(input)
	if err != nil {
		return 0, err
	}

	dial := 50
	zeroCount := 0

	for _, m := range moves {
		r := DialState{dial, 0}
		r = moveCountPassingZero(DialState{dial, zeroCount}, m)
		dial = r.dial
		zeroCount = r.zeroCount
	}

	return zeroCount, nil
}

type DialState struct {
	dial      int
	zeroCount int
}

func moveCountEndOnZero(r DialState, m Move) DialState {
	d := (r.dial + map[string]int{"L": -m.clicks, "R": m.clicks}[m.dir] + 100) % 100
	return DialState{d, r.zeroCount + bToI(d == 0)}
}

func moveCountPassingZero(r DialState, m Move) DialState {
	d := r.dial + map[string]int{"L": -m.clicks, "R": m.clicks}[m.dir]
	newDial := d % 100
	if newDial < 0 {
		newDial += 100
	}
	v := newDial != 0 && r.dial != 0 && map[string]bool{"L": newDial > r.dial, "R": newDial < r.dial}[m.dir]

	x := max(-m.clicks, m.clicks) / 100
	i := x + bToI(v) + bToI(newDial == 0)
	return DialState{newDial, r.zeroCount + i}
}

type Move struct {
	dir    string
	clicks int
}

func getMoves(input string) ([]Move, error) {
	var processedLines []Move
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		direction := line[0:1]
		distance, err := strconv.Atoi(line[1:])

		if err != nil {
			return nil, err
		}

		processedLines = append(processedLines, Move{
			dir:    direction,
			clicks: distance,
		})
	}

	return processedLines, nil
}

func bToI(b bool) int {
	if b {
		return 1
	}
	return 0
}
