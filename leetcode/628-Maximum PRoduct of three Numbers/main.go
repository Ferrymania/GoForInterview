/*
Given an integer array, find three numbers whose product is maximum and output the maximum product.
Input: [1,2,3]
Output: 6

Input: [1,2,3,4]
Output: 24

The length of the given array will be in range [3,104] and all elements are in the range [-1000, 1000].
Multiplication of any three numbers in the input won't exceed the range of 32-bit signed integer.
*/
package main

import "fmt"

/*
brute force 遍历每个数  O(n^3)

第二种方法
先对数组排序，最大值的情况要么是三个最大的正整数 要么是两个最小的负数和一个最大的正整数
 */
func maximumProduct(nums []int) int {
	length := len(nums)
	if len(nums)<3{
		return 0
	}
	arr :=quickSort(nums)
	maxProject := max(arr[0]*arr[1]*arr[length-1],arr[length-3]*arr[length-2]*arr[length-1])
	return maxProject
}

func quickSort(arr []int)[]int{
	//基线条件，数组为空或只包含一个元素的数组是有序的
	if len(arr) < 2{
		return arr
	}
	pivot := arr[0]  //选定首项为基准值
	less := []int{}  //小于基准值的数字组成的数组
	greater :=[]int{} //大于基准值的数字组成的数组
	for _,num :=range arr[1:]{
		if pivot > num{
			less = append(less,num)   //比基准值小，移向左边的数组
		}else {
			greater = append(greater,num)  //比基准值大，移向右边的数组
		}
	}
	less = append(quickSort(less),pivot)  //再对左边的子数组进行排序
	greater = quickSort(greater)  //再对右边的子数组进行排序
	return append(less,greater...)
}

func max(a,b int)int{
	if a > b{
		return a
	}
	return b
}

func main(){
	arr := []int{-3,-4,-6,1,2,3,4}
	fmt.Println(maximumProduct(arr))
}