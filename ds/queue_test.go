package ds_test

import (
	"testing"

	"github.com/nmmugia/dsa/ds"
	"github.com/stretchr/testify/assert"
)

/*
Queue requirements:

Enqueueing: The queue should be able to add new elements to the back of the queue.
	-Test Case 1: Test that the size of the queue increases by 1 after enqueueing an element.

Dequeueing: The queue should be able to remove elements from the front of the queue.
	- Test Case 1: Test that the size of the queue decreases by 1 after dequeueing an element.
	- Test Case 2: Test that the element removed from the front of the queue is the first element.

Peek: The queue should be able to return the first element without removing it.
	- Test Case 1: Test that peeking at the queue returns the first element.
	- Test Case 2: Test that the size of the queue remains the same after peeking.

Empty queue: The queue should handle being empty correctly.
	- Test Case 1: Test that dequeueing from an empty queue returns an error.
	- Test Case 2: Test that peeking at an empty queue returns an error.

Capacity: The queue should have a capacity limit.
	- Test Case 1: Test that adding elements to the queue up to its capacity does not return an error.
	- Test Case 2: Test that adding elements to the queue beyond its capacity returns an error.

Clear: The queue should be able to remove all elements and become empty.
	- Test Case 1: Test that clearing the queue removes all elements.
	- Test Case 2: Test that the size of the queue is 0 after clearing.
*/

func TestEnqueueing(t *testing.T) {
	t.Run("Test that the size of the queue increases by 1 after enqueueing an element.",
		func(t *testing.T) {
			q := ds.NewQueue(5)
			assert.Nil(t, q.Enqueueing(1))
			assert.Equal(t, 1, q.Size())
			assert.Nil(t, q.Enqueueing("kamadssa"))
			assert.Equal(t, 2, q.Size())
		})
}

func TestPeek(t *testing.T) {
	t.Run("Test that the size of the queue decreases by 1 after dequeueing an element.",
		func(t *testing.T) {
			q := ds.NewQueue(5)
			assert.Nil(t, q.Enqueueing(1))
			assert.Nil(t, q.Enqueueing("sada"))
			assert.Equal(t, 2, q.Size())
			assert.Equal(t, 1, q.Dequeueing())
			assert.Equal(t, 1, q.Size())

		})
	t.Run("Test that the element removed from the front of the queue is the first element.",
		func(t *testing.T) {
			q := ds.NewQueue(5)
			assert.Nil(t, q.Enqueueing(1))
			assert.Equal(t, 1, q.Size())
			assert.Nil(t, q.Enqueueing("sada"))
			assert.Equal(t, 2, q.Size())
			assert.Equal(t, 1, q.Dequeueing())
			assert.Equal(t, 1, q.Size())
			assert.Equal(t, "sada", q.Dequeueing())
			assert.Equal(t, 0, q.Size())
		})
}

func TestDequeueing(t *testing.T) {
	t.Run("Test that peeking at the queue returns the first element.",
		func(t *testing.T) {
			q := ds.NewQueue(5)
			assert.Nil(t, q.Enqueueing(1))
			assert.Nil(t, q.Enqueueing("sada"))
			assert.Equal(t, 1, q.Peek())
		})
	t.Run("Test that the size of the queue remains the same after peeking.",
		func(t *testing.T) {
			q := ds.NewQueue(5)
			assert.Nil(t, q.Enqueueing(1))
			assert.Nil(t, q.Enqueueing("sada"))
			assert.Equal(t, 1, q.Peek())
			assert.Equal(t, 2, q.Size())
		})
}

func TestEmptyQueue(t *testing.T) {
	t.Run("Test that dequeueing from an empty queue returns an error.",
		func(t *testing.T) {
			q := ds.NewQueue(5)
			assert.Equal(t, ds.ErrNoSuchElement, q.Dequeueing())
		})
	t.Run("Test that peeking at an empty queue returns an error.",
		func(t *testing.T) {
			q := ds.NewQueue(5)
			assert.Equal(t, ds.ErrNoSuchElement, q.Peek())
		})
}

func TestCapacity(t *testing.T) {
	t.Run("Test that adding elements to the queue up to its capacity does not return an error.",
		func(t *testing.T) {
			q := ds.NewQueue(2)
			assert.Nil(t, q.Enqueueing(1))
			assert.Nil(t, q.Enqueueing("sada"))
		})
	t.Run("Test that adding elements to the queue beyond its capacity returns an error.",
		func(t *testing.T) {
			q := ds.NewQueue(2)
			assert.Nil(t, q.Enqueueing(1))
			assert.Nil(t, q.Enqueueing("sada"))
			assert.Equal(t, ds.ErrCapacityLimit, q.Enqueueing("sada"))
		})
}

func TestClear(t *testing.T) {
	t.Run("Test that clearing the queue removes all elements.",
		func(t *testing.T) {
			q := ds.NewQueue(2)
			assert.Nil(t, q.Enqueueing(1))
			assert.Nil(t, q.Enqueueing("sada"))
			assert.Nil(t, q.Clear())
			assert.Equal(t, ds.ErrNoSuchElement, q.Dequeueing())
		})
	t.Run("Test that the size of the queue is 0 after clearing.",
		func(t *testing.T) {
			q := ds.NewQueue(2)
			assert.Nil(t, q.Enqueueing(1))
			assert.Nil(t, q.Enqueueing("sada"))
			assert.Nil(t, q.Clear())
			assert.Equal(t, 0, q.Size())
		})
}
