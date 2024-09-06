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

func Part2(lines []string) {
	cardsCopies := make(map[int]int)

	for cardNum, line := range lines {
		matches := 0

		parts := strings.Split(strings.Split(line, ": ")[1], " | ")
		winningNumbers := strings.Fields(parts[0])
		numbers := strings.Fields(parts[1])

		for _, num := range numbers {
			if slices.Contains(winningNumbers, num) {
				matches++
			}
		}

		copies, _ := cardsCopies[cardNum]
		copies++
		cardsCopies[cardNum] = copies

		for i := 0; i < copies; i++ {
			for j := 1; j <= matches; j++ {
				c, _ := cardsCopies[cardNum+j]
				cardsCopies[cardNum+j] = c + 1
			}
		}
	}

	cards := 0
	for _, copies := range cardsCopies {
		cards += copies
	}

	println(cards)
}
