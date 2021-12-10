package stack_test

import (
	"github.com/aceakash/advent-of-code-2021/day10/stack"
	"testing"
)
import "github.com/stretchr/testify/assert"

const EmptyString = ""

func TestNewStackReturnsNothingWhenPopped(t *testing.T) {
	stringStack := stack.NewStringStack()

	item, ok := stringStack.Pop()
	assert.Equal(t, EmptyString, item)
	assert.False(t, ok)
}

func TestPushingAnItemThenPoppingIt_ReturnsTheItem(t *testing.T) {
	ss := stack.NewStringStack()

	ss.Push("first")
	popped, ok := ss.Pop()

	assert.True(t, ok)
	assert.Equal(t, "first", popped)
}

func TestPushingTwoItemThenPoppingTwice_ReturnsTheItemsInLifoOrder(t *testing.T) {
	ss := stack.NewStringStack()

	ss.Push("first")
	ss.Push("second")

	popped, ok := ss.Pop()
	assert.True(t, ok)
	assert.Equal(t, "second", popped)

	popped, ok = ss.Pop()
	assert.True(t, ok)
	assert.Equal(t, "first", popped)
}

func TestPushingAnItemThenPoppingTwice_ReturnsOnlyOneItem(t *testing.T) {
	ss := stack.NewStringStack()

	ss.Push("first")
	ss.Pop()

	_, ok := ss.Pop()
	assert.False(t, ok)
}

//func TestPushingAnItemAndPeekingIt(t *testing.T) {
////	s := New
////}



