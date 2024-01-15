package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoublyLinkedLists(t *testing.T) {

	t.Run("should return a pointer to doubleLinkedList", func(t *testing.T) {
		got := InitDoubleList()

		want := &DoubleList{}

		assert.Equal(t, got, want)
	})

	t.Run("should set tail and head to same value if first insertion", func(t *testing.T) {
		list := InitDoubleList()
		node := &doubleNode{value: 1}

		list.firstInsertion(node)

		assert.Equal(t, list.head, list.tail)

	})

}

func BenchmarkDoublyLinkedList(b *testing.B) {

	b.Run("insert first", func(b *testing.B) {
		list := InitDoubleList()

		for i := 0; i < b.N; i++ {
			list.InsertFirst(i)
		}
	})

	list := InitDoubleList()
	b.Run("insert last", func(b *testing.B) {

		for i := 0; i < b.N; i++ {
			list.InsertLast(i)
		}
	})

	b.Run("insert At", func(b *testing.B) {
		last := list.tail.value

		for i := 0; i < b.N; i++ {
			list.InsertAt(last, i)
		}
	})

	b.Run("delete", func(b *testing.B) {

		for i := 0; i < b.N; i++ {
			list.Delete(i)
		}
	})

}
