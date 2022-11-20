package main

import (
	"fmt"

	"github.com/meauxses/DataStructuresGo/LinkedList"
)

func main() {
	//Testing for ints
	list := LinkedList.NewList[int]()

	list.Append(1)
	list.Prepend(2)
	list.Append(3)
	list.Prepend(4)
	list.Print()
	fmt.Println("Size:", list.Size())


	list.Remove(4)
	list.Print()
	fmt.Println("Size:", list.Size())
	list.Remove(3)
	list.Print()
	fmt.Println("Size:", list.Size())
	list.Remove(4)
	fmt.Println(list.Pop())
	fmt.Println(list.Pop())
	fmt.Println(list.Pop())
	list.Append(1)
	list.Prepend(2)
	list.Append(3)
	list.Prepend(4)
	fmt.Println("Index 0: ", list.FindAtIndex(0))
	fmt.Println("Index 1: ", list.FindAtIndex(1))
	fmt.Println("Index 2: ", list.FindAtIndex(2))
	fmt.Println("Index 3: ", list.FindAtIndex(3))
	fmt.Println("Index 4: ", list.FindAtIndex(4))

	fmt.Println(list.Dequeue())
	fmt.Println("Size:", list.Size())
	fmt.Println(list.Dequeue())
	fmt.Println("Size:", list.Size())
	fmt.Println(list.Dequeue())
	fmt.Println("Size:", list.Size())

	//testing for strings
	slist := LinkedList.NewList[string]()

	slist.Append("I")
	slist.Prepend("am")
	slist.Append("not")
	slist.Prepend("here")
	slist.Print()
	fmt.Println("Size:", slist.Size())


	slist.Remove("here")
	slist.Print()
	fmt.Println("Size:", slist.Size())
	slist.Remove("not")
	slist.Print()
	fmt.Println("Size:", slist.Size())
	slist.Remove("not")
	fmt.Println(slist.Pop())
	fmt.Println(slist.Pop())
	fmt.Println(slist.Pop())
	slist.Append("I")
	slist.Prepend("You")
	slist.Append("Me")
	slist.Prepend("Us")
	fmt.Println("Index 0: ", slist.FindAtIndex(0))
	fmt.Println("Index 1: ", slist.FindAtIndex(1))
	fmt.Println("Index 2: ", slist.FindAtIndex(2))
	fmt.Println("Index 3: ", slist.FindAtIndex(3))
	fmt.Println("Index 4: ", slist.FindAtIndex(4))

}
