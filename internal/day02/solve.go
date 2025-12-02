package day02

import (
	"strconv"
	"strings"
)

func PartOne(input string) (int, error) {
	list := strings.Split(input, ",")

	sum := 0
	for _, value := range list {
		pair := strings.Split(value, "-")
		start, _ := strconv.Atoi(pair[0])
		end, _ := strconv.Atoi(pair[1])
		for i := start; i <= end; i++ {
			s := strconv.Itoa(i)
			left := s[:len(s)/2]
			right := s[len(s)/2:]

			if left == right {
				sum += i
			}
		}
	}

	return sum, nil
}
