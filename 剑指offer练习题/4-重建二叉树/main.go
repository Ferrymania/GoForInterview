/*
输入某二叉树的前序遍历和中序遍历的结果，请重建出该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。例如输入前序遍历序列{1,2,4,7,3,5,6,8}和中序遍历序列{4,7,2,1,5,3,8,6}，则重建二叉树并返回。

 // 思路：前序遍历序列中，某元素右边的序列是该节点左子树和右子树的前序遍历序列；
    //      中序遍历序列中，某元素左边的序列是该节点左子树的中序遍历序列，右边同理
    //      根据上述特性，找出该节点左右子树的前序和中序遍历序列，然后可以建立递归关系
 */
package main

import "fmt"

type TreeNode struct {
	Value int
	Left *TreeNode
	Right *TreeNode
}

//pre前序遍历  in中序遍历
func reConstructBinaryTree(pre ,in []int) *TreeNode{
	if(len(pre)==0 || len(in)==0){
		return nil
	}
	return ConstructBinaryTree(pre,0,len(pre)-1,in,0,len(in)-1)
}
/*
pre  前序遍历数组
preStart  前序遍历起始索引
preEnd  前序遍历数组最大索引
in   中序遍历数组
inStart 中序遍历起始索引
inEnd  中序遍历数组最大索引
 */
func ConstructBinaryTree(pre []int,preStart int,preEnd int,in []int,inStart int ,inEnd int)*TreeNode{
	//递归截止条件
	if preStart>preEnd || inStart > inEnd {
		return nil
	}
	//根据前序遍历获取根节点
	root := &TreeNode{Value:pre[preStart]}
	//遍历查找在中序遍历中的根
	for i := inStart;i<=inEnd;i++ {
		if in[i] == pre[preStart] {
			root.Left = ConstructBinaryTree(pre,preStart+1,preStart+i-inStart,in,inStart,i-1)
			root.Right = ConstructBinaryTree(pre,preStart+i-inStart+1,preEnd,in,i+1,inEnd)
		}
	}
	return root
}
func main() {
	pre := []int{1,2,4,7,3,5,6,8}
	in := []int{4,7,2,1,5,3,8,6}
	root := reConstructBinaryTree(pre,in)
	fmt.Println(root.Value)
}
