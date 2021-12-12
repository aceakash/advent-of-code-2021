package day11

/*
1234
5678

width = 4, height = 2

7 is at 2, 1 at at cells[6]
 */

type Grid struct {
	width, height int
	cells         []int
}

func (g *Grid) getCellIndex(x, y int) int {
	return y*g.width + x
}

func (g *Grid) Set(x int, y int, value int) {
	g.cells[g.getCellIndex(x, y)] = value
}

func (g *Grid) Get(x int, y int) int {
	return g.cells[g.getCellIndex(x, y)]
}

func NewGrid(width int, height int) *Grid {
	return &Grid{
		width:  width,
		height: height,
		cells:  make([]int, width*height),
	}
}
