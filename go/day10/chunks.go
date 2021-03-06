package day10

import (
	"github.com/aceakash/advent-of-code-2021/day10/stack"
	"sort"
	"strings"
)

var openingToClosingMap = map[string]string{
	"(": ")",
	"{": "}",
	"<": ">",
	"[": "]",
}

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

	completionScores := make([]int, len(incompleteLines))
	for i, line := range incompleteLines {
		seq := sequenceRequiredToClose(line)
		completionScores[i] = calcCompletionScore(seq)
	}
	return median(completionScores)
}

func median(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	sort.Ints(nums)
	midIndex := (len(nums) + 1) / 2
	return nums[midIndex]
}

func calcCompletionScore(seq string) int {
	values := map[rune]int {
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}
	total := 0
	for _, r := range seq {
		total = (5 * total) + values[r]
	}
	return total
}

func sequenceRequiredToClose(line string) string {
	/*
		continuously remove all () {} [] <>  till none remain
		[({(<(())[]>[[{[]{<()<>>   --->   [({([[{{

		then just close all open ones (starting from the top of the stack or end of the string)
	*/
	for {
		newLine := removePairs(line)
		if len(newLine) == len(line) {
			break
		}
		line = newLine
	}
	seq := findSimpleClosingSeq(line)
	return seq
}

func findSimpleClosingSeq(line string) string {
	seq := ""
	for _, r := range line {
		seq = openingToClosingMap[string(r)] + seq
	}
	return seq
}

func removePairs(line string) string {
	s1 := strings.ReplaceAll(line, "{}", "")
	s2 := strings.ReplaceAll(s1, "[]", "")
	s3 := strings.ReplaceAll(s2, "<>", "")
	s4 := strings.ReplaceAll(s3, "()", "")
	return s4
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
