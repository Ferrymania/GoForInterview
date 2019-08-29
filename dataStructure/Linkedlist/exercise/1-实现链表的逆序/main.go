/*
给定一个带头节点的单链表，请将其逆序
example：head->1->2->3->4->5    to  head->5->4->3->2->1
 */
package main

import (
	"fmt"
	"github.com/Ferrymania/GoForInterview/dataStructure/Linkedlist/singlylinkedlist"
)
// 就地逆序，修改当前结点指针域的指向，让其指向它的前驱结点
func Reverse(head *singlylinkedlist.Node){
	if head == nil || head.Next == nil {
		return
	}

	var pre ,cur *singlylinkedlist.Node //定义前驱结点和当前结点
	next := head.Next  //后继结点     pre-> cur ->next
	for next != nil {
		cur = next.Next				// pre<-cur next
		next.Next = pre
		pre = next
		next = cur
	}
	head.Next = pre //修改头结点的指向
}
//递归法：先逆序除第一个结点以外的子链表，接着把第一个结点添加到逆序的子链表的后面。
func recursiveReverseChild(node *singlylinkedlist.Node)*singlylinkedlist.Node {
	if node != nil || node.Next != nil {
		return node
	}
	newHead := recursiveReverseChild(node.Next)
	node.Next.Next = node
	node.Next = nil
	return newHead
}

func recursiveReverse(node *singlylinkedlist.Node){
	firstNode := node.Next
	newHead := recursiveReverseChild(firstNode)
	node.Next = newHead
}

func main(){
	head := &singlylinkedlist.Node{}
	CreateNode(head,8)
	PrintNode(head)
	Reverse(head)
	PrintNode(head)
	recursiveReverse(head)
	PrintNode(head)
}

func CreateNode(node *singlylinkedlist.Node,length int){
	cur := node
	for i:= 1;i<length;i++ {
		cur.Next = &singlylinkedlist.Node{}
		cur.Next.Value = i
		cur = cur.Next
	}
}

func PrintNode(head *singlylinkedlist.Node){
	for cur:=head.Next;cur != nil;cur = cur.Next{
		if cur.Next != nil {
			fmt.Print(cur.Value,"->")
		}else{
			fmt.Print(cur.Value,"\n")
		}
	}
}