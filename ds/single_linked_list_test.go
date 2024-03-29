package ds_test

import (
	"testing"

	"github.com/nmmugia/dsa/ds"
	"github.com/stretchr/testify/assert"
)

/*
Requirements:

Test the creation of an empty list: Ensure that a new instance of the Single Linked List data structure can be created and that it starts with zero elements.

Test adding elements to the list: Verify that elements can be added to the list, one at a time or in bulk, and that the size of the list reflects the number of elements added.

Test removing elements from the list: Confirm that elements can be removed from the list, one at a time or in bulk, and that the size of the list reflects the number of elements removed.

Test retrieving elements from the list: Test that elements can be retrieved from the list by index or by value, and that the correct element is returned.

Test the size of the list: Verify that the size of the list is correctly updated as elements are added or removed.

Test searching for elements in the list: Check that the list can search for elements by value, returning the correct element or indicating that it is not present in the list.

Test inserting elements at a specific index: Test that elements can be inserted at a specific index in the list, and that the size of the list reflects the addition of the new element.

Test modifying elements in the list: Verify that elements in the list can be modified, and that the modified element is correctly reflected in the list.

Test clearing the list: Confirm that the list can be cleared of all elements, and that the size of the list is zero after clearing.

Test handling edge cases: Test the behavior of the list when it is empty, has only one element, or has multiple elements. Test the behavior of the list when attempting to add, remove, or retrieve elements that are out of bounds or not present in the list.
*/

func TestCreationSingleLinkedList(t *testing.T) {
	t.Run(`Test the creation of an empty list: Ensure that a new instance of the Single Linked\n
		List data structure can be created and that it starts with zero elements.`,
		func(t *testing.T) {
			sl := ds.NewSingleLinkedList()
			assert.True(t, sl.IsEmpty())
			assert.Equal(t, 0, sl.Size())
		})
}

func TestManipulationSingleLinkedList(t *testing.T) {
	t.Run(`Test adding elements to the list: Verify that elements can be added
		to the list, one at a time or in bulk, and that the size of the list
		reflects the number of elements added.`,
		func(t *testing.T) {
			sl := ds.NewSingleLinkedList()
			assert.Nil(t, sl.Add("test"))
			assert.False(t, sl.IsEmpty())
			assert.Equal(t, 1, sl.Size())
			assert.Nil(t, sl.Add("2"))
			assert.Equal(t, 2, sl.Size())
		})
	t.Run(`Test removing elements from the list: Confirm that elements can be
		removed from the list, one at a time or in bulk, and that the size of
		the list reflects the number of elements removed.`,
		func(t *testing.T) {
			sl := ds.NewSingleLinkedList()
			assert.Nil(t, sl.Add("test"))
			assert.Equal(t, 1, sl.Size())
			assert.Nil(t, sl.Add("2"))
			assert.Equal(t, 2, sl.Size())
			assert.Nil(t, sl.Remove(0))
			assert.Equal(t, 1, sl.Size())
		})
	t.Run(`Test retrieving elements from the list: Test that elements can be retrieved
		from the list by index or by value, and that the correct element is returned.`,
		func(t *testing.T) {
			sl := ds.NewSingleLinkedList()
			assert.Nil(t, sl.Add("test"))
			assert.Nil(t, sl.Add("1"))
			assert.Equal(t, 2, sl.Size())
			assert.Equal(t, "test", sl.Search(0))
			assert.Equal(t, "1", sl.Search(1))
			assert.Nil(t, sl.Add("2"))
			assert.Nil(t, sl.Add("3"))
			assert.Equal(t, 4, sl.Size())
			assert.Equal(t, "3", sl.Search(3))
		})
	t.Run(`Test inserting elements at a specific index: Test that elements can be
		inserted at a specific index in the list, and that the size of the list
		reflects the addition of the new element.`,
		func(t *testing.T) {
			sl := ds.NewSingleLinkedList()
			assert.Nil(t, sl.Add("0"))
			assert.Nil(t, sl.Add("1"))
			assert.Equal(t, "1", sl.Search(1))
			assert.Equal(t, 2, sl.Size())
			assert.Nil(t, sl.Insert(1, "pushed"))
			assert.Equal(t, 3, sl.Size())
			assert.Equal(t, "pushed", sl.Search(1))
			assert.Equal(t, "0", sl.Search(0))
			assert.Equal(t, "1", sl.Search(2))
		})
	t.Run(`Test modifying elements in the list: Verify that elements in the
		list can be modified, and that the modified element is correctly reflected in the list.`,
		func(t *testing.T) {
			sl := ds.NewSingleLinkedList()
			assert.Nil(t, sl.Add("test"))
			assert.Nil(t, sl.Add("1"))
			assert.Nil(t, sl.Insert(1, "pushed"))
			// values should be [test, pushed, 1]
			assert.Equal(t, 3, sl.Size())
			assert.Equal(t, "pushed", sl.Search(1))
			assert.Nil(t, sl.Update(1, "updated1"))
			assert.Equal(t, "updated1", sl.Search(1))
			assert.Nil(t, sl.Update(0, "updated0"))
			assert.Equal(t, "updated0", sl.Search(0))
			assert.Nil(t, sl.Update(2, "updated2"))
			assert.Equal(t, "updated2", sl.Search(2))
		})
	t.Run(`Test clearing the list: Confirm that the list can be cleared of all elements,
		and that the size of the list is zero after clearing.`,
		func(t *testing.T) {
			sl := ds.NewSingleLinkedList()
			assert.Nil(t, sl.Add("test"))
			assert.Nil(t, sl.Add("1"))
			assert.Nil(t, sl.Insert(1, "pushed"))
			// values should be [test, pushed, 1]
			assert.Equal(t, 3, sl.Size())
			assert.Nil(t, sl.Clear())
			assert.Equal(t, 0, sl.Size())
		})
	t.Run(`Test handling edge cases: Test the behavior of the list when it is empty, has
		only one element, or has multiple elements. Test the behavior of the list when attempting
		to add, remove, or retrieve elements that are out of bounds or not present in the list.`,
		func(t *testing.T) {
			sl := ds.NewSingleLinkedList()
			assert.Nil(t, sl.Add("test"))
			assert.Nil(t, sl.Add("1"))
			assert.Nil(t, sl.Insert(1, "pushed"))
			// values should be [test, pushed, 1]

			assert.Equal(t, 3, sl.Size())
			assert.Error(t, sl.Insert(500, "pushed"))
			assert.Equal(t, 3, sl.Size()) // stay the same since the above line failed inserting
			assert.Error(t, sl.Remove(500))
			assert.Equal(t, 3, sl.Size()) // stay the same since the above line failed removing
			assert.Error(t, sl.Update(500, "update"))
			assert.Error(t, sl.Search(500).(error))
		})
}
