package main

import (
	"fmt"
)
type SliceStack struct {
	arr [](interface{})
	stackSize int


}

func(s *SliceStack) IsEmpty() bool {
	return s.stackSize == 0
}

func(s *SliceStack) GetSize() int {
	return s.stackSize
}

func(s *SliceStack) Peek() interface{} {
	if s.IsEmpty() {
		panic("stack is empty")
	}

	return s.arr[s.stackSize-1]
}

func(s *SliceStack) Pop() interface{} {
	if s.IsEmpty() {
		panic("stack is empty ,no element pop")
	}
	s.stackSize--
	ret := s.arr[s.stackSize]
	s.arr = s.arr[:s.stackSize]
	return ret
}

func(s *SliceStack)Push(t interface{}){
	s.arr=append(s.arr,t)
	s.stackSize++	
}

func main() {
	fmt.Println("creat a stack by array")
	s := &SliceStack{
		arr:make([]int,0),
	}
}
