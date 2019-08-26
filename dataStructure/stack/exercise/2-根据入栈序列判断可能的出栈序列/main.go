/*
输入两个整数序列，其中一个序列表示栈的push顺序，判断另一个序列有没有可能是对应的pop顺序
 */
package main

import (
	"fmt"
	"github.com/Ferrymania/GoForInterview/dataStructure/stack/stackViaArray/arraystack"
)

/*
使用一个栈来模拟入栈顺序
1.把push序列依次入栈，直到栈顶元素等于pop序列的第一个元素，然后栈顶元素出栈
，pop序列移动到第二个元素
2.如果栈顶继续等于pop序列现在的元素，则继续出栈并pop后移；否则对push序列继续入栈
3.如果push序列已经全部入栈，但是POP序列未全部遍历，而且栈顶元素不等于当前pop元素，
那么这个序列不是一个可能的出栈序列。
时间复杂度 O(n) 空间复杂度O(n)
 */

func IsPopSerial(pushArr,popArr []int)bool {
	pushLen := len(pushArr)
	popLen := len(popArr)
	if pushLen == 0 || popLen ==0 || pushLen != popLen{
		return false
	}
	pushIndex :=0
	popIndex := 0

	s := arraystack.Constructor(pushLen)
	for pushIndex<pushLen {
		s.Push(pushArr[pushIndex])
		pushIndex++

		for !s.IsEmpty() && s.Peek() == popArr[popIndex]{
			s.Pop()
			popIndex++
		}

	}
	if s.IsEmpty() && popIndex == popLen {
		return true
	}
	return false
}

func main() {
	pushArr := []int{1,2,3,4,5}
	popArr := []int{3,2,5,4,1}
	fmt.Println("isPOPSerial ",IsPopSerial(pushArr,popArr))
}
