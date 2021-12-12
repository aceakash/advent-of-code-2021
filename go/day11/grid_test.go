package day11

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGrid_CanSetAValueAndGetItBack(t *testing.T) {
	grid := NewGrid(1, 1)

	grid.Set(0, 0, 99)

	assert.Equal(t, 99, grid.Get(0, 0))
}
