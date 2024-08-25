package day1

import (
	"strings"
	"unicode"
)

func RunPart2(lines []string) {
	sum := 0

	for _, line := range lines {
		firstNum := 0
		lastNum := 0

		for _, char := range translateNumberWords(line) {
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

func translateNumberWords(line string) string {
	var sb strings.Builder

	nums := map[string]string{"one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9"}

	for i, char := range line {
		if unicode.IsDigit(char) {
			sb.WriteRune(char)
			continue
		}

		for word, num := range nums {
			if strings.HasPrefix(line[i:], word) {
				sb.WriteString(num)
			}
		}
	}

	return sb.String()
}
