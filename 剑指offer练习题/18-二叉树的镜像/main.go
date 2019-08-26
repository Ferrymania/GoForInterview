/*
输入一个二叉树，输出它的镜像
镜像定义：
          4                            4
	2		   6                    6     2
1	  3  	5	   7             7    5  3  1
 */
package main

import "fmt"

type treeNode struct {
	data int
	leftChild *treeNode
	rightChild *treeNode
}

func MirrorRecursively(root *treeNode){
	//根节点为空或没有叶子结点 直接返回
	if root ==nil || root.leftChild == nil && root.rightChild ==nil  {
		return
	}
	var temp *treeNode
	temp = root.leftChild
	root.leftChild = root.rightChild
	root.rightChild = temp
	if root.leftChild != nil {
		MirrorRecursively(root.leftChild)
	}
	if root.rightChild != nil {
		MirrorRecursively(root.rightChild)
	}
}

func main(){
	root := &treeNode{data:8}
	root.leftChild = &treeNode{data:6}
	root.rightChild = &treeNode{data:10}
	root.leftChild.leftChild = &treeNode{data:5}
	root.leftChild.rightChild = &treeNode{data:7}
	root.rightChild.leftChild = &treeNode{data:9}
	root.rightChild.rightChild = &treeNode{data:11}
	MirrorRecursively(root)
	fmt.Println("\t\t\t",root.data)
	fmt.Println("\t\t",root.leftChild.data,"\t",root.rightChild.data)
	fmt.Println("\t",root.leftChild.leftChild.data,"  ",root.leftChild.rightChild.data,"  ",root.rightChild.leftChild.data,"  ",root.rightChild.rightChild.data)
}
