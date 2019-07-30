/*
This is a package about SinglyLinkedlist
*/
package main

import (
	"fmt"
)

type Node struct { //define each node strcut
	value interface{}
	next  *Node
}

func (n *Node) String() string { //print node
	return fmt.Sprint(n.value)
}

type LinkedList struct { // define  linkedList struct ,head point to first node
	head *Node
	size int
}

func CreateLinkList() *LinkedList { // init one LinkedList
	return &LinkedList{
		head: &Node{},
	}
}

func (ll *LinkedList) GetSize() int {
	return ll.size
}

func (ll *LinkedList) IsEmpty() bool {
	return ll.size == 0
}

func (ll *LinkedList) Add(index int, value interface{}) { //Add one node in the index position of LinkedList
	if index < 0 || index > ll.size {
		panic("Illegal index")
	}

	prev := ll.head
	for i := 0; i < index; i++ {
		prev = prev.next
	}

	prev.next = &Node{value, prev.next}
	ll.size++

}

func (ll *LinkedList) AddToHead(value interface{}) { //Add node in the head of linkedList
	ll.Add(0, value)
}

func (ll *LinkedList) AddToTail(value interface{}) { //Add node in the tail of linkedList
	ll.Add(ll.size, value)
}

func (ll *LinkedList) Get(index int) interface{} { //query node of index position
	if index < 0 || index >= ll.size {
		panic("Illegal index")
	}

	currentNode := ll.head.next
	for i := 0; i < index; i++ {
		currentNode = currentNode.next
	}

	return currentNode.value
}

func (ll *LinkedList) GetHead() interface{} { //query head node
	return ll.Get(0)
}

func (ll *LinkedList) GetTail() interface{} { //query tail node
	return ll.Get(ll.size - 1)
}

func (ll *LinkedList) Set(index int, v interface{}) { //set index position node with a new value
	if index < 0 || index >= ll.size {
		panic("Illegal index")
	}

	currentNode := ll.head.next
	for i := 0; i < index; i++ {
		currentNode = currentNode.next
	}
	currentNode.value = v
}

func (ll *LinkedList) Contains(v interface{}) bool { //decide whether the linkedList contains the value
	currentNode := ll.head.next

	for currentNode != nil {
		if currentNode.value == v {
			return true
		}
		currentNode = currentNode.next
	}

	return false
}

func (ll *LinkedList) Delete(index int) interface{} {
	if index < 0 || index >= ll.size {
		panic("Delete failed.Index is illeagal")
	}

	prev := ll.head
	for i := 0; i < index; i++ {
		prev = prev.next
	}

	retNode := prev.next
	prev.next = retNode.next
	retNode.next = nil
	ll.size--
	return retNode.value
}

func (ll *LinkedList) DeleteFirst() {
	ll.Delete(0)
}

func (ll *LinkedList) DeleteLast() {
	ll.Delete(ll.size - 1)
}

func (ll *LinkedList) PrintLinkedList() {
	currentNode := ll.head
	if currentNode == nil {
		fmt.Println("LinkedList is empty")
	}
	fmt.Printf("%+v\n", *currentNode)

	for currentNode.next != nil {
		currentNode = currentNode.next
		fmt.Printf("%+v\n", *currentNode)
	}
}
func main() {
	linkedList := CreateLinkList()

	for i := 0; i < 5; i++ {
		linkedList.AddToHead(i)
		fmt.Println(linkedList)
	}

	linkedList.Add(2, 666)
	fmt.Println(linkedList)
	linkedList.PrintLinkedList()
	fmt.Println("Does linkedlist contanis 666", linkedList.Contains(666))
	linkedList.Delete(4)
	linkedList.PrintLinkedList()
}
