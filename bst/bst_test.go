package bst

import (
	"math/rand"
	"testing"
)

func makeTestTree() *tree {
	tree := NewTree(10)
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(3)
	tree.Insert(7)
	tree.Insert(12)
	tree.Insert(18)
	return tree
}

func TestNewTree(t *testing.T) {
	tree := NewTree(8)
	if tree == nil {
		t.Fatal("expected a non-nil tree")
	}
	if tree.Length() != 1 {
		t.Errorf("expected length 1, got %d", tree.Length())
	}
	if tree.root == nil || tree.root.val != 8 {
		t.Errorf("expected root value 8, got %d", tree.root.val)
	}
}

func TestInsert(t *testing.T) {
	tree := NewTree(10)

	tree.Insert(5)
	tree.Insert(15)

	if tree.Length() != 3 {
		t.Errorf("expected length 3, got %d", tree.Length())
	}

	if !tree.Search(5) || !tree.Search(10) || !tree.Search(15) {
		t.Errorf("tree should contain 5, 10, 15")
	}

	if tree.Search(20) {
		t.Errorf("tree should not contain 20")
	}
}

func TestSearch(t *testing.T) {
	tree := makeTestTree()

	cases := []struct {
		val      int
		expected bool
	}{
		{10, true},
		{3, true},
		{7, true},
		{18, true},
		{100, false},
		{-1, false},
	}

	for _, c := range cases {
		if result := tree.Search(c.val); result != c.expected {
			t.Errorf("Search(%d) = %t; want %t", c.val, result, c.expected)
		}
	}
}

func TestDelete(t *testing.T) {

	// 1. Deleting a leaf
	// 2. Deleting a node with one child
	// 3. Deleting a node with two children
	// 4. Deleting the root
	// 5. Deleting a non-existent node

}

func buildTestData(size int) (*tree, []int) {

	vals := make([]int, 0, size)
	tree := NewTree(rand.Intn(size * 10))
	vals = append(vals, tree.root.val)

	for i := 1; i < size; i++ {
		v := rand.Intn(size * 10)
		tree.Insert(v)
		vals = append(vals, v)
	}

	return tree, vals
}

func linearSearch(vals []int, target int) bool {
	for _, v := range vals {
		if v == target {
			return true
		}
	}
	return false
}

func BenchmarkSearchBST(b *testing.B) {
	tree, _ := buildTestData(100000)

	target := 99999
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		tree.Search(target)
	}
}

func BenchmarkSearchLinear(b *testing.B) {
	_, vals := buildTestData(100000)

	target := 99999
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		linearSearch(vals, target)
	}
}
