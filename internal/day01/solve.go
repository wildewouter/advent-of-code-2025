package day01

import (
	"strconv"
	"strings"
)

func Solve(input string) (int, error) {
	moves, err := getMoves(input)
	if err != nil {
		return 0, err
	}

	dial := 50
	zeroCount := 0

	for _, m := range moves {
		r := MoveResult{dial, 0}
		r = move(MoveResult{dial, zeroCount}, m)
		dial = r.newDial
		zeroCount = r.zeroCount
	}

	return zeroCount, nil
}

type MoveResult struct {
	newDial   int
	zeroCount int
}

func move(r MoveResult, m Move) MoveResult {
	d := (r.newDial + map[string]int{"L": -m.clicks, "R": m.clicks}[m.dir] + 100) % 100
	return MoveResult{d, r.zeroCount + bToI(d == 0)}
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
