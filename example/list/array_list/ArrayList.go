package main

import (
	"fmt"

	"github.com/mehditeymorian/structures/list"
)

func main() {
	al := list.NewArrayList[int]()
	// al.Add(1, 2, 3, 4)
	// fmt.Println(al.Size())
	// al.Add(456)
	// fmt.Println(al.Size())
	// fmt.Println(al.Empty())
	// fmt.Println(al.GetAll())

	al.Add(1, 2, 3)
	fmt.Println(al.GetAll())
	al.Remove(1)
	fmt.Println(al.GetAll())

	doubledSum := []interface{}{al.Map(func(item int) int {
		return item * item
	}).Reduce(0, func(prev, current int) int {
		return prev + current
	})}

	fmt.Println(doubledSum)

}
