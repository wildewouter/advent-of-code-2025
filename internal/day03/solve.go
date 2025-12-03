package day03

import (
	"strconv"
	"strings"
)

func PartOne(input string) (int, error) {
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

		result := MaxWithIndex(parsedBank[:len(parsedBank)-1])
		result2 := MaxWithIndex(parsedBank[result.index+1:])

		joltage, err := strconv.Atoi(strconv.Itoa(result.max) + strconv.Itoa(result2.max))

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
