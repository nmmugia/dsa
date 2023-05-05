package ds

import "fmt"

type BTree struct {
	root *node
	size int
}

type node struct {
	key    int
	value  interface{}
	parent *node
	left   *node
	right  *node
}

func (n *node) GetValue() interface{} {
	if n == nil {
		return nil
	}
	return n.value
}
func (n *node) GetKey() interface{} {
	if n == nil {
		return nil
	}
	return n.key
}
func (n *node) GetParent() *node {
	if n == nil {
		return nil
	}
	return n.parent
}
func (n *node) GetLeft() *node {
	if n == nil {
		return nil
	}
	return n.left
}
func (n *node) GetRight() *node {
	if n == nil {
		return nil
	}
	return n.right
}

func NewBTree() *BTree {
	return &BTree{
		root: nil,
		size: 0,
	}
}

func (bt *BTree) IsEmpty() bool {
	return bt.size <= 0
}

func (bt *BTree) Size() int {
	return bt.size
}

func (bt *BTree) Insert(key int, value interface{}) error {
	if bt.Size() <= 0 {
		bt.root = &node{
			key:   key,
			value: value,
		}
		bt.size++
		return nil
	}
	currentNode := bt.root
	for currentNode != nil {
		if key < currentNode.key {
			if currentNode.left == nil {
				currentNode.left = &node{
					key:    key,
					value:  value,
					parent: currentNode,
				}
				bt.size++
				return nil
			}
			currentNode = currentNode.left
		} else if key > currentNode.key {
			if currentNode.right == nil {
				currentNode.right = &node{
					key:    key,
					value:  value,
					parent: currentNode,
				}
				bt.size++
				return nil
			}
			currentNode = currentNode.right
		} else {
			return fmt.Errorf("error value already exists inserting value %d", value)
		}
	}
	return fmt.Errorf("error inserting value %d", value)
}

func (bt *BTree) GetRoot() *node {
	return bt.root
}

func (bt *BTree) Remove(index int) error {
	currentNode := bt.Search(index)

	// if the deleted node is leaf
	if currentNode.left == nil && currentNode.right == nil {
		// if current node is root
		if currentNode.parent == nil {
			bt.root = nil
		} else if currentNode == currentNode.parent.left {
			currentNode.parent.left = nil
		} else {
			currentNode.parent.right = nil
		}
	} else if currentNode.left == nil {
		if currentNode.parent == nil {
			bt.root = currentNode.right
		} else if currentNode == currentNode.parent.left {
			currentNode.parent.left = currentNode.right
		} else {
			currentNode.parent.right = currentNode.right
		}
	} else if currentNode.right == nil {
		if currentNode.parent == nil {
			bt.root = currentNode.right
		} else if currentNode == currentNode.parent.left {
			currentNode.parent.left = currentNode.left
		} else {
			currentNode.parent.right = currentNode.left
		}
	} else {
		successor := currentNode.GetRight()
		for successor.left != nil {
			successor = successor.left
		}
		bt.Remove(successor.key)
		successor.left = currentNode.left
		successor.right = currentNode.right
		if currentNode.parent == nil {
			bt.root = successor
		} else if currentNode.parent.right == currentNode {
			currentNode.right = successor
		} else {
			currentNode.left = successor
		}
		return nil
	}
	bt.size--
	return nil
}

func (bt *BTree) Search(index int) *node {
	currentNode := bt.root
	for currentNode != nil {
		if index < currentNode.key {
			currentNode = currentNode.left
		} else if index > currentNode.key {
			currentNode = currentNode.right
		} else {
			return currentNode
		}
	}
	return nil
}

func Preorder(n *node, result *[]int) {
	if n == nil {
		return
	}
	*result = append(*result, n.key)
	Preorder(n.left, result)
	Preorder(n.right, result)
}

func Inorder(n *node, result *[]int) {
	if n == nil {
		return
	}
	Inorder(n.left, result)
	*result = append(*result, n.key)
	Inorder(n.right, result)
}

func Postorder(n *node, result *[]int) {
	if n == nil {
		return
	}
	Postorder(n.left, result)
	Postorder(n.right, result)
	*result = append(*result, n.key)
}

func (bt *BTree) Traverse() (result []int) {
	if bt.IsEmpty() {
		return nil
	}

	// Preorder traversal
	// fmt.Println(Preorder(bt.root))
	Preorder(bt.root, &result)

	// Inorder traversal
	Inorder(bt.root, &result)

	// Postorder traversal
	Postorder(bt.root, &result)

	// Level order traversal
	levelorder := func(n *node) {
		if n == nil {
			return
		}
		var queue []*node
		queue = append(queue, n)
		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]
			result = append(result, node.key)
			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.right != nil {
				queue = append(queue, node.right)
			}
		}
	}
	levelorder(bt.root)

	return result
}
