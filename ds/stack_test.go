package ds_test

import (
	"testing"

	"github.com/nmmugia/dsa/ds"
	"github.com/stretchr/testify/assert"
)

/*
Stack
Requirements:
- A stack is empty on init
- A stack has size of 0 on construction
- After n time pushes to an empty stack (n > 0), the stack is non empty and the size of the stack is n
- if one pushes x then peeks, the value returned is x, but the size stays the same
- peeking into an empty stack return error: ErrNoSuchElement
- if one pushes x then pops, the value popped is x, and the size is one less than the old size
- if the size is n, then after n pops, the stack is empty and has zero size
- popping from an empty stack return error: ErrNoSuchElement
*/

func TestNewSet(t *testing.T) {

	t.Run("A stack is empty on construction", func(t *testing.T) {
		s := ds.NewStack()
		assert.True(t, s.IsEmpty())
	})

	t.Run("A stack has size of 0 on construction", func(t *testing.T) {
		s := ds.NewStack()
		assert.Equal(t, 0, s.Size())
	})

}

func TestInsert(t *testing.T) {
	t.Run(" After n time pushes to an empty stack (n > 0), the stack is non empty and the size of the stack is n",
		func(t *testing.T) {
			s := ds.NewStack()
			assert.Nil(t, s.Push(1))
			assert.Nil(t, s.Push("assdad"))
			assert.Nil(t, s.Push(3))
			assert.False(t, s.IsEmpty())
			assert.Equal(t, 3, s.Size())
		})

	t.Run("if one pushes x then peeks, the value returned is x, but the size stays the same",
		func(t *testing.T) {
			s := ds.NewStack()
			assert.Nil(t, s.Push("as"))
			assert.Equal(t, "as", s.Peek())
			assert.Equal(t, 1, s.Size())
			assert.Nil(t, s.Push(1))
			assert.Equal(t, 1, s.Peek())
			assert.Equal(t, 2, s.Size())
		})

	t.Run("peeking into an empty stack return error: ErrNoSuchElement",
		func(t *testing.T) {
			s := ds.NewStack()
			assert.Equal(t, ds.ErrNoSuchElement, s.Peek())
		})

	t.Run("if one pushes x then pops, the value popped is x, and the size is one less than the old size",
		func(t *testing.T) {
			s := ds.NewStack()
			assert.Nil(t, s.Push(1))
			assert.Nil(t, s.Push(2))
			assert.Equal(t, 2, s.Size())
			assert.Equal(t, 2, s.Pop())
			assert.Equal(t, 1, s.Size())
			assert.Equal(t, 1, s.Pop())
		})
	t.Run("if the size is n, then after n pops, the stack is empty and has zero size",
		func(t *testing.T) {
			s := ds.NewStack()
			assert.Nil(t, s.Push(1))
			assert.Equal(t, 1, s.Size())
			assert.Equal(t, 1, s.Pop())
			assert.Equal(t, 0, s.Size())
		})
	t.Run("popping from an empty stack return error: ErrNoSuchElement",
		func(t *testing.T) {
			s := ds.NewStack()
			assert.Equal(t, ds.ErrNoSuchElement, s.Pop())
		})
}
