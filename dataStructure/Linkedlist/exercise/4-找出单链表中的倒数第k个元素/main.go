/*
找出单链表中的倒数第k个元素，
 */
package main

import "fmt"

type node struct {
	value int
	next *node
}
//顺序遍历两遍法。首先遍历一遍单链表，求出整个单链表的长度n,然后把求
//倒数第k个元素转换成求第n-k个元素，再去遍历一次单链表进行两次遍历
func findLastK(head *node,k int)int{
	if head ==nil || head.next == nil {
		panic("Linkedlist is empty")
	}
	length := 0
	for cur := head.next;cur !=nil;cur=cur.next{
		length++
	}
	cur :=head.next
	for i:=0;i<length-k;i++ {
		cur = cur.next
	}
	return cur.value
}
//快慢指针查找。在查找过程中，设置两个指针，让其中一个指针比另一个指针先前移k步，然后两个
//指针同时往前移动。循环直到先行的指针值为null时，另一个指针所指的位置就是所要找的位置。
func FindLastK(head *node,k int) int {
	if head == nil || head.next ==nil {
		panic("linkedList is empty")
	}
	slow := head.next
	fast := head.next
	i :=0
	for;i<k && fast !=nil;i++{
		fast=fast.next
	}
	if i<k{
		return -1
	}
	for fast!=nil {
		slow = slow.next
		fast  = fast.next
	}
	return slow.value
}

func main() {
	head := &node{}
	CreateNode(head)
	PrintNode(head)
	fmt.Println("last 2 is ",findLastK(head,2))
	fmt.Println("last 1 is ",FindLastK(head,1))
}

func CreateNode(head *node){
	cur := head
	arr :=[]int{1,3,6,5,4,7}
	for i:= 0;i<len(arr);i++ {
		cur.next = &node{}
		cur.next.value = arr[i]
		cur = cur.next
	}
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