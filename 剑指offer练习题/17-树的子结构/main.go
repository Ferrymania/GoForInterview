/*
输入两棵二叉树A和B，判断B是不是A的子结构。
 */
package main

import "fmt"

/*
分成两步：第一步在树A中找到和B的根节点的值一样的结点R，
第二步再判断A中以R为根结点的子数是不是包含和树B一样的结构
 */

type treeNode struct {
	data int
	leftChild *treeNode
	rightChild *treeNode
}
func HasSubTree(aRoot,bRoot *treeNode)bool{
	result := false
	if aRoot!= nil && bRoot != nil {
		if aRoot.data == bRoot.data{           //比较根节点
			result = DoesTreeAHaveTreeB(aRoot,bRoot)
		}
		if !result {
			result = HasSubTree(aRoot.leftChild,bRoot)  //继续比较左子结点
		}
		if !result {
			result = HasSubTree(aRoot.rightChild,bRoot) //继续比较右子结点
		}
	}
	return result
}

func DoesTreeAHaveTreeB(aRoot,bRoot *treeNode)bool{
	if bRoot == nil {
		return true
	}
	if aRoot == nil {
		return false
	}
	if aRoot.data != bRoot.data{
		return false
	}
	return DoesTreeAHaveTreeB(aRoot.leftChild,bRoot.leftChild)&&DoesTreeAHaveTreeB(aRoot.rightChild,bRoot.rightChild)
}


func main() {
	treeA := &treeNode{data:8}
	treeA.leftChild = &treeNode{data:8}
	treeA.rightChild = &treeNode{data:7}
	treeA.leftChild.leftChild = &treeNode{data:9}
	treeA.leftChild.rightChild = &treeNode{data:2}
	treeA.leftChild.rightChild.leftChild = &treeNode{data:4}
	treeA.leftChild.rightChild.rightChild = &treeNode{data:7}

	treeB :=&treeNode{data:8}
	treeB.leftChild = &treeNode{data:9}
	treeB.rightChild = &treeNode{data:2}

	fmt.Println("Does treeA has treeB ",HasSubTree(treeA,treeB))
}