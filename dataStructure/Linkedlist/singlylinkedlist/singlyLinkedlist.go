/*
This is a package about SinglyLinkedlist
*/
package singlylinkedlist

import (
	"fmt"
)
//define each node strcut
type Node struct {
	Value interface{}
	Next  *Node
}

func NewNode() *Node {
	return &Node{}
}
//print node
func (n *Node) String() string {
	return fmt.Sprint(n.Value)
}

// define  linkedList struct ,head point to first node
type LinkedList struct {
	Head *Node
	Size int
}

func CreateLinkList() *LinkedList { // init one LinkedList
	return &LinkedList{
		Head: &Node{},
	}
}

func (ll *LinkedList) GetSize() int {
	return ll.Size
}

func (ll *LinkedList) IsEmpty() bool {
	return ll.Size == 0
}

//Add one node in the index position of LinkedList
func (ll *LinkedList) Add(index int, value interface{}) {
	if index < 0 || index > ll.Size {
		panic("Illegal index")
	}

	prev := ll.Head
	for i := 0; i < index; i++ {
		prev = prev.Next
	}

	prev.Next = &Node{value, prev.Next}
	ll.Size++

}

//Add node in the head of linkedList
func (ll *LinkedList) AddToHead(value interface{}) {
	ll.Add(0, value)
}
//Add node in the tail of linkedList
func (ll *LinkedList) AddToTail(value interface{}) {
	ll.Add(ll.Size, value)
}
//query node of index position
func (ll *LinkedList) Get(index int) interface{} {
	if index < 0 || index >= ll.Size {
		panic("Illegal index")
	}

	currentNode := ll.Head.Next
	for i := 0; i < index; i++ {
		currentNode = currentNode.Next
	}

	return currentNode.Value
}

//query head node
func (ll *LinkedList) GetHead() interface{} {
	return ll.Get(0)
}
//query tail node
func (ll *LinkedList) GetTail() interface{} {
	return ll.Get(ll.Size - 1)
}
//set index position node with a new value
func (ll *LinkedList) Set(index int, v interface{}) {
	if index < 0 || index >= ll.Size {
		panic("Illegal index")
	}

	currentNode := ll.Head.Next
	for i := 0; i < index; i++ {
		currentNode = currentNode.Next
	}
	currentNode.Value = v
}
//decide whether the linkedList contains the value
func (ll *LinkedList) Contains(v interface{}) bool {
	currentNode := ll.Head.Next

	for currentNode != nil {
		if currentNode.Value == v {
			return true
		}
		currentNode = currentNode.Next
	}

	return false
}

func (ll *LinkedList) Delete(index int) interface{} {
	if index < 0 || index >= ll.Size {
		panic("Delete failed.Index is illeagal")
	}

	prev := ll.Head
	for i := 0; i < index; i++ {
		prev = prev.Next
	}

	retNode := prev.Next
	prev.Next = retNode.Next
	retNode.Next = nil
	ll.Size--
	return retNode.Value
}

func (ll *LinkedList) DeleteFirst() {
	ll.Delete(0)
}

func (ll *LinkedList) DeleteLast() {
	ll.Delete(ll.Size - 1)
}

func (ll *LinkedList) PrintLinkedList() {
	currentNode := ll.Head
	if currentNode == nil {
		fmt.Println("LinkedList is empty")
	}
	fmt.Printf("%+v\n", *currentNode)

	for currentNode.Next != nil {
		currentNode = currentNode.Next
		fmt.Printf("%+v\n", *currentNode)
	}
}
