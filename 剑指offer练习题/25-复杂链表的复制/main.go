/*
请实现一个函数复制一个复杂链表。在复杂链表中，每个结点除了有一个next指针指向下一个结点外，还有
一个sibling指向链表中的任意节点或者null
 */
package main

type ComplexListNode struct {
	value int
	next *ComplexListNode
	sibling *ComplexListNode
}

/*
第一步根据原始链表的每个结点N创建对应的N'，并把N'链接在N的后面
A->A'->B->B'->C->C'
 */
func CloneNodes(head *ComplexListNode){
	pNode := head
	for pNode != nil {
		pCloned := &ComplexListNode{}
		pCloned.value = pNode.value
		pCloned.next = pNode.next

		pNode.next = pCloned
		pNode = pCloned.next
	}
}
/*
第二步设置复制出来的结点的sibling。假设原始链表上的N的sibling指向结点S，
那么其对应复制出来的N'是N的next指向的结点，同样S'也是S的next指向的结点。
N->S
N'->S'
 */
func ConnectSiblingNodes(head *ComplexListNode){
	pNode :=head
	for pNode != nil {
		pCloned := pNode.next
		if  pNode.sibling != nil {
			pCloned.sibling = pNode.sibling.next
		}
		pNode = pCloned.next
	}
}
/*
第三步把这个长链表拆分成两个链表，把奇数位置的结点用next链接起来就是原始链表，
把偶数位置的结点用next链接起来就是复制出来的链表
 */
func ReconnectNodes(head *ComplexListNode)*ComplexListNode{
	pNode := head
	var PClonedHead,pClonedNode *ComplexListNode
	if pNode != nil {
		PClonedHead = pNode.next
		pClonedNode = pNode.next
		pNode.next = pClonedNode.next
		pNode = pNode.next
	}
	for pNode != nil {
		pClonedNode.next = pNode.next
		pClonedNode = pClonedNode.next
		pNode.next = pClonedNode.next
		pNode = pNode.next
	}
	return PClonedHead
}


func Clone(head *ComplexListNode)*ComplexListNode{
	CloneNodes(head)
	ConnectSiblingNodes(head)
	return ReconnectNodes(head)
}