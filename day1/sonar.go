package day1

import (
	. "github.com/aceakash/advent-of-code-2021"
)

func PartOne(input string) int {
	readings := MustParseInputAsInts(input)

	count := 0
	for i := 1; i < len(readings); i++ {
		if readings[i] > readings[i-1] {
			count++
		}
	}
	return count
}

func PartTwo(input string) int {
	readings := MustParseInputAsInts(input)

	count := 0
	for i := 3; i < len(readings); i++ {
		prevWindow := readings[i-1] + readings[i-2] + readings[i-3]
		currWindow := readings[i] + readings[i-1] + readings[i-2]
		if currWindow > prevWindow {
			count++
		}
	}
	return count
}
