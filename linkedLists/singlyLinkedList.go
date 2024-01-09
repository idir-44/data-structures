package linkedlists

import (
	"fmt"
)

type node struct {
	value int
	next  *node
}

type List struct {
	length int
	head   *node
	tail   *node
}

func InitList() *List {
	return &List{}
}

func (l *List) Length() int {
	return l.length
}

func (l *List) firstInsertion(n *node, value int) {
	n.value = value
	l.head = n
	l.tail = n
	l.length++

	return
}

func (l *List) InsertFirst(value int) {
	n := &node{}

	if l.head == nil && l.tail == nil {
		l.firstInsertion(n, value)

		return
	}

	n.value = value
	n.next = l.head
	l.head = n
	l.length++
}

func (l *List) InsertLast(value int) {
	n := &node{}

	if l.head == nil && l.tail == nil {
		l.firstInsertion(n, value)

		return
	}

	n.value = value
	l.tail.next = n
	l.tail = n

	l.length++
}

func (l *List) InsertAfter(index, value int) error {
	n := &node{}
	if l.head == nil && l.tail == nil {
		l.firstInsertion(n, value)

		return nil
	}

	i := l.head

	for i != nil {
		if i.value == index {
			n.value = value
			n.next = i.next
			i.next = n
			l.length++

			return nil
		}
		i = i.next
	}

	return fmt.Errorf("value not found")
}

func (l *List) Delete(value int) error {
	if l.head == nil {
		return fmt.Errorf("empty List")
	}

	n := l.head

	if l.head.value == value {
		l.head = n.next
		n = nil
		l.length--

		return nil
	}

	for n.next != nil {
		if n.next.value == value {
			if n.next == l.tail {
				l.tail = n
				n.next = nil
				l.length--

				return nil
			}

			n.next = n.next.next
			l.length--

			return nil
		}
		n = n.next
	}

	return fmt.Errorf("value not found")
}

func (l *List) PrintList() {
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
