/*
Serialization is the process of converting a data structure or object into a sequence of bits so that it can be stored in a file or memory buffer, or transmitted across a network connection link to be reconstructed later in the same or another computer environment.

Design an algorithm to serialize and deserialize a binary tree. There is no restriction on how your serialization/deserialization algorithm should work. You just need to ensure that a binary tree can be serialized to a string and this string can be deserialized to the original tree structure.

You may serialize the following tree:

    1
   / \
  2   3
     / \
    4   5

as "[1,2,3,null,null,4,5]"
 */
package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct{
	data int
	leftChild *TreeNode
	rightChild *TreeNode
}

//采用广度遍历
func serialize(root *TreeNode)string{
	var s string
	if root == nil {
		return s
	}
	queue := &LinkedQueue{}
	queue.EnQueue(root)
	s += strconv.Itoa(root.data) + " "
	for !queue.IsEmpty() {
		tempNode := queue.DeQueue()
		if tempNode.leftChild != nil {
			queue.EnQueue(tempNode.leftChild)
			s +=strconv.Itoa(tempNode.leftChild.data) + " "
		}else{
			s += "# "
		}
		if tempNode.rightChild != nil {
			queue.EnQueue(tempNode.rightChild)
			s +=strconv.Itoa(tempNode.rightChild.data) + " "
		}else{
			s += "# "
		}


	}
	return s
}

func deserialize(s string)*TreeNode {
	if len(s) == 0 {
		return nil
	}
	serNode := strings.Split(s," ")
	fmt.Println(serNode)
	queue := &LinkedQueue{}
	value,_ := strconv.Atoi(serNode[0])
	root := &TreeNode{data:value}
	queue.EnQueue(root)
	i := 1
	var v int
	for i<len(serNode) {
		node := queue.DeQueue()
		if serNode[i] == "#" {
			node.leftChild = nil
		}else {
			v,_ = strconv.Atoi(serNode[i])
			node.leftChild = &TreeNode{data:v}
		}
		i++
		if serNode[i] == "#" {
			node.rightChild = nil
		}else {
			v,_ = strconv.Atoi(serNode[i])
			node.rightChild = &TreeNode{data:v}
		}
		i++
		if node.leftChild != nil {
			queue.EnQueue(node.leftChild)
		}
		if node.rightChild != nil {
			queue.EnQueue(node.rightChild)
		}
	}
	return root
}
func main(){
	root := &TreeNode{data:1}
	root.leftChild = &TreeNode{data:2}
	root.rightChild = &TreeNode{data:3}
	root.rightChild.leftChild = &TreeNode{data:4}
	root.rightChild.rightChild = &TreeNode{data:5}
	s := serialize(root)
	fmt.Println(s)
	r := deserialize(s)
	fmt.Println(r.data)
}

//定义一个队列
type node struct {
	value *TreeNode
	next *node
}

type LinkedQueue struct {
	head  *node
	end  *node
}

func(q *LinkedQueue) IsEmpty() bool{
	return q.head == nil
}

func(q *LinkedQueue) GetSize() int {
	if q.IsEmpty() {
		return 0
	}
	size :=0
	for cur:=q.head;cur !=nil;cur = cur.next{
		size++
	}
	return size
}

func(q *LinkedQueue)EnQueue(value *TreeNode){
	e := &node{value:value}
	if q.IsEmpty() {
		q.head = e
		q.end = e
	}else {
		q.end.next = e
		q.end = e
	}
}

func(q *LinkedQueue)DeQueue()*TreeNode{
	if q.head == nil {
		panic("queue is already empty")
	}
	node := q.head.value
	q.head = q.head.next
	if q.head == nil {
		q.end = nil
	}
	return node
}