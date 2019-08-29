/*
输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否为该栈
的弹出顺序。假设压入栈的所有数字均不相等
 */
package main

import (
	"fmt"
	"github.com/Ferrymania/GoForInterview/dataStructure/stack/stackViaArray/arraystack"
)

/*
判断一个序列是不是栈的弹出序列的规律：如果下一个弹出的数字刚好是栈顶数字，那么直接弹出。
如果下一个弹出的数字不在栈顶，我们把压栈序列中还没有入栈的数字压入辅助栈，直到把下一个需要弹出的数字压入栈顶为止。
如果所有的数字都压入栈了仍然没有找到下一个弹出的数字，那么该序列不可能是一个弹出序列
 */

func IsPopOrder(arrPush,arrPop []int)bool {
	if arrPop == nil || arrPush == nil || len(arrPush)!= len(arrPop){
		return  false
	}
	stack :=arraystack.Constructor(len(arrPush))
	popIndex := 0
	for i:=0;i<len(arrPush);i++{
		stack.Push(arrPush[i])          //按顺序压入每个元素
		for !stack.IsEmpty() && stack.Peek() == arrPop[popIndex]{  //如果栈顶元素等于当前出栈元素
			stack.Pop()
			popIndex++
		}
	}
	return stack.IsEmpty()
}

func main(){
	arrPush := []int{1,2,3,4,5}
	arrPop := []int{4,5,3,2,1}
	fmt.Println(" Is this Pop Order",IsPopOrder(arrPush,arrPop))

}