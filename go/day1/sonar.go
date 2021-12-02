package day1

import (
	. "github.com/aceakash/advent-of-code-2021"
)

func PartOne(input string) int {
	readings := MustParseInputAsInts(input)
	return solveForWindowSize(readings, 1)
}

func PartTwo(input string) int {
	readings := MustParseInputAsInts(input)
	return solveForWindowSize(readings, 3)
}

func solveForWindowSize(readings []int, windowSize int) int {
	count := 0
	for i := windowSize; i < len(readings); i++ {
		prevWindow := sumWindow(readings, i-1, windowSize)
		currWindow := sumWindow(readings, i, windowSize)
		if currWindow > prevWindow {
			count++
		}
	}
	return count
}

func sumWindow(readings []int, lastIndex int, windowSize int) int {
	sum := 0
	for i := 0; i < windowSize; i++ {
		sum += readings[lastIndex-i]
	}
	return sum
}
