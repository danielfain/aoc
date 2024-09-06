package day4

import (
	"math"
	"slices"
	"strings"
)

func Part1(lines []string) {
	points := 0

	for _, line := range lines {
		matches := 0

		parts := strings.Split(strings.Split(line, ": ")[1], " | ")
		winningNumbers := strings.Fields(parts[0])
		numbers := strings.Fields(parts[1])

		for _, num := range numbers {
			if slices.Contains(winningNumbers, num) {
				matches++
			}
		}

		points += int(math.Pow(float64(2), float64(matches-1)))
	}

	println(points)
}
