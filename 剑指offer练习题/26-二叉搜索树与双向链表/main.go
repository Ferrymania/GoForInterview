/*
输入一棵二叉搜索树，将该二叉树转换成一个排序的双向链表。要求不能创建任何新的节点，只能
调整树中结点指针的指向。
example
      10
	6   14      -> 4 =6 =8 = 10 =12 =14 =16
  4  8 12 16
 */
package main

type BST struct {
	data int
	leftChild *BST
	rightChild *BST
}

/*
在二叉搜索树中，左子结点的值总是小于父结点的值，右子节点的值总是大于父结点的值。因此我们在
转换成排序双向链表时，原先指向左子结点的指针调整为链表中指向前一个结点的指针，原先指向右子
节点的指针调整为链表中指向后一个结点的指针。由于要求链表是有序的，可以借助二叉树中序遍历，
因为中序遍历算法的特点就是从小到大访问结点。当遍历访问到根结点时，假设根结点的左侧已经处理
好，只需将根结点与上次访问的最近结点（左子树中最大值结点）的指针连接好即可。进而更新当前链
表的最后一个结点指针。同时中序遍历过程正好是转换成链表的过程，可采用递归方法处理。
 */
func convert(root *BST)*BST{
	head := &BST{}
	end := &BST{}  //一直保存前一个结点
	convertRecursively(root,head,end) //中序遍历
	return head
}

func convertRecursively(root ,head,end *BST){
	if root == nil {
		return
	}
	convert(root.leftChild)  //遍历左子树，直到最左边的叶子节点
	if end == nil {
		head = root
		end = root
	}else {
		end.rightChild = root
		root.leftChild = end
		end = root
	}
	convert(root.rightChild)
}
