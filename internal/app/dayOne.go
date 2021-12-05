package app

import "strconv"

func DayOne(input []string) int {
	increased := 0
	if len(input) < 2 {
		return increased
	}
	prev, _ := strconv.Atoi(input[0])

	for _, s := range input[1:] {
		val, _ := strconv.Atoi(s)
		if val > prev {
			increased++
		}
		prev = val
	}

	return increased
}