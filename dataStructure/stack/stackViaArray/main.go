package main

import (
	"fmt"
	"github.com/Ferrymania/GoForInterview/dataStructure/stack/stackViaArray/arraystack"
)

func main() {
	fmt.Println("creat a stack by array")
	s := arraystack.Constructor(0)
	for i:=0;i<6;i++{
		s.Push(i*2)
	}
	fmt.Println("stack size is ",s.GetSize())
	fmt.Println(s)
	fmt.Println("stack top element is ",s.Peek())
	s.Pop()
	fmt.Println("after pop",s)
}