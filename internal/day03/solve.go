package day03

import (
	"strconv"
	"strings"
)

func PartOne(input string) (int, error) {
	return doIt(input, 2)
}

func PartTwo(input string) (int, error) {
	return doIt(input, 12)
}

func doIt(input string, amount int) (int, error) {
	lines := strings.Split(input, "\n")

	var banksJoltages []int

	for _, bank := range lines {
		var parsedBank []int
		batteries := strings.Split(bank, "")
		for _, battery := range batteries {
			v, err := strconv.Atoi(battery)
			if err != nil {
				return 0, err
			}
			parsedBank = append(parsedBank, v)
		}

		var joltages []string
		position := 0
		for i := 0; i < amount; i++ {
			end := len(parsedBank) - amount + len(joltages) + 1
			result := MaxWithIndex(parsedBank[position:end])
			joltages = append(joltages, strconv.Itoa(result.max))
			position = position + result.index + 1
		}

		result := ""

		for _, j := range joltages {
			result += j
		}

		joltage, err := strconv.Atoi(result)

		if err != nil {
			return 0, err
		}

		banksJoltages = append(banksJoltages, joltage)
	}

	return Sum(banksJoltages), nil
}

type MaxResult struct {
	max   int
	index int
}

func MaxWithIndex(arr []int) MaxResult {
	maxi := MaxResult{arr[0], 0}
	for i, v := range arr {
		if v > maxi.max {
			maxi = MaxResult{v, i}
		}
	}
	return maxi
}

func Sum(arr []int) int {
	var sum int
	for _, v := range arr {
		sum += v
	}
	return sum
}
