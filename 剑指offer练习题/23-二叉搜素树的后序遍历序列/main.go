/*
输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历的结果。如果是则返回true，否则
返回false。假设输入的数组的任意两个数字都互不相同
*/
package main

import "fmt"

type treeNode struct {
	data int
	leftChild *treeNode
	rightChild *treeNode
}

/*
二叉查找树的特点：对于任意一个结点，它的左子树上所有结点的值都小于这个结点的值，它的右子树上所有结点的值都
大于这个结点的值。
从而在后序遍历得到的序列中，最后一个数字是树的根结点的值。数组中前面的数字可以分为两部分：第一部分是左子树
根结点的值，它们都比根结点的值小；第二部分是右子树结点的值，它们都比根结点的值大
 */
//start:首元素 end 尾元素
func VerifySquenceOfBST(sequence []int,start,end int)bool {
	if sequence == nil {
		return  false
	}
	if len(sequence) == 1 {
		return  true
	}
	root :=sequence[end] //最后一个结点必定是根结点
	var i,j int
	//找到第一个大于root的值，那么前面所有的结点都位于root的左子树上
	for i=start;i<end;i++{
		if sequence[i]>root {
			break
		}
	}
	//如果序列是后序遍历的序列，那么从i开始的所有值都应该大于根结点root的值
	for j=i;j<end;j++{
		if sequence[j] <root{
			return false
		}
	}
	leftIsPostOrder := true
	rightIsPostOrder := true
	if i>start {
		leftIsPostOrder = VerifySquenceOfBST(sequence,start,i-1)
	}

	if j<end {
		rightIsPostOrder = VerifySquenceOfBST(sequence,i,end-1)
	}
	return leftIsPostOrder && rightIsPostOrder
}

func main(){
	sequence := []int{1,3,2,5,7,6,4}
	fmt.Println("Is Post oreder of BST",VerifySquenceOfBST(sequence,0,len(sequence)-1))
}