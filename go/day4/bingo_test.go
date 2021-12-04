package day4

import (
	"fmt"
	"testing"

	. "github.com/aceakash/advent-of-code-2021"
	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	actual := PartOne(MustReadFile())

	assert.Equal(t, 28082, actual)
}

func TestPartTwo(t *testing.T) {
	actual := PartTwo(MustReadFile())

	assert.Equal(t, 8224, actual)
}

func TestNewBoard_CanBePrinted(t *testing.T) {
	board := NewBoard(`0 1 2 3 4
5 6 7 8 9
10 11 12 13 14
15 16 17 18 19
20 21 22 23 24`)

	expected := `
  00   01   02   03   04 
  05   06   07   08   09 
  10   11   12   13   14 
  15   16   17   18   19 
  20   21   22   23   24 `

	assert.Equal(t, expected, board.String())
}

func TestNewBoard_CanBePrinted_WithSomeCellsmarked(t *testing.T) {
	board := NewBoard(`0 1 2 3 4
5 6 7 8 9
10 11 12 13 14
15 16 17 18 19
20 21 22 23 24`)

	board.PlayNumber(2)
	fmt.Println(board)

	expected := `
  00   01  (02)  03   04 
  05   06   07   08   09 
  10   11   12   13   14 
  15   16   17   18   19 
  20   21   22   23   24 `

	assert.Equal(t, expected, board.String())
}

func TestBoard_IsSolvedForAnyCols_ReturnsTrueWhenSolvedForFirstColumn(t *testing.T) {
	board := NewBoard(`0 1 2 3 4
5 6 7 8 9
10 11 12 13 14
15 16 17 18 19
20 21 22 23 24`)

	board.PlayNumber(0)
	board.PlayNumber(5)
	board.PlayNumber(10)
	board.PlayNumber(15)
	board.PlayNumber(20)

	fmt.Println(board)
	assert.Equal(t, true, board.IsSolvedForAnyCols())
}

func TestBoard_IsSolvedForAnyCols_ReturnsFalseWhenNotSolvedForAnyColumn(t *testing.T) {
	board := NewBoard(`0 1 2 3 4
5 6 7 8 9
10 11 12 13 14
15 16 17 18 19
20 21 22 23 24`)

	board.PlayNumber(0)
	board.PlayNumber(5)
	board.PlayNumber(10)
	board.PlayNumber(15)
	board.PlayNumber(21)

	fmt.Println(board)
	assert.Equal(t, false, board.IsSolvedForAnyCols())
}

func TestRegex(t *testing.T) {
	input := "1   2 3   4   5"
	exp := "1 2 3 4 5"

	assert.Equal(t, exp, CollapseSpaces(input))

}


