package solutions

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func DayTwo(input []string) (int, int, error) {
	ax, ay := 0, 0
	bx, by, aim := 0, 0, 0

	for _, s := range input {
		data := strings.Split(s, " ")
		dir := data[0]
		delta, _ := strconv.Atoi(data[1])

		switch dir {
		case "forward":
			ax += delta
			bx += delta
			by += aim * delta
		case "up":
			ay -= delta
			aim -= delta
		case "down":
			ay += delta
			aim += delta
		default:
			return 0, 0, errors.New(fmt.Sprintf("Unrecognized direction %v", dir))
		}
	}

	return ax * ay, bx * by, nil
}
