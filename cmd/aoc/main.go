package main

import (
	"advent-of-code-2021/internal/app/solutions"
	"advent-of-code-2021/internal/pkg/common"
	"fmt"
	"os"
)

const usage = `%s CHALLENGE [INPUT]

CHALLENGE    Which day's challenge to complete
INPUT        Input filename, if not provided reads from STDIN`

func main() {
	var input []string
	var err error
	if len(os.Args) < 2 {
		fmt.Printf("Must provide CHALLENGE argument\n")
		fmt.Printf(usage, os.Args[0])
		fmt.Println()
		os.Exit(1)
	}

	if len(os.Args) == 3 {
		input, err = common.ParseFile(os.Args[2])
	} else {
		input, err = common.ParseStdin()
	}
	if err != nil {
		fmt.Printf("Error parsing input: %v", err)
		os.Exit(2)
	}

	switch os.Args[1] {
	case "1":
		solution := solutions.DayOne(input)
		fmt.Printf("Day 1's solution is: %v\n", solution)
	case "2":
		solution, err := solutions.DayTwo(input)
		if err != nil {
			fmt.Printf("Error finding Day Two solution: %v", err)
			os.Exit(3)
		}
		fmt.Printf("Day 1's solution is: %v\n", solution)
	default:
		fmt.Printf("No solution implemented for challenge %v\n", os.Args[1])
	}
}
