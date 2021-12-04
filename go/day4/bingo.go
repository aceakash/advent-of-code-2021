package day4

import (
	"fmt"
	"strings"

	. "github.com/aceakash/advent-of-code-2021"
)

func PartOne(input string) int {
	randomNums, boards := parseInput(input)

	var solvedBoard *Board = nil
	lastNumberPlayed := 0
	for _, rn := range randomNums {
		if solvedBoard != nil {
			break
		}
		for _, b := range boards {
			b.PlayNumber(rn)
			if b.IsSolved() {
				solvedBoard = b
				lastNumberPlayed = rn
				goto outOfLoop
			}
		}
	}
outOfLoop:
	if solvedBoard == nil {
		panic("no board was solved")
	}
	return solvedBoard.SumOfUnmarkedNumbers() * lastNumberPlayed
}

func parseInput(input string) ([]int, []*Board) {
	lines := strings.Split(input, "\n")

	nums := MustParseCsvOfInts(lines[0])

	boardStrings := strings.Split(strings.Join(lines[2:], "\n"), "\n\n")
	boards := make([]*Board, len(boardStrings))
	for i, bs := range boardStrings {
		b := NewBoard(bs)
		boards[i] = b
	}

	return nums, boards
}

func NewBoard(input string) *Board {
	cells := []Cell{}
	lines := strings.Split(input, "\n")
	//fmt.Println(len(lines))
	for _, l := range lines {
		//fmt.Println("line", l)
		l = CollapseSpaces(l)
		fmt.Println("line", l)
		l2 := strings.TrimSpace(strings.Join(strings.Split(l, " "), ","))
		fmt.Println("l2", l2)
		nums := MustParseCsvOfInts(l2)
		//fmt.Println("nums", nums)
		for _, n := range nums {
			cells = append(cells, Cell{n, false})
		}
	}

	return &Board{cells}
}

type Cell struct {
	num      int
	isMarked bool
}

type Board struct {
	cells []Cell
}

func (b *Board) PlayNumber(num int) {
	for i, c := range b.cells {
		if c.num == num {
			b.cells[i].isMarked = true
		}
	}
	for _, c := range b.cells {
		if c.isMarked {
			fmt.Println("===========", c)
		}
	}

}

func (b *Board) IsSolved() bool {
	return b.IsSolvedForAnyRows() || b.IsSolvedForAnyCols()
}

func (b *Board) IsSolvedForAnyRows() bool {
	// todo: generalise this to any size?
	for i := 0; i < 5; i++ {
		count := 0
		for j := 0; j < 5; j++ {
			idx := 5*i + j
			if b.cells[idx].isMarked {
				count++
			}
		}
		if count == 5 {
			return true
		}
	}
	return false
}

func (b *Board) IsSolvedForAnyCols() bool {
	// todo: generalise this to any size?
	for i := 0; i < 5; i++ {
		count := 0
		for j := 0; j < 5; j++ {
			idx := 5*j + i
			if b.cells[idx].isMarked {
				count++
			}
		}
		if count == 5 {
			return true
		}
	}
	return false
}

func (b *Board) SumOfUnmarkedNumbers() int {
	sum := 0
	for _, c := range b.cells {
		if !c.isMarked {
			sum += c.num
		}
	}
	return sum
}

func (b *Board) String() string {
	fmt.Println("b cells count", len(b.cells))
	str := ""
	for i, cell := range b.cells {
		if i % 5 == 0 {
			str += "\n"
		}
		decoratedNum := ""
		if cell.isMarked {
			decoratedNum = fmt.Sprintf("(%02d)", cell.num)
		} else {
			decoratedNum = fmt.Sprintf(" %s ", fmt.Sprintf("%02d", cell.num))
		}
		str = fmt.Sprintf("%s %s", str, decoratedNum)

	}
	return str
}
