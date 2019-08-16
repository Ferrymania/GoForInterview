/*
从无序列表中移除重复项
example 1->3->1->5->5->7 to 1->3->5->7
 */


package main

import (
	"fmt"
)

type node struct {
	Value int
	Next *node
}
//顺序删除。利用两层循环，外层循环用一个指针从第一个结点开始遍历整个链表，然后内层循环用另一个指针遍历其余结点，将与外层循环遍历到的指针所指结点的数据域相同的结点删除。
//时间复杂度O(n^2) 空间复杂度O(1)
func removeDup(head *node){
	if head == nil || head.Next == nil {
		return
	}
	for outerCur:=head.Next;outerCur != nil ;outerCur = outerCur.Next{
		innerPre :=outerCur
		innerCur := outerCur.Next
		for ;innerCur!=nil;{
			if innerCur.Value == outerCur.Value {
				innerPre.Next = innerCur.Next
				innerCur = innerCur.Next
			}else {
				innerPre = innerCur
				innerCur = innerCur.Next
			}
		}
	}
}

func main() {
	head := &node{}
	CreateNode(head)
	PrintNode(head)
	removeDup(head)
	PrintNode(head)
}

func CreateNode(head *node){
	cur := head
	arr :=[]int{1,3,1,5,5,7}
	for i:= 0;i<len(arr);i++ {
		cur.Next = &node{}
		cur.Next.Value = arr[i]
		cur = cur.Next
	}
}

func PrintNode(head *node){
	for cur:=head.Next;cur != nil;cur = cur.Next{
		if cur.Next != nil {
			fmt.Print(cur.Value,"->")
		}else{
			fmt.Print(cur.Value,"\n")
		}
	}
}