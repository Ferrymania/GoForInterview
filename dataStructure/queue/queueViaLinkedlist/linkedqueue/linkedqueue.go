/*
分别使用两个指针指向队列的首元素与尾元素
 */
package linkedqueue

import (
	"bytes"
	"fmt"
)

type node struct {
	value int
	next *node
}

type LinkedQueue struct {
	head  *node
	end  *node
}

func(q *LinkedQueue) IsEmpty() bool{
	return q.head == nil
}

func(q *LinkedQueue) GetSize() int {
	if q.IsEmpty() {
		return 0
	}
	size :=0
	for cur:=q.head;cur !=nil;cur = cur.next{
		size++
	}
	return size
}

func(q *LinkedQueue)EnQueue(value int){
	e := &node{value:value}
	if q.IsEmpty() {
		q.head = e
		q.end = e
	}else {
		q.end.next = e
		q.end = e
	}
}

func(q *LinkedQueue)DeQueue(){
	if q.head == nil {
		panic("queue is already empty")
	}
	q.head = q.head.next
	if q.head == nil {
		q.end = nil
	}
}

func(q *LinkedQueue)GetFront() int {
	if q.IsEmpty() {
		panic ("queue is empty")
	}
	return q.head.value
}

func(q *LinkedQueue)GetBack() int {
	if q.IsEmpty() {
		panic("queue is empty")
	}
	return q.end.value
}

func(q *LinkedQueue) String() string {
	var buffer bytes.Buffer
	if q.IsEmpty() {
		buffer.WriteString("queue is nil")
	}
	buffer.WriteString("front [")

	for cur:=q.head;cur!=nil;cur = cur.next {
		buffer.WriteString(fmt.Sprint(cur.value))
		if cur.next != nil{
			buffer.WriteString(" ")
		}
	}
	buffer.WriteString("] back")
	return buffer.String()
}