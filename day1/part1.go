package day1

import (
	"unicode"
)

func RunPart1(lines []string) {
	sum := 0

	for _, line := range lines {
		firstNum := 0
		lastNum := 0

		for _, char := range line {
			if unicode.IsDigit(char) {
				if firstNum == 0 {
					firstNum = int(char - '0')
				}
				lastNum = int(char - '0')
			}
		}

		sum += (firstNum * 10) + lastNum
	}

	println(sum)
}
