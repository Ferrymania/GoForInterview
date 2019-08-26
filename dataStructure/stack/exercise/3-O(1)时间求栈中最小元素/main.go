/*
在O(1)时间内求栈中最小元素
 */
package main

import (
	"fmt"
	"github.com/Ferrymania/GoForInterview/dataStructure/stack/stackViaArray/arraystack"
)

/*
利用空间换取时间。在实现的时候使用两个栈结构，一个栈用来存储数据，另外一个栈用来存储栈的最小元素。
如果当前入栈的元素比原来栈中的最小值还小，则把这个值压入保存最小元素的栈中；
在出栈的时候，如果当前出栈的元素恰好为当前栈中的最小值，则保存最小值的栈顶元素也出栈，使得当前最小值变为当前最小值入栈前的那个最小值
 */

type MinStack struct {
	elemStack *arraystack.SliceStack
	minStack *arraystack.SliceStack
}

func(p *MinStack)Push(data int){
	p.elemStack.Push(data)
	if p.minStack.IsEmpty() {
		p.minStack.Push(data)
	}else {
		if data <= p.minStack.Peek().(int){
			p.minStack.Push(data)
		}
	}
}

func(p *MinStack)Pop() int {
	top := p.elemStack.Pop().(int)
	if top == p.Min() {
		p.minStack.Pop()
	}
	return top
}

func(p *MinStack)Min() int {
	if p.minStack.IsEmpty(){
		return 0
	}else {
		return p.minStack.Peek().(int)
	}
}

func main() {
	s := &MinStack{elemStack:&arraystack.SliceStack{},minStack:&arraystack.SliceStack{}}
	s.Push(5)
	fmt.Println("stack min is ",s.Min())
	s.Push(2)
	s.Push(4)
	s.Push(6)
	fmt.Println("stack min is ",s.Min())
}