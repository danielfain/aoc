package main

import (
	"os"
	"strings"

	"fain.dev/aoc/day1"
	"fain.dev/aoc/day2"
	"fain.dev/aoc/day3"
	"fain.dev/aoc/day4"
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
	lines = lines[:len(lines)-1]

	switch day {
	case "day1":
		if part == "part1" {
			day1.RunPart1(lines)
		} else {
			day1.RunPart2(lines)
		}
	case "day2":
		if part == "part1" {
			day2.RunPart1(lines)
		} else {
		}
	case "day3":
		if part == "part1" {
			day3.Part1(lines)
		}
	case "day4":
		if part == "part1" {
			day4.Part1(lines)
		} else {
			day4.Part2(lines)
		}
	}
}
