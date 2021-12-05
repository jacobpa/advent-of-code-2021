# Advent of Code 2021
Collection of solutions for Advent of Code 2021, written in Go. Input can be provided via `STDIN` or as a text file.

## Building
```shell
go build advent-of-code-2021/cmd/aoc
```

## Usage
```shell
aoc CHALLENGE [INPUT]

CHALLENGE    Which day's challenge to complete
INPUT        Input filename, if not provided reads from STDIN
```

### Examples
```shell
aoc 1 day1.txt
cat day1.txt | aoc 1
```