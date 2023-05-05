package ds_test

import (
	"testing"

	"github.com/nmmugia/dsa/ds"
	"github.com/stretchr/testify/assert"
)

/*
test requirements for a Binary Tree data structure:

Test that a new empty tree is created correctly:
Create a new tree and check that its root node is nil.
Check that its size is 0.


Test that nodes can be inserted into the tree correctly:
Insert a single node into an empty tree and check that its root is set correctly and its size is 1.
Insert multiple nodes into a tree and check that their order is preserved and their parent/child relationships are set correctly.
Insert nodes with both left and right children and check that their relationships are set correctly.
Insert nodes that already has the same value, then it should return error

Test that nodes can be removed from the tree correctly.:
Remove the only node in a tree and check that its root is nil and its size is 0
Remove nodes from the left and right of the tree and check that their parent/child relationships are updated correctly.
Remove nodes with both left and right children and check that their relationships are updated correctly.


Test that nodes can be searched for correctly:
Search for the root node and check that its value is correct.
Search for nodes in the left and right subtrees and check that their values are correct.
Search for a node that is not in the tree and check that nil is returned.


Test that the size of the tree is updated correctly:
Insert and remove nodes from the tree and check that its size is updated correctly each time.


Test that the tree can be traversed correctly:
Traverse an empty tree and check that no nodes are returned.
Traverse a non-empty tree in preorder, inorder, and postorder and check that all nodes are returned in the correct order.
Traverse a non-empty tree in level order and check that all nodes are returned in the correct order.


Test that the height of the tree is calculated correctly:
Calculate the height of an empty tree and check that it is 0.
Calculate the height of a non-empty tree and check that it is correct.


Test error handling for invalid operations:
Attempt to remove a node from an empty tree and check that an error is returned.
Attempt to remove a node that is not in the tree and check that an error is returned.
Attempt to search for a node with an invalid value and check that nil is returned.
*/

func TestBTreeCreation(t *testing.T) {
	t.Run(`Create a new tree and check that its root node is nil.`,
		func(t *testing.T) {
			bt := ds.NewBTree()
			assert.True(t, bt.IsEmpty())
		})
	t.Run(`Check that its size is 0.`,
		func(t *testing.T) {
			bt := ds.NewBTree()
			assert.Equal(t, 0, bt.Size())
		})
}

func TestBTreeInsertion(t *testing.T) {
	t.Run(`Insert a single node into an empty tree and check that
		its root is set correctly and its size is 1.`,
		func(t *testing.T) {
			bt := ds.NewBTree()
			assert.Nil(t, bt.Insert(10, "test"))
			assert.Equal(t, 1, bt.Size())
			assert.NotNil(t, bt.GetRoot())
			assert.Equal(t, "test", bt.GetRoot().GetValue())
			assert.Equal(t, 10, bt.GetRoot().GetKey())
		})
	t.Run(`Insert multiple nodes into a tree and check that their order is
	preserved and their parent/child relationships are set correctly.`,
		func(t *testing.T) {
			bt := ds.NewBTree()
			assert.Nil(t, bt.Insert(10, "ten "))
			assert.Nil(t, bt.Insert(5, "sgo"))
			assert.Nil(t, bt.Insert(2, "s"))
			assert.Equal(t, 3, bt.Size())
		})
	t.Run(`Insert nodes with both left and right children and check
		that their relationships are set correctly.`,
		func(t *testing.T) {
			bt := ds.NewBTree()
			assert.Nil(t, bt.Insert(10, "aii"))
			assert.Nil(t, bt.Insert(5, "aii"))
			assert.Nil(t, bt.Insert(13, "aii"))
			assert.Nil(t, bt.Insert(12, "aii"))
			assert.Nil(t, bt.Insert(2, "aii"))
			assert.Equal(t, 5, bt.Size())
		})
	t.Run(`Insert nodes that already has the same key, then it should return error `,
		func(t *testing.T) {
			bt := ds.NewBTree()
			assert.Nil(t, bt.Insert(10, "s"))
			assert.Error(t, bt.Insert(10, "ss"))
			assert.Equal(t, 1, bt.Size())
		})
}

