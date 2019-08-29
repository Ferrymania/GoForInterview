/*
从上往下打印出二叉树的每个结点，同一层的结点按照从左到右的顺序打印
 */
package main

import "fmt"
type TreeNode struct {
	data int
	leftChild *TreeNode
	rightChild *TreeNode
}
/*
使用队列
对于每一个节点，首先访问本节点，再将其子节点放置入先进先出的队列中
*/

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

func(q *LinkedQueue)EnQueue(tNode *TreeNode){
	e := &node{value:tNode}
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
	tNode := q.head.value
	q.head = q.head.next
	if q.head == nil {
		q.end = nil
	}
	return tNode
}

func PrintLevelOrder(root *TreeNode){
	if root == nil {
		return
	}
	queue := &LinkedQueue{}
	queue.EnQueue(root)
	for !queue.IsEmpty() {
		tempNode := queue.DeQueue()
		fmt.Print(tempNode.data," ")

		if tempNode.leftChild != nil {
			queue.EnQueue(tempNode.leftChild)
		}

		if tempNode.rightChild != nil {
			queue.EnQueue(tempNode.rightChild)
		}
	}
}

func main() {
	root := &TreeNode{data:1}
	root.leftChild = &TreeNode{data:2}
	root.rightChild = &TreeNode{data:3}
	root.leftChild.leftChild = &TreeNode{data:4}
	root.leftChild.rightChild = &TreeNode{data:5}
	PrintLevelOrder(root)
}
