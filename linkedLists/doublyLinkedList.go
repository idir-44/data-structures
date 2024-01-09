package linkedlists

import "fmt"

type doubleNode struct {
	value int
	next  *doubleNode
	prev  *doubleNode
}

type DoubleList struct {
	head   *doubleNode
	tail   *doubleNode
	length int
}

func InitDoubleList() *DoubleList {
	return &DoubleList{}
}

func (l *DoubleList) Length() int {
	return l.length
}

func (l *DoubleList) firstInsertion(node *doubleNode) {

	l.head = node
	l.tail = node
	l.length++

	return
}

func (l *DoubleList) InsertFirst(value int) {
	n := &doubleNode{value: value}

	if l.head == nil {
		l.firstInsertion(n)
		return
	}

	n.next = l.head
	n.next.prev = n
	l.head = n
	l.length++

}

func (l *DoubleList) InsertLast(value int) {
	n := &doubleNode{value: value}

	if l.head == nil {
		l.firstInsertion(n)
		return
	}

	n.prev = l.tail
	l.tail.next = n
	l.tail = n
	l.length++

}

func (l *DoubleList) InsertAt(index, value int) error {
	n := &doubleNode{value: value}

	if l.head == nil {
		l.firstInsertion(n)
		return nil
	}

	if index == l.head.value {
		l.InsertFirst(value)
		n = nil
		return nil
	}

	i := l.head
	for i != nil {
		if i.value == index {
			if i == l.tail {
				l.InsertLast(value)
				n = nil
				return nil
			}

			n.next = i.next
			n.prev = i
			i.next.prev = n
			i.next = n
			l.length++

			return nil
		}

		i = i.next
	}

	return fmt.Errorf("value not found")
}

func (l *DoubleList) PrintList() {
	if l.head != nil {
		node := l.head
		for node != nil {
			fmt.Printf("node: %d\n", node.value)
			node = node.next
		}
	} else {
		fmt.Println("empty list")
	}
}

func (l *DoubleList) PrintReverseList() {
	if l.tail != nil {
		node := l.tail
		for node != nil {
			fmt.Printf("node: %d\n", node.value)
			node = node.prev
		}
	} else {
		fmt.Println("empty list")
	}
}
