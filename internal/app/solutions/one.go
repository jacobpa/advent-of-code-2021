package solutions

import (
	"strconv"
)

func DayOne(input []string) (int, int) {
	if len(input) < 2 {
		return 0, 0
	}
	intInput := convertInts(input)
	slidingWindowInput := make([]int, len(input)-2)

	for i, val := range intInput {
		if i < len(input) -2 {
			slidingWindowInput[i] = val + intInput[i+1] + intInput[i+2]
		}
	}

	return countIncreasing(intInput), countIncreasing(slidingWindowInput)
}

func countIncreasing(input []int) int {
	if len(input) < 2 {
		return 0
	}

	increased := 0
	prev := input[0]

	for _, s := range input[1:] {
		if s > prev {
			increased++
		}
		prev = s
	}

	return increased
}

func convertInts(input []string) []int {
	out := make([]int, len(input))
	for i, s := range input {
		j, _ := strconv.Atoi(s)
		out[i] = j
	}

	return out
}
