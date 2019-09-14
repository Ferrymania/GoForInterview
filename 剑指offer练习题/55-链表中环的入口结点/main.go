/*
如果一个链表包含环，如何找出环的入口节点
 */
package main

import "fmt"

type node struct {
	data int
	next *node
}
/*
定义快慢指针，如果走的快的指针追上了走得慢的指针，那么链表就包含环
如果走得快的指针走到了链表的末尾（next指向NULL）都没有追上第一个指针，那么链表
就不包含环
找出环中任意一个节点->得到环中节点的数目->找到环的入口节点
 */

//找到环中任意一个节点
func MeetingNode(head *node)*node{
	if head == nil {
		return nil
	}
	pSlow := head.next
	if pSlow == nil {
		return nil
	}
	pFast := pSlow.next
	for pFast!=nil && pSlow != nil {
		if pFast == pSlow {
			return pFast
		}
		pSlow = pSlow.next
		pFast = pFast.next
		if pFast!= nil {
			pFast = pFast.next
		}
	}
	return nil
}
//得出环中节点数目，并找到环的入口节点
func EntryNodeOfLoop(head *node)*node{
	meetingNode := MeetingNode(head)
	if meetingNode == nil {
		return nil
	}

	nodesInLoop :=1
	pNode1 := meetingNode
	for pNode1.next != meetingNode {
		pNode1 = pNode1.next
		nodesInLoop++
	}

	pNode1 = head;
	for i:=0;i<nodesInLoop;i++{
		pNode1 = pNode1.next
	}

	pNode2 := head
	for pNode1 != pNode2{
		pNode1 = pNode1.next
		pNode2 = pNode2.next
	}
	return pNode1
}

func main(){
	head := &node{data:1}
	head.next = &node{data:2}
	head.next.next = &node{data:3}
	head.next.next.next = &node{data:4}
	head.next.next.next.next = &node{data:5}
	head.next.next.next.next.next = head.next.next
	entry := EntryNodeOfLoop(head)
	fmt.Println(entry.data)
}

