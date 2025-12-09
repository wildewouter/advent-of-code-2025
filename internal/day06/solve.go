package day06

import (
	"regexp"
	"strconv"
	"strings"
)

var re = regexp.MustCompile(`\s+`)

func PartOne(input string) (int, error) {
	lines, operations := getOperations(input)

	for _, line := range lines {
		stringValues := strings.Split(strings.TrimSpace(re.ReplaceAllString(line, " ")), " ")
		for i, value := range stringValues {
			v, _ := strconv.Atoi(value)
			operations[i].values = append(operations[i].values, v)
		}
	}

	answer := calc(operations)

	return answer, nil
}

func PartTwo(input string) (int, error) {
	lines, operations := getOperations(input)

	runeGrid := map[Coord]rune{}

	for y, line := range lines {
		runes := []rune(line)

		for x, r := range runes {
			runeGrid[Coord{x: x, y: y}] = r
		}
	}

	var groups [][]int

	series := 0

	for x := 0; x <= len(lines[0]); x++ {
		values := ""
		for y := 0; y < len(lines); y++ {
			r := runeGrid[Coord{x, y}]

			if r == ' ' {
				continue
			}

			values += string(r)
		}

		if values == "" {
			series++
			continue
		}

		if len(groups) <= series {
			groups = append(groups, []int{})
		}

		atoi, _ := strconv.Atoi(values)
		groups[series] = append(groups[series], atoi)
	}

	for i, _ := range operations {
		for _, v := range groups[i] {
			operations[i].values = append(operations[i].values, v)
		}
	}

	answer := calc(operations)

	return answer, nil
}

type Coord struct {
	x int
	y int
}

func getOperations(input string) ([]string, []Operation) {
	lines := strings.Split(input, "\n")

	operatorLine := lines[len(lines)-1]
	operators := strings.Split(strings.TrimSpace(re.ReplaceAllString(operatorLine, " ")), " ")

	lines = lines[:len(lines)-1]

	var operations []Operation

	for _, operator := range operators {
		operations = append(operations, Operation{operator, []int{}})
	}
	operators = nil
	return lines, operations
}

func calc(operations []Operation) int {
	answer := 0

	for _, operation := range operations {
		opAns := operation.values[0]

		for _, v := range operation.values[1:] {
			switch operation.operator {
			case "+":
				opAns += v
			case "*":
				opAns *= v
			}
		}

		answer += opAns
	}

	return answer
}

type Operation struct {
	operator string
	values   []int
}
