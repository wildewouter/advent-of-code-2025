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

func PartTwo(input string) (int, error) {
	list := strings.Split(input, ",")

	sum := 0

	for _, value := range list {
		pair := strings.Split(value, "-")
		start, _ := strconv.Atoi(pair[0])
		end, _ := strconv.Atoi(pair[1])
		for i := start; i <= end; i++ {
			s := strconv.Itoa(i)

			for endI := 1; endI <= len(s)/2; endI++ {
				if isInvalidIdWithLength(s, endI) {
					sum += i
					break
				}
			}
		}
	}

	return sum, nil
}

func isInvalidIdWithLength(complete string, seqLength int) bool {
	if len(complete)%seqLength != 0 {
		return false
	}

	numParts := len(complete) / seqLength
	firstPart := complete[:seqLength]

	for i := 1; i < numParts; i++ {
		start := i * seqLength
		end := start + seqLength
		if complete[start:end] != firstPart {
			return false
		}
	}

	return true
}
