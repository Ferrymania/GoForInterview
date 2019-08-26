/*
输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有的奇数位于数组的前半部分，所有的偶数位于数组的后半部分。
 */
package main

import "fmt"

func reOrderArray(arr []int){
	headIndex :=0
	tailIndex := headIndex+len(arr) -1
	if headIndex == tailIndex {
		return
	}
	var temp int
	for headIndex < tailIndex {
		for !isEven(arr[headIndex]) && headIndex < tailIndex {
			headIndex++
		}
		for isEven(arr[tailIndex]) && headIndex <tailIndex{
			tailIndex--
		}
		if headIndex < tailIndex{
			temp =  arr[headIndex]
			arr[headIndex] = arr[tailIndex]
			arr[tailIndex] = temp
		}
	}
}

func isEven(num int)bool {
	if num%2 == 0 {
		return true
	}
	return false
}

func main() {
	arr := []int{1,2,3,4,5}
	reOrderArray(arr)
	fmt.Println(arr)
}