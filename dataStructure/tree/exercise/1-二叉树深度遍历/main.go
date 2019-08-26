/*
实现二叉树的深度遍历，包括前序遍历、中序遍历和后序遍历
 */
package main

import "fmt"

type node struct {
	data int
	leftChild *node
	rightChild *node
}

//前序遍历
func PreOder(root *node){
	if root == nil {
		return
	}
	fmt.Print(root.data," ")
	PreOder(root.leftChild)
	PreOder(root.rightChild)
}

//中序遍历
func InOrder(root *node){
	if root == nil {
		return
	}
	InOrder(root.leftChild)
	fmt.Print(root.data," ")
	InOrder(root.rightChild)
}

//后序遍历
func PostOrder(root *node){
	if root == nil {
		return
	}
	PostOrder(root.leftChild)
	PostOrder(root.rightChild)
	fmt.Print(root.data," ")
}

func main() {
	root := &node{data:1}
	root.leftChild = &node{data:2}
	root.rightChild = &node{data:3}
	root.leftChild.leftChild = &node{data:4}
	root.leftChild.rightChild = &node{data:5}

	fmt.Println("preOrder :")
	PreOder(root)
	fmt.Println("\nInOrder :")
	InOrder(root)
	fmt.Println("\npostOrder :")
	PostOrder(root)
}