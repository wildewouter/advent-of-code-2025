package day06

import (
	"regexp"
	"strconv"
	"strings"
)

func PartOne(input string) (int, error) {
	lines := strings.Split(input, "\n")

	re := regexp.MustCompile(`\s+`)

	operatorLine := lines[len(lines)-1]
	operators := strings.Split(strings.TrimSpace(re.ReplaceAllString(operatorLine, " ")), " ")

	lines = lines[:len(lines)-1]

	var operations []Operation

	for _, operator := range operators {
		operations = append(operations, Operation{operator, []int{}})
	}
	operators = nil

	for _, line := range lines {
		stringValues := strings.Split(strings.TrimSpace(re.ReplaceAllString(line, " ")), " ")
		for i, value := range stringValues {
			v, _ := strconv.Atoi(value)
			operations[i].values = append(operations[i].values, v)
		}
	}

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

	return answer, nil
}

func PartTwo(input string) (int, error) {
	return 0, nil
}

type Operation struct {
	operator string
	values   []int
}
