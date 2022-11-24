package main

import (
	"fmt"

	"github.com/mehditeymorian/structures/list"
)

func main() {
	ll := list.NewLinkedList[int]()
	ll.Add(1, 2, 3, 4, 5, 6)
	fmt.Println(ll.GetAll())
	ll.Remove(0)
	fmt.Println(ll.GetAll())
	ll.Remove(4)
	fmt.Println(ll.GetAll())
	ll.Remove(2)
	fmt.Println(ll.GetAll())

	item, ok := ll.Get(2)
	fmt.Print(*item, ok)

}
