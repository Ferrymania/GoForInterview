/*
翻转栈的所有元素，例如输入栈{1,2,3,4,5}，翻转之后{5,4,3,2,1}
 */
package main

import (
	"fmt"
	"github.com/Ferrymania/GoForInterview/dataStructure/stack/stackViaArray/arraystack"
)

//若申请一个队列存储栈中元素，再出队存入栈中，空间复杂度高
//采用递归方法解决。将当前栈的最底元素移到栈顶，其他元素顺次下移一位，然后对
//不包含栈顶元素的子栈进行同样的操作，递归下去，直到栈为空

func moveBottomToTop(s *arraystack.SliceStack){
	if s.IsEmpty(){
		return
	}
	top1 := s.Pop()  //弹出栈顶元素
	if !s.IsEmpty() {
		moveBottomToTop(s)
		top2 := s.Pop()
		s.Push(top1)
		s.Push(top2)
	}else {
		s.Push(top1)
	}
}
func ReverseStack(s *arraystack.SliceStack){
	if s.IsEmpty(){
		return
	}

	//把栈底元素移到栈顶
	moveBottomToTop(s)
	top:=s.Pop()
	ReverseStack(s)
	s.Push(top)
}

func main() {
	s := arraystack.Constructor(10)
	for i:=1;i<10;i++ {
		s.Push(i*2)
	}
	fmt.Println(s)
	ReverseStack(s)
	fmt.Println(s)
}