/*
输入一个递增排序的数组和一个数字s，在数组中查找两个数，使得它们的和正好是s。
如果有多对数字的和等于s，则输出任意一对即可
 */
package main

import "fmt"

/*
双指针法
第一个指针指向数组第一个元素
第二个指针指向数组最后一个元素
如果两个指向的数之和大于s，则尾指针向前进
如果两数之和小于s，则头指针向后移
*/

func FindNumbersWithSum(data []int,s int)bool{
	if data == nil || len(data)<2 {
		return false
	}
	length := len(data)
	head :=0
	end := length-1
	for head < end {
		if data[head] + data[end] == s{
			fmt.Printf("%d + %d :=s",data[head],data[end])
			return true
		}else if data[head] + data[end] > s{
			end--
		}else {
			head ++
		}

	}
	return false
}

func main(){
	arr := []int{1,2,4,7,11,15}
	FindNumbersWithSum(arr,15)
}