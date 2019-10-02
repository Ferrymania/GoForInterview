/*
借助栈，将二叉树的递归遍历算法转换为非递归算法。
先扫描（并非访问）根结点的所有左节点并将它们一一进栈。然后出栈一个结点*p(*p没有左孩子结点
或者左孩子结点均已访问过)，则访问它。然后扫描该结点的右孩子结点，将其进栈，再扫描该右孩子的
结点，将其进栈，再扫描该右孩子结点的所有左节点并一一进栈，如此继续，直到栈为空为止。
*/
package main

import "fmt"

type node struct {
	data int
	leftChild *node
	rightChild *node
}

type stack struct {
	arr []*node
	size int
}
func constructor()*stack{
	return &stack{arr:[]*node{}}
}
func(s *stack)isEmpty()bool{
	return s.size == 0
}
func(s *stack)push(n *node){
	s.arr = append(s.arr,n)
	s.size++
}
func(s *stack)pop()*node{
	if s.isEmpty(){
		return nil
	}
	s.size--
	n := s.arr[s.size]
	s.arr = s.arr[:s.size]
	return n
}
func preTravsal(root *node){
	s:=constructor()
	if root == nil {
		return
	}
	n := root //n是遍历指针
	for !s.isEmpty() || n!= nil {
		//一直遍历到左子树最下边，边遍历边保存根结点到栈中
		if n!= nil  {
			s.push(n)
			n = n.leftChild
		}else {
			// 根指针退站，访问根结点，遍历右子树
			n = s.pop()
			fmt.Println(n.data)
			n = n.rightChild
		}

	}
}

func main() {
	root := &node{data:1}
	root.leftChild = &node{data:2}
	root.rightChild = &node{data:3}
	root.leftChild.leftChild = &node{data:4}
	root.leftChild.rightChild = &node{data:5}
	fmt.Println("InOrder :")
	inTravsal(root)
}