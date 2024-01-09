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
