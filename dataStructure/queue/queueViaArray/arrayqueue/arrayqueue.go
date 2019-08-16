/*
arrayqueue make a simple queue by slice
 */
package arrayqueue

import (
	"bytes"
	"fmt"
)

type Queue struct {
	arr []int
	front int  //record the first  element location of queue
	rear int   //record the last element location of queue + 1
}

func CreateQueue(capacity int) *Queue {
	return &Queue{arr:make([]int,0,capacity)}
}
//Determine if the queue is empty
func(q *Queue)IsEmpty() bool {
	return q.front == q.rear
}
//return the size of queue
func(q *Queue)Size() int {
	return q.rear - q.front
}
//Get the front element of queue
func(q *Queue)GetFront()int{
	if q.IsEmpty() {
		panic("Get front error.queue is empty")
	}
	return q.arr[q.front]
}

//Get the back element of the queue
func(q *Queue)GetBack() int{
	if q.IsEmpty() {
		panic("Get back error.queue is empty")
	}
	return q.arr[q.rear-1]
}
//Delete the front element of the queue
func(q *Queue)DeQueue() int {
	if q.IsEmpty() {
		panic("Delete front error.queue is empty")
	}
	front := q.arr[0]
	q.arr = q.arr[1:]
	q.rear--
	return front
}
//insert element to the back of the queue
func(q *Queue)EnQueue(e int){
	q.arr = append(q.arr,e)
	q.rear++
}
//rewrite the queue type string method
func(q *Queue) String() string {
	var buffer bytes.Buffer
	if q.IsEmpty() {
		buffer.WriteString("queue is nil")
	}
	buffer.WriteString("front [")
	for i:=0;i<q.Size();i++ {
		buffer.WriteString(fmt.Sprint(q.arr[i]))
		if i != q.Size()-1{
			buffer.WriteString(" ")
		}
	}
	buffer.WriteString("] back")
	return buffer.String()
}
