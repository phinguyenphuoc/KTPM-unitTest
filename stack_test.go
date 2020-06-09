package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEmpty(t *testing.T) {
	t.Run("Test stack isEmpty return true", func(t *testing.T) {
		var stack Stack
		assert.Equal(t, stack.IsEmpty(), true)
	})
	t.Run("Test stack isEmpty return false", func(t *testing.T) {
		var stack Stack
		stack.Push("1")
		assert.Equal(t, stack.IsEmpty(), false)
	})
}

func TestPush(t *testing.T) {
	t.Run("Test push one time success", func(t *testing.T) {
		var stack Stack
		assert.Equal(t, stack.Push("1"), "1")
		assert.Equal(t, stack.IsEmpty(), false)
	})
	t.Run("Test push two time success", func(t *testing.T) {
		var stack Stack
		stack.Push("1")
		assert.Equal(t, stack.Push("2"), "2")
		assert.Equal(t, stack.IsEmpty(), false)
	})
	t.Run("Test push after pop success", func(t *testing.T) {
		var stack Stack
		stack.Push("1")
		stack.Pop()
		assert.Equal(t, stack.Push("2"), "2")
		assert.Equal(t, stack.IsEmpty(), false)
	})
}

func TestPop(t *testing.T) {
	t.Run("Test pop when stack empty", func(t *testing.T) {
		var stack Stack
		assert.Equal(t, stack.IsEmpty(), true)
		_, success := stack.Pop()
		assert.Equal(t, success, false)
	})
	t.Run("Test pop when stack is not empty", func(t *testing.T) {
		var stack Stack
		assert.Equal(t, stack.IsEmpty(), true)
		stack.Push("1")
		assert.Equal(t, stack.IsEmpty(), false)
		value, success := stack.Pop()
		assert.Equal(t, success, true)
		assert.Equal(t, value, "1")
	})
}

func TestPeek(t *testing.T) {
	t.Run("Peek return right value after push", func(t *testing.T) {
		var stack Stack
		stack.Push("1")
		assert.Equal(t, stack.Peek(), "1")
	})
	t.Run("Peek return -1 when stack is empty", func(t *testing.T) {
		var stack Stack
		assert.Equal(t, stack.Peek(), "-1")
	})
}

func TestClear(t *testing.T) {
	t.Run("Clear when stack is Empty", func(t *testing.T) {
		var stack Stack
		assert.Equal(t, stack.IsEmpty(), true)
		stack.Clear()
		assert.Equal(t, stack.Peek(), "-1")
		assert.Equal(t, stack.IsEmpty(), true)
	})
	t.Run("Clear when stack is not Empty", func(t *testing.T) {
		var stack Stack
		stack.Push("1")
		assert.Equal(t, stack.IsEmpty(), false)
		stack.Clear()
		assert.Equal(t, stack.Peek(), "-1")
		assert.Equal(t, stack.IsEmpty(), true)
	})
}
