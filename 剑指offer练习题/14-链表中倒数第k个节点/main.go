/*
输入一个链表，输出该链表中倒数第K个结点
 */
package main

import "fmt"

/*
方法一：若有n个结点，倒数第k个结点即为顺数第n-k+1个结点，需遍历两次
方法二： 双指针法。定义两个指针，第一个指针从链表的头指针向前走k-1,第二个指针保持不动
从第k步开始，第二个指针也开始从链表的头指针开始遍历。由于两个指针的距离保持在k-1，当第一个
指针到达链表的尾结点时，第二个指针正好是倒数第k个结点
 */

type node struct {
	value int
	next *node
}

func FindKthToTail(head *node,k int)*node{
	if head == nil || k ==0 {
		return nil
	}
	pAhead := head
	pBehind := head
	for i:=0;i<k-1;i++{
		//判断k是不是大于链表长度
		if pAhead.next != nil {
			pAhead = pAhead.next
		}else {
			return  nil
		}

	}
	for pAhead.next!= nil{
		pAhead = pAhead.next
		pBehind = pBehind.next
	}
	return pBehind
}

func main(){
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

	kNode := FindKthToTail(head,1)
	fmt.Println()
	fmt.Println(kNode.value)
}