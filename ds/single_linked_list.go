package ds

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
	for i := 0; i < sl.Size(); i++ {
		if i+1 == index {
			if currentNode.next != nil {
				currentNode.next = currentNode.next.next
			}
			break
		}
		currentNode = currentNode.next
	}
	sl.length--
	return nil
}
