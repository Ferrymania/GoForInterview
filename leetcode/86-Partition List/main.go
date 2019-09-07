/*
Given a linked list and a value x, partition it such that all nodes less than x come before nodes greater than or equal to x.
You should preserve the original relative order of the nodes in each of the two partitions.
Example:
Input: head = 1->4->3->2->5->2, x = 3
Output: 1->2->2->4->3->5
 */
package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return  head
	}

	smallHead := &ListNode{}
	bigHead := &ListNode{}
	small := smallHead
	big := bigHead
	for head!= nil {
		if head.Val < x {
			small.Next = head
			small= small.Next
		}else {
			big.Next = head
			big = big.Next
		}
		head = head.Next
	}
	big.Next = nil
	small.Next = bigHead.Next
	return smallHead.Next
}

func main() {
	head := &ListNode{Val:1}
	head.Next = &ListNode{Val:4}
	head.Next.Next = &ListNode{Val:3}
	head.Next.Next.Next = &ListNode{Val:2}
	head.Next.Next.Next.Next = &ListNode{Val:5}
	head.Next.Next.Next.Next.Next = &ListNode{Val:2}
	new := partition(head ,3)
	cur := new
	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.Next

	}
}
