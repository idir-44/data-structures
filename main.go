package main

import (
	"fmt"

	linkedlists "github.com/idir-44/dataStructures/linkedLists"
)

func main() {
	list := linkedlists.Init()

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
