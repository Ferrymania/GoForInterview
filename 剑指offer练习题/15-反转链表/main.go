/*
定义一个函数，输入一个链表的头结点，反转该链表并输出反转后链表的头结点
 */
package main

import "fmt"

type node struct {
	value int
	next *node
}

/*
定义三个指针，分别指向当前遍历到的结点，它的前结点以及后结点
 */
func ReverseList(head *node)*node{
	var pReverseHead *node
	pNode :=head
	var pPrev,pNext *node
	if head == nil || head.next == nil {
		return head
	}
	for pNode != nil {
		pNext = pNode.next
		if pNext == nil {
			pReverseHead = pNode
		}
		pNode.next = pPrev
		pPrev = pNode
		pNode = pNext
	}
	return pReverseHead
}

func main() {
	head := &node{value:1}
	temp :=head
	for i:=1;i<6;i++{
		temp.next = &node{value:i+1}
		temp = temp.next
	}
	temp = head
	for temp!= nil {
		fmt.Print(temp.value," ")
		temp = temp.next
	}

	reverseHead := ReverseList(head)
	fmt.Println()
	temp = reverseHead
	for temp!= nil {
		fmt.Print(temp.value," ")
		temp = temp.next
	}
}