package arraystack

import (
	"bytes"
	"fmt"
)

type SliceStack struct {
	arr []interface{}
	stackSize int
}

func Constructor(capacity int) * SliceStack {
	return &SliceStack{
		arr:make([]interface{},0,capacity),
	}
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

func(s *SliceStack) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("Stack: ")
	buffer.WriteString("[")
	for i:=0;i<s.GetSize();i++ {
		buffer.WriteString(fmt.Sprint(s.arr[i]))
		if i != s.GetSize()-1 {
			buffer.WriteString(", ")
		}
	}

	buffer.WriteString("] top")
	return buffer.String()
}
