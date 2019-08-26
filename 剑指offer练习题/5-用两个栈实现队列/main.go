/*
用两个栈实现一个队列，完成在队列尾部插入节点和在队列头部删除结点的功能
 */
package main

import (
	"fmt"
	"github.com/Ferrymania/GoForInterview/dataStructure/stack/stackViaArray/arraystack"
)

/*
使用栈A和栈B模拟队列Q，A为插入栈，B为弹出栈，以实现队列Q
入队列，入栈A即可，出队列要考虑
（1） 如果栈B不为空，则直接弹出栈B的数据
(2)如果栈B为空，则依次弹出栈A的数据，放入栈B中，再弹出栈B的数据
 */

type StackQueue struct {
	aStack *arraystack.SliceStack
	bStack *arraystack.SliceStack
}

func(q *StackQueue)EnQueue(data int){
	q.aStack.Push(data)
}

func(q *StackQueue)DeQueue() int {
	if q.bStack.IsEmpty() {
		for !q.aStack.IsEmpty(){
			q.bStack.Push(q.aStack.Pop())
		}
	}

	return q.bStack.Pop().(int)
}

func main() {
	q := &StackQueue{aStack:&arraystack.SliceStack{},bStack:&arraystack.SliceStack{}}
	q.EnQueue(2)
	q.EnQueue(5)
	fmt.Println("queue frist element is ",q.DeQueue())
	fmt.Println("queue first element is ",q.DeQueue())
}