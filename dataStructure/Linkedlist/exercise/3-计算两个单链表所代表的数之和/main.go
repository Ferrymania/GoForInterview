/*
给定两个单链表，链表的每个结点代表一位数，计算两个数的和。例如，输入链表（3->1->5)和
(5->9->2)输出8->0->8，注意个位数在链表头
 */
package main

import (
	"fmt"
)

type node struct {
	value int
	next *node
}
//对链表中的结点直接进行相加操作，把相加的和存储到新的链表中对应的结点中，同时还要记录结点相加后的进位
//时间复杂度O(n) 空间复杂度O(n)
func add(h1 ,h2 *node)*node{
	if h1 == nil || h1.next == nil {
		return h2
	}
	if h2 == nil || h2.next ==nil {
		return h1
	}
	c := 0 //记录进位
	sum :=0 //记录两个结点相加的值
	p1 := h1.next  //遍历h1
	p2 := h2.next  //遍历h2
	resultHead := &node{}  //相加后链表头节点
	p := resultHead  //指向新链表最后一个结点
	for p1 != nil && p2 != nil {
		p.next = &node{}
		sum = p1.value + p2.value+c
		p.next.value = sum %10  //两结点相加和
		c = sum/10  //进位
		p = p.next
		p1 = p1.next
		p2 = p2.next
	}

	//链表h2比h1 长，只考虑h2剩余的结点
	if p1 == nil {
		for p2 != nil {
			p.next = &node{}
			sum = p2.value +c
			p.next.value = sum %10
			c =sum/10
			p = p.next
			p2 = p2.next
		}
	}
	//链表h1比h2 长，只考虑h1剩余的结点
	if p2 == nil {
		for p1 != nil {
			p.next = &node{}
			sum = p1.value +c
			p.next.value = sum %10
			c =sum/10
			p = p.next
			p1 = p1.next
		}
	}
	//两链表所有结点计算完还有进位，增加新的节点
	if c==1 {
		p.next = &node{}
		p.next.value = 1
	}
	return resultHead
}

func main() {
	h1,h2 := createList()
	PrintNode(h1)
	PrintNode(h2)
	s := add(h1,h2)
	PrintNode(s)
}

func createList()(h1,h2 *node){
	h1 =&node{}
	h2 =&node{}
	cur :=h1
	for i:=1;i<7;i++ {
		cur.next = &node{}
		cur.next.value = i+2
		cur = cur.next
	}
	cur = h2
	for i:=9;i>4;i-- {
		cur.next = &node{}
		cur.next.value = i
		cur = cur.next
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
