package day10

import (
	"fmt"
	"strings"
)

func FindFirstWrongCloser(input string) string {
	if len(strings.TrimSpace(input)) == 0 {
		return ""
	}

	stack := NewStack()
	firstWrongCloser := ""
	for _, r := range input {
		fmt.Println(stack)
		sr := string(r)
		if isOpening(sr) {
			stack.Push(sr)
			continue
		}
		if isClosing(sr) {
			popped := stack.Pop()
			if popped != sr {
				firstWrongCloser = sr
				break
			}
			continue
		}
	}

	return firstWrongCloser
}

func isClosing(sr string) bool {
	return strings.Contains("}])>", sr)
}

func isOpening(sr string) bool {
	return strings.Contains("({[<", sr)
}



