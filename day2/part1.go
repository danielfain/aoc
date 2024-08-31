package day2

import (
	"strconv"
	"strings"
)

func RunPart1(games []string) {
	sum := 0

	for id, game := range games {
		rounds := strings.Split(game, ": ")[1]

		maxRed := 0
		maxGreen := 0
		maxBlue := 0

		for _, round := range strings.Split(rounds, "; ") {
			for _, hand := range strings.Split(round, ", ") {
				parts := strings.Split(hand, " ")
				count, _ := strconv.Atoi(parts[0])
				color := strings.TrimSpace(parts[1])

				if color == "red" {
					maxRed = max(count, maxRed)
				} else if color == "green" {
					maxGreen = max(count, maxGreen)
				} else if color == "blue" {
					maxBlue = max(count, maxBlue)
				}
			}
		}

		if maxRed <= 12 && maxGreen <= 13 && maxBlue <= 14 {
			sum += id + 1
		}
	}

	println(sum)
}
