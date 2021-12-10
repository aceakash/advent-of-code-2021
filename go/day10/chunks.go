package day10

import (
	"github.com/aceakash/advent-of-code-2021/day10/stack"
	"strings"
)

func PartOne(input string) int {
	lines := parseLines(input)

	totalSyntaxErrorScore := 0
	for _, line := range(lines) {
		sec := syntaxErrorScore(line)
		totalSyntaxErrorScore += sec
	}
	return totalSyntaxErrorScore
}

func parseLines(input string) []string {
	return strings.Split(input, "\n")
}

func syntaxErrorScore(line string) int {
	wrong := FindFirstWrongCloser(line)
	if wrong == "" {
		return 0
	}
	scores := map[string]int {
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}
	return scores[wrong]
}

func FindFirstWrongCloser(line string) string {
	if len(strings.TrimSpace(line)) == 0 {
		return ""
	}
	stringStack, firstWrongCloser := stack.NewStringStack(), ""
	for _, r := range line {
		sr := string(r)
		if isOpening(sr) {
			stringStack.Push(sr)
			continue
		}
		if isClosing(sr) {
			popped, ok := stringStack.Pop()
			if !ok {
				firstWrongCloser = sr
				break
			}
			if !isClosingFor(sr, popped) {
				firstWrongCloser = sr
				break
			}
		}
	}
	return firstWrongCloser
}

// ----
var openingToClosingMap = map[string]string{
	"(": ")",
	"{": "}",
	"<": ">",
	"[": "]",
}

func isClosing(sr string) bool {
	for _, v := range openingToClosingMap {
		if v == sr {
			return true
		}
	}
	return false
}

func isClosingFor(closing string, opening string) bool {
	c, ok := openingToClosingMap[opening]
	if ok && c == closing {
		return true
	}
	return false
}

func isOpening(sr string) bool {
	for k, _ := range openingToClosingMap {
		if k == sr {
			return true
		}
	}
	return false
}