func TestRemoval(t *testing.T) {
	t.Run(`Remove the only node in a tree and check that its root is nil and its size is 0`,
		func(t *testing.T) {
			bt := ds.NewBTree()
			assert.Nil(t, bt.Insert(10, "ss"))
			assert.Equal(t, 1, bt.Size())
			assert.Nil(t, bt.Remove(10))
			assert.Nil(t, bt.GetRoot())
			assert.Equal(t, 0, bt.Size())
		})
	t.Run(`Remove nodes from the left and right of the tree and check that their parent/child relationships are updated correctly.`,
		func(t *testing.T) {
			bt := ds.NewBTree()
			assert.Nil(t, bt.Insert(10, "10"))
			assert.Nil(t, bt.Insert(5, "5"))
			assert.Nil(t, bt.Insert(15, "15"))
			assert.Equal(t, 3, bt.Size())
			assert.Nil(t, bt.Remove(15))
			assert.Equal(t, 2, bt.Size())
			assert.Nil(t, bt.Remove(5))
			assert.Equal(t, 1, bt.Size())
		})
	t.Run(`Remove nodes with both left and right children and check that their relationships are updated correctly.`,
		func(t *testing.T) {
			bt := bTreeExample()
			assert.Equal(t, 9, bt.Size())
			assert.Nil(t, bt.Remove(14))
			assert.Equal(t, 8, bt.Size())
		})

}

func TestSearch(t *testing.T) {

	t.Run(`Search for the root node and check that its value is correct.`,
		func(t *testing.T) {
			bt := bTreeExample()
			assert.Equal(t, "8", bt.Search(8).GetValue())
		})
	t.Run(`Search for nodes in the left and right subtrees and check that their values are correct.`,
		func(t *testing.T) {
			bt := bTreeExample()
			assert.Equal(t, "3", bt.Search(3).GetValue())
			assert.Equal(t, "10", bt.Search(10).GetValue())
		})
	t.Run(`Search for a node that is not in the tree and check that nil is returned.`,
		func(t *testing.T) {
			bt := bTreeExample()
			assert.Nil(t, bt.Search(18))
		})
}

func TestSize(t *testing.T) {
	t.Run(`Insert and remove nodes from the tree and check that its size is
		updated correctly each time.`,
		func(t *testing.T) {
			bt := bTreeExample()
			assert.Equal(t, 9, bt.Size())
			assert.Nil(t, bt.Insert(50, "aii"))
			assert.Equal(t, 10, bt.Size())
			assert.Nil(t, bt.Remove(50))
			assert.Equal(t, 9, bt.Size())
		})
}

func TestTraversion(t *testing.T) {
	t.Run(`Traverse an empty tree and check that no nodes are returned.`,
		func(t *testing.T) {
			bt := ds.NewBTree()
			assert.Nil(t, bt.Traverse())
		})
	t.Run(`Traverse a non-empty tree in preorder, inorder, and postorder and check that all nodes are returned in the correct order.`,
		func(t *testing.T) {
			bt := bTreeExample()
			result := []int{}
			ds.Preorder(bt.GetRoot(), &result)
			assert.Equal(t, []int{8, 3, 1, 6, 4, 7, 10, 14, 13}, result)
			result = []int{}
			ds.Inorder(bt.GetRoot(), &result)
			assert.Equal(t, []int{1, 3, 4, 6, 7, 8, 10, 13, 14}, result)
			result = []int{}
			ds.Postorder(bt.GetRoot(), &result)
			assert.Equal(t, []int{1, 4, 7, 6, 3, 13, 14, 10, 8}, result)
		})
	t.Run(`Traverse a non-empty tree in level order and check that all nodes are returned in the correct order.`,
		func(t *testing.T) {

		})

}

func bTreeExample() *ds.BTree {
	bt := ds.NewBTree()
	bt.Insert(8, "8")
	bt.Insert(3, "3")
	bt.Insert(1, "1")
	bt.Insert(6, "6")
	bt.Insert(4, "4")
	bt.Insert(7, "7")
	bt.Insert(10, "10")
	bt.Insert(14, "14")
	bt.Insert(13, "13")
	return bt
}
