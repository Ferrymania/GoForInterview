package main

import (
	"fmt"
	"github.com/Ferrymania/GoForInterview/dataStructure/stack/stackViaLinkedlist/linkedstack"
)

func main() {
	fmt.Println("create stack by linkedList")
	ls := linkedstack.NewLinkedStack()
	ls.Push(1)
	ls.Push(3)
	ls.Push(5)
	fmt.Println(ls)
	fmt.Println("stack top is ",ls.Peek())
	e := ls.Pop()
	fmt.Println(" pop top element",e)
	fmt.Println("stack size is ",ls.Size())

}
