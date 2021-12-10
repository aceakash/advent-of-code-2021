package day10

import (
	"fmt"
	"github.com/aceakash/advent-of-code-2021/day10/stack"
	"strings"
)

func FindFirstWrongCloser(input string) string {
	fmt.Println(input)
	if len(strings.TrimSpace(input)) == 0 {
		return ""
	}

	stringStack := stack.NewStringStack()
	fmt.Println("stringStack when created", stringStack)
	firstWrongCloser := ""
	for _, r := range input {
		sr := string(r)
		fmt.Println("sr", sr)
		fmt.Println("stringStack", stringStack)
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



