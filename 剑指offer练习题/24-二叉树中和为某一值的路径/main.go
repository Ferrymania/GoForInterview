/*
输入一棵二叉树和一个整数，打印出二叉树中结点值的和为输入整数的所有路径。
从树的根结点开始往下一直到叶结点所经过的结点形成一条路径
 */
package main

import "fmt"

type treeNode struct {
	data int
	leftChild *treeNode
	rightChild *treeNode
}

/*
对二叉树进行先序遍历，把遍历的路径记录下来，当遍历到叶子节点时，判断
当前的路径上所有结点数据的和是否等于给定的整数，如果相等则输出路径信息
 */
//root 二叉树根结点，target 给定的整数,sum 当前路径上所有结点的和，path 存储从根结点遍历到当前结点的路径
func FindPath(root *treeNode,target,sum int,path []int){
	if root == nil {
		return
	}
	sum += root.data
	path = append(path,root.data)

	if root.leftChild == nil && root.rightChild == nil && sum == target {
		//当前节点是叶子节点并且与目标整数相等
		for _,v :=range path {
			fmt.Print(v," ")
		}
		fmt.Println()
	}
	//遍历左子树
	if root.leftChild != nil {
		FindPath(root.leftChild,target,sum,path)
	}
	//遍历右子树
	if root.rightChild != nil {
		FindPath(root.rightChild,target,sum,path)
	}
	//返回父节点之前，在路径上删除当前结点
	sum -= path[len(path)-1]
	path = path[:len(path)-1]

}

func main() {
	root := &treeNode{data:10}
	root.leftChild = &treeNode{data:5}
	root.rightChild = &treeNode{data:12}
	root.leftChild.leftChild = &treeNode{data:4}
	root.leftChild.rightChild = &treeNode{data:7}
	FindPath(root,22,0,make([]int,0))
}

