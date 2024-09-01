package day3

import (
	"strconv"
	"strings"
	"unicode"
)

func Part1(lines []string) {
	sum := 0

	for row, line := range lines {
		var curNumStr strings.Builder
		foundSymbol := false

		for col, char := range line {
			if unicode.IsDigit(char) {
				if hasAdjacentSymbol(lines, col, row) {
					foundSymbol = true
				}

				curNumStr.WriteRune(char)
			} else {
				if foundSymbol {
					curNum, _ := strconv.Atoi(curNumStr.String())
					sum += curNum
					foundSymbol = false
				}
				curNumStr.Reset()
			}
		}

		if foundSymbol {
			curNum, _ := strconv.Atoi(curNumStr.String())
			sum += curNum
		}
	}

	println(sum)
}

func hasAdjacentSymbol(lines []string, x, y int) bool {
	directions := [][]int{{-1, -1}, {0, -1}, {1, -1},
		{-1, 0}, {1, 0},
		{-1, 1}, {0, 1}, {1, 1}}

	for _, direction := range directions {
		dx := direction[0] + x
		dy := direction[1] + y

		if dx < 0 || dy < 0 || dx >= len(lines[0]) || dy >= len(lines) {
			continue
		}

		if isSymbol(lines[dy][dx]) {
			return true
		}
	}

	return false
}

func isSymbol(char byte) bool {
	return !unicode.IsDigit(rune(char)) && char != '.'
}
