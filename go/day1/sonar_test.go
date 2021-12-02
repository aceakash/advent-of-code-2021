package day1_test

import (
	. "github.com/aceakash/advent-of-code-2021"
	"github.com/aceakash/advent-of-code-2021/day1"
	"testing"
)

func TestPartOne(t *testing.T) {
	testData := []struct {
		scenario string
		input    string
		want     int
	}{
		{`example`, `199
200
208
210
200
207
240
269
260
263`, 7},

		{`prob to solve`, MustReadFile(), 1553},
	}

	for _, td := range testData {
		t.Run(td.scenario, func(t *testing.T) {
			got := day1.PartOne(td.input)
			if got != td.want {
				t.Errorf("Wanted %d, got %d for %v", td.want, got, td.input)
			}
		})
	}
}

func TestPartTwo(t *testing.T) {
	testData := []struct {
		scenario string
		input    string
		want     int
	}{
		{`example`, `199
200
208
210
200
207
240
269
260
263`, 5},

		{`prob to solve`, MustReadFile(), 1597},
	}

	for _, td := range testData {
		t.Run(td.scenario, func(t *testing.T) {
			got := day1.PartTwo(td.input)
			if got != td.want {
				t.Errorf("Wanted %d, got %d for %v", td.want, got, td.input)
			}
		})
	}
}
