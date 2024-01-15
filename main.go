package main

import (
	"fmt"

	linkedlists "github.com/idir-44/dataStructures/linkedLists"
)

func main() {
	// singlyLinkedListFunctions()
	// doublyLinkedListsFunctions()

}

func doublyLinkedListsFunctions() {
	list := linkedlists.InitDoubleList()

	for i := 0; i < 20; i++ {
		list.InsertLast(i)
	}

	list.InsertAt(10, 69)

	list.InsertAt(69, 666)

	list.Delete(0)

	list.Delete(19)

	list.PrintList()

	fmt.Println("/////// Reverse ////////")

	// list.PrintReverseList()

}

func singlyLinkedListFunctions() {
	list := linkedlists.InitList()

	for i := 0; i < 20; i++ {
		list.InsertLast(i)
	}

	err := list.InsertAfter(8, 69)
	if err != nil {
		fmt.Println(err)
	}

	err = list.InsertAfter(69, 666)
	if err != nil {
		fmt.Println(err)
	}

	err = list.Delete(666)
	if err != nil {
		fmt.Println(err)
	}

	err = list.Delete(0)
	if err != nil {
		fmt.Println(err)
	}

	err = list.Delete(19)
	if err != nil {
		fmt.Println(err)
	}

	list.PrintList()

	fmt.Printf("my list: %d\n", list.Length())
}
