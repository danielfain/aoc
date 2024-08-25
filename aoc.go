package main

import (
	"os"
	"strings"

	"fain.dev/aoc/day1"
)

func main() {
	args := os.Args[1:]

	day := args[0]
	part := args[1]

	buf, err := os.ReadFile(day + "/input.txt")
	if err != nil {
		panic("Could not read input file")
	}

	lines := strings.Split(string(buf), "\n")

	switch day {
	case "day1":
		if part == "part1" {
			day1.RunPart1(lines)
		} else {
			day1.RunPart2(lines)
		}
	}
}
