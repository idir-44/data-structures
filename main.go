package main

import (
	"math/rand"

	"github.com/idir-44/dataStructures/bst"
)

func main() {
	t := bst.NewTree(45)

	for i := 0; i < 16; i++ {
		t.Insert(rand.Intn(100))
	}

}
