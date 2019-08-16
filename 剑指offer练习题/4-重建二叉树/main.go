/*
输入某二叉树的前序遍历和中序遍历的结果，请重建出该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。例如输入前序遍历序列{1,2,4,7,3,5,6,8}和中序遍历序列{4,7,2,1,5,3,8,6}，则重建二叉树并返回。

 // 思路：前序遍历序列中，某元素右边的序列是该节点左子树和右子树的前序遍历序列；
    //      中序遍历序列中，某元素左边的序列是该节点左子树的中序遍历序列，右边同理
    //      根据上述特性，找出该节点左右子树的前序和中序遍历序列，然后可以建立递归关系
 */
package main

type TreeNode struct {
	Value int
	Left *TreeNode
	Right *TreeNode
}

func reConstructBinaryTree(pre ,in []int) *TreeNode{
	if(len(pre)==0 || len(in)==0){
		return nil
	}
	node := &TreeNode{Value:pre[0]}
	for i:=range in {
		node.Left
	}


}
func main() {

}
