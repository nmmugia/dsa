package ds

import (
	"fmt"
)

type Node struct {
	value interface{}
	next  *Node
}

type SingleLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func NewSingleLinkedList() *SingleLinkedList {
	return &SingleLinkedList{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

func (sl *SingleLinkedList) IsEmpty() bool {
	return sl.length <= 0
}

func (sl *SingleLinkedList) Size() int {
	return sl.length
}

func (sl *SingleLinkedList) Add(value interface{}) error {
	currentNode := &Node{
		value: value,
		next:  nil,
	}
	if sl.head == nil {
		sl.head = currentNode
	}
	if sl.tail != nil {
		sl.tail.next = currentNode
	}
	sl.tail = currentNode
	sl.length++
	return nil
}

func (sl *SingleLinkedList) Remove(index int) error {
	if index == 0 {
		if sl.head.next != nil {
			sl.head = sl.head.next
		} else {
			sl.head = nil
		}
		sl.length--
		return nil
	}
	currentNode := sl.head
	for i := 0; i < sl.Size()-1; i++ {
		if i+1 == index && currentNode.next != nil {
			currentNode.next = currentNode.next.next
			sl.length--
			return nil
		}
		currentNode = currentNode.next
	}
	return fmt.Errorf("error failed removing node, node isn't founded, index:%d", index)
}

func (sl *SingleLinkedList) Search(index int) interface{} {
	if index == 0 {
		return sl.head.value
	}
	currentNode := sl.head
	for i := 0; i < sl.Size()-1; i++ {
		if i+1 == index && currentNode.next != nil {
			return currentNode.next.value
		}
		currentNode = currentNode.next
	}
	return fmt.Errorf("node is not found, index: %d", index)
}

func (sl *SingleLinkedList) Insert(index int, value interface{}) error {
	if index == sl.Size() || (index == 0 && sl.Size() == 0) {
		return sl.Add(value)
	}
	currentNode := sl.head
	for i := 0; i < sl.Size()-1; i++ {
		if i+1 == index && currentNode.next != nil {
			insertedNode := &Node{
				value: value,
				next:  currentNode.next,
			}
			currentNode.next = insertedNode
			sl.length++
			return nil
		}
		currentNode = currentNode.next
	}
	return fmt.Errorf("error failed inserting node, index unbound, index:%d", index)
}

func (sl *SingleLinkedList) Update(index int, value interface{}) error {
	if index < 0 || index >= sl.Size() {
		return fmt.Errorf("error failed update node, index unbound, index:%d", index)
	}
	if index == 0 {
		sl.head.value = value
	}
	currentNode := sl.head
	for i := 0; i < sl.Size()-1; i++ {
		if i+1 == index && currentNode.next != nil {
			currentNode.next.value = value
			break
		}
		currentNode = currentNode.next
	}
	return nil
}

func (sl *SingleLinkedList) Clear() error {
	sl.head = nil
	sl.tail = nil
	sl.length = 0
	return nil
}
