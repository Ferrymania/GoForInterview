/*
输入两个递增排序的链表，合并这两个链表并使新链表中的结点仍是按照递增排序的
 */
package main

import "fmt"

type node struct {
	value int
	next *node
}

func merge(head1,head2 *node) *node {
	if head1 == nil {
		return head2
	}else if head2 ==nil{
		return head1
	}

	mergeHead := &node{}
	if head1.value < head2.value {
		mergeHead = head1
		mergeHead.next = merge(head1.next,head2)
	}else {
		mergeHead = head2
		mergeHead.next = merge(head1,head2.next)
	}
	return mergeHead
}

func main() {
	head1 := &node{}
	CreateNode(head1,7,true)
	PrintNode(head1)
	head2 :=&node{}
	CreateNode(head2,7,false)
	PrintNode(head2)
	merge := merge(head1,head2)
	PrintNode(merge)

}

func CreateNode(head *node ,end int,isOdd bool){
	head.value = 2
	for i:= 2;i<end;i++{
		head.next = &node{}
		if isOdd {
			head.next.value = i*2+1
		}else {
			head.next.value = i*2
		}

		head = head.next
	}
}

func PrintNode(head *node){
	for cur:=head;cur != nil;cur = cur.next{
		if cur.next != nil {
			fmt.Print(cur.value,"->")
		}else{
			fmt.Print(cur.value,"\n")
		}
	}
}