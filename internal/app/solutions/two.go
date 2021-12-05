package solutions

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func DayTwo(input []string) (int, error) {
	x, y := 0, 0

	for _, s := range input {
		data := strings.Split(s, " ")
		dir := data[0]
		delta, _ := strconv.Atoi(data[1])

		switch dir {
		case "forward":
			x += delta
		case "up":
			y -= delta
		case "down":
			y += delta
		default:
			return 0, errors.New(fmt.Sprintf("Unrecognized direction %v", dir))
		}
	}

	return x * y, nil
}
