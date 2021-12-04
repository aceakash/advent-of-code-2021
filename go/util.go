package advent_of_code_2021

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func MustReadFile() string {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}
	return strings.TrimSpace(string(b))
}

func MustParseInt(s string) int {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return int(i)
}

func MustParseInputAsInts(input string) []int {
	lines := strings.Split(input, "\n")
	readings := make([]int, len(lines))
	for i, line := range lines {
		readings[i] = MustParseInt(line)
	}
	return readings
}

func MustParseCsvOfInts(csv string) []int {
	numStrs := strings.Split(csv, ",")

	nums := make([]int, len(numStrs))
	for i, v := range numStrs {
		nums[i] = MustParseInt(v)
	}
	return nums
}

func CollapseSpaces(input string) string {
	r := regexp.MustCompile(`[ ]+`)
	return strings.TrimSpace(r.ReplaceAllString(input, " "))
}
