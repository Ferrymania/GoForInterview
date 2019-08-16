/*
输入两个链表，找出它们的第一个公共结点。
 */
package main

import "fmt"

type node struct {
	value int
	next *node
}

//蛮力法。在第一个链表上顺序遍历每个结点，每遍历到一个结点的时候,在第二个链表上顺序遍历每个结点，
//如果在第二个链表上有一个结点和第一个链表上的结点一样，说明两个链表在这个结点上重合。
//时间复杂度 O(mn)

//比较简单的方法。首先遍历两个链表得到它们的长度，知道较长的的链表多几个结点。在第二次遍历的时候，
//在较长的链表上先走若干步，接着再同时在两个链表上遍历，找到的第一个相同的结点就是它们的第一个公共结点
func isIntersect(h1,h2 *node) *node {
	length1 := getListLength(h1)
	length2 := getListLength(h2)
	lengthDiff :=0
	var headLong,headShort *node
	if length1 > length2 {
		lengthDiff = length1 - length2
		headLong = h1
		headShort = h2
	}else {
		lengthDiff = length2 - length1
		headLong = h2
		headShort = h1
	}

	for i:=0;i<lengthDiff;i++ {
		headLong = headLong.next
	}

	for headLong != nil && headShort != nil && headLong!= headShort {
		headLong = headLong.next
		headShort = headShort.next
	}

	commonNode := headLong
	return commonNode
}

func getListLength(head *node) int {
	if head == nil || head.next ==nil {
		return 0
	}
	length := 0
	for cur:=head.next;cur !=nil;cur = cur.next{
		length ++
	}
	return length
}

func main() {
	head1 := &node{}
	retNode :=CreateLinkedList(head1,5)
	head2 := &node{}
	cur :=head2
	for i:=1;i<5;i++ {
		cur.next = &node{}
		cur.next.value = i
		cur = cur.next
	}
	cur.next = retNode
	PrintNode(head1)
	PrintNode(head2)
	interNode :=isIntersect(head1,head2)
	if interNode == nil {
		fmt.Println("these two linkedlist not intersect")
	}else {
		fmt.Println("these two linkedlist intersected in",interNode.value)
	}

}

func CreateLinkedList(head *node,intersect int)(retNode *node){
	cur := head
	for i:=1;i<8;i++{
		cur.next = &node{}
		cur.next.value = i+2
		cur = cur.next
		if i == intersect {
			retNode = cur
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