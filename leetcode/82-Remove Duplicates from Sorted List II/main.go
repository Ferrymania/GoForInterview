/*
Given a sorted linked list, delete all nodes that have duplicate numbers, leaving only distinct numbers from the original list.
Input: 1->2->3->3->4->4->5
Output: 1->2->5
Input: 1->1->1->2->3
Output: 2->3
 */
package main

import "fmt"

type ListNode struct{
	Val int
	Next *ListNode
}
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	fakeHead := &ListNode{}
	fakeHead.Next = head
	pre := fakeHead
	cur := head
	for cur != nil{
		for cur.Next != nil && cur.Val == cur.Next.Val{
			cur = cur.Next
		}
		if pre.Next == cur {
			pre = pre.Next
		}else {
			pre.Next = cur.Next
		}
		cur = cur.Next
	}
	return fakeHead.Next
}

func main(){
	head := &ListNode{Val:1}
	head.Next = &ListNode{Val:1}
	head.Next.Next= &ListNode{Val:2}
	head.Next.Next.Next = &ListNode{Val:3}
	head.Next.Next.Next.Next = &ListNode{Val:4}
	head.Next.Next.Next.Next.Next = &ListNode{Val:4}
	cur := head
	for cur != nil {
		fmt.Print(cur.Val," ")
		cur = cur.Next
	}
	fmt.Println()
	new := deleteDuplicates(head)
	cur = new
	for cur != nil {
		fmt.Print(cur.Val," ")
		cur = cur.Next
	}
}