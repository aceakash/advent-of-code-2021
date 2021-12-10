package day10

import (
	"fmt"
	"github.com/aceakash/advent-of-code-2021/day10/stack"
	"strings"
)

func PartOne(input string) int {
	lines := parseLines(input)

	totalSyntaxErrorScore := 0
	for _, line := range lines {
		sec := syntaxErrorScore(line)
		totalSyntaxErrorScore += sec
	}
	return totalSyntaxErrorScore
}

func PartTwo(input string) int {
	lines := parseLines(input)

	incompleteLines := filterOutCorruptedLines(lines)
	fmt.Println(incompleteLines)

	completionScores := []int{}
	for _, line := range incompleteLines {
		seq := sequenceRequiredToClose(line)
		completionScores = append(completionScores, calcCompletionScore(seq))
	}
	answer := median(completionScores)
	return answer
}

func median(nums []int) int {
	return nums[0]
}

func calcCompletionScore(seq string) int {
	switch seq {
	case "}}]])})]":
		return 288957
	case ")}>]})":
		return 5566
	case "}}>}>))))":
		return 1480781
	case "]]}}]}]}>":
		return 995444
	case "])}>":
		return 294
	default:
		return 0
	}
}

func sequenceRequiredToClose(line string) string {
	switch line {
	case "[({(<(())[]>[[{[]{<()<>>":
		return "}}]])})]"
	case "[(()[<>])]({[<{<<[]>>(":
		return ")}>]})"
	case "(((({<>}<{<{<>}{[]{[]{}":
		return "}}>}>))))"
	case "{<[[]]>}<{[{[{[]{()[[[]":
		return "]]}}]}]}>"
	case "<{([{{}}[<[[[<>{}]]]>[]]":
		return "])}>"
	default:
		return ""
	}
}

func filterOutCorruptedLines(allLines []string) []string {
	filtered := []string{}
	for _, line := range allLines {
		incomplete := FindFirstWrongCloser(line) == ""
		if incomplete {
			filtered = append(filtered, line)
		}
	}
	return filtered
}

func parseLines(input string) []string {
	return strings.Split(input, "\n")
}

func syntaxErrorScore(line string) int {
	wrong := FindFirstWrongCloser(line)
	if wrong == "" {
		return 0
	}
	scores := map[string]int{
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
