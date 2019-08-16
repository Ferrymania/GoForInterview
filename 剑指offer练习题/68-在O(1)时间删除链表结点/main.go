/*
给定单向链表的头指针和一个结点指针，定义一个函数在O(1)时间
删除该节点
 */
package main

import "fmt"

type node struct {
	value int
	next *node
}
//删除一个结点，通常从链表的头结点开始，顺序遍历查找要删除的结点，
//并在链表中删除该结点，时间复杂度为O(n)

//如果把下一个结点的内容复制到需要删除的结点上覆盖原有的内容，
//再把下一个结点删除，就相当于把需要删除的结点删除
//需要注意的是删除的如果是链表的尾结点和一个结点的情况
//时间复杂度[(n-1)*O(1)+O(n)]/n  = O(1)
func deleteNode(head ,toBeDeleted *node){
	if head == nil|| head.next ==nil || toBeDeleted ==nil {
		return
	}

	//删除结点不是尾结点
	if toBeDeleted.next != nil {
		pNext := toBeDeleted.next
		toBeDeleted.value = pNext.value
		toBeDeleted.next = pNext.next
		pNext =nil
	}else if head.next == toBeDeleted { //只有一个结点
		toBeDeleted = nil
		head.next = nil
	}else {  //链表中有多个结点，删除尾结点
		cur  := head.next
		for cur.next != toBeDeleted {
			cur = cur.next
		}
		cur.next = nil
		toBeDeleted = nil
	}
}

func main() {
	head := &node{}
	toBeDeleted := CreateNode(head)
	PrintNode(head)
	deleteNode(head,toBeDeleted)
	PrintNode(head)
}

func CreateNode(head *node)(toBeDeleted *node){
	cur := head
	arr :=[]int{1,3,2,6,8,7}
	for i:= 0;i<len(arr);i++ {
		cur.next = &node{}
		cur.next.value = arr[i]
		cur = cur.next
		if i ==4 {
			toBeDeleted = cur
		}
	}
	return
}

func PrintNode(head *node){
	for cur:=head.next;cur != nil;cur = cur.next{
		if cur.next != nil {
			fmt.Print(cur.value,"->")
		}else{
			fmt.Print(cur.value,"\n")
		}
	}
}