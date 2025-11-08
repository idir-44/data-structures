package bst

import "fmt"

type treeNode struct {
	val   int
	left  *treeNode
	right *treeNode
}

type tree struct {
	length int
	root   *treeNode
}

func NewTree(val int) *tree {
	return &tree{length: 1, root: &treeNode{val: val}}
}

func (t *tree) Length() int {
	return t.length
}

func (t *tree) Search(needle int) bool {
	return search(t.root, needle)
}

func (t *tree) Insert(val int) {
	node := &treeNode{val: val}
	if t.root == nil {
		t.root = node
		t.length++
		return
	}

	if err := insert(t.root, node); err == nil {
		t.length++
	}
}

func (t *tree) Delete(value int) {

}

func (t *tree) Print() {
	inOrderWalk(t.root)
}

func inOrderWalk(node *treeNode) {
	if node.left != nil {
		inOrderWalk(node.left)
	}

	fmt.Println(node.val)

	if node.right != nil {
		inOrderWalk(node.right)
	}
}

func insert(node *treeNode, toInsert *treeNode) error {
	if toInsert.val < node.val {
		if node.left != nil {
			insert(node.left, toInsert)
		} else {
			node.left = toInsert
			return nil
		}
	}

	if toInsert.val > node.val {
		if node.right != nil {
			insert(node.right, toInsert)
		} else {
			node.right = toInsert
			return nil
		}
	}

	return fmt.Errorf("couldn't insert %d", toInsert.val)
}

func search(node *treeNode, needle int) bool {
	if node.val == needle {
		return true
	}

	if needle < node.val && node.left != nil {
		return search(node.left, needle)
	}

	if needle > node.val && node.right != nil {
		return search(node.right, needle)
	}

	return false
}
