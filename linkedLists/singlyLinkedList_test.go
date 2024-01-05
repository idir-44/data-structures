package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedLists(t *testing.T) {

	t.Run("should return a pointer to list", func(t *testing.T) {
		got := Init()

		want := &List{}

		assert.Equal(t, got, want)
	})

	t.Run("should set tail and head to same value if first insertion", func(t *testing.T) {
		list := Init()
		node := &node{}

		list.firstInsertion(node, 1)

		assert.Equal(t, list.head, list.tail)

	})

}
