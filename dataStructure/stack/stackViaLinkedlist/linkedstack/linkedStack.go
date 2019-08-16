package linkedstack

import (
	"bytes"
	"fmt"
	"github.com/Ferrymania/GoForInterview/dataStructure/Linkedlist/singlylinkedlist"
)

type LinkedStack struct {
	head *singlylinkedlist.Node
}


func NewLinkedStack() *LinkedStack{
	return &LinkedStack{head:&singlylinkedlist.Node{}}
}
func(ls *LinkedStack)IsEmpty() bool {
	return ls.head.Next == nil
}

func(ls *LinkedStack) Size() int {
	size := 0
	node := ls.head.Next
	for node !=nil {
		node = node.Next
		size++
	}
	return size
}

func(ls *LinkedStack) Push(e interface{}){
	node := &singlylinkedlist.Node{Value:e,Next:ls.head.Next}
	ls.head.Next = node
}

func(ls *LinkedStack) Pop() interface{} {
	node := ls.head.Next
	if node != nil {
		ls.head.Next = node.Next
		return node.Value
	}
	return nil
}

func(ls *LinkedStack) Peek() interface{} {
	if ls.head.Next != nil {
		return ls.head.Next.Value
	}
	return nil
}

func(ls *LinkedStack ) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("Stack: ")
	buffer.WriteString("[")
	node := ls.head.Next
	for i:=0;i<ls.Size();i++ {
		buffer.WriteString(fmt.Sprint(node.Value))
		node = node.Next
		if i != ls.Size()-1 {
			buffer.WriteString(", ")
		}
	}

	buffer.WriteString("] top")
	return buffer.String()
}