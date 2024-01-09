package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: search how to properly test linked lists

func TestLinkedLists(t *testing.T) {

	t.Run("should return a pointer to list", func(t *testing.T) {
		got := InitList()

		want := &List{}

		assert.Equal(t, got, want)
	})

	t.Run("should set tail and head to same value if first insertion", func(t *testing.T) {
		list := InitList()
		node := &node{}

		list.firstInsertion(node, 1)

		assert.Equal(t, list.head, list.tail)

	})

}
