/*
输入一个整型数组，数组里有整数也有负数。数组中的一个或连续多个整数组成一个子数组
求所有子数组的和的最大值。要求时间复杂度为O（n)
example
{1,-2,3,10,-4,7,2,-5}
最大和
为18 {3,10,-4,7,2}
 */
package main

import "fmt"

/*
对于一个数组中的一个数x，若是x的左边的数加起来非负，那么加上x能使得值变大，这样我们认为x之
前的数的和对整体和是有贡献的。如果前几项加起来是负数，则认为有害于总和。
我们用cur记录当前值, 用max记录最大值，如果cur<0,则舍弃之前的数，让cur等于当前的数字，否则，
cur = cur+当前的数字。若cur和大于max更新max。
 */

func FindGreatestSumOfSubArray(arr []int)int{
	if arr == nil {
		return 0
	}
	length :=  len(arr)
	cur :=arr[0]
	max := arr[0]
	for i:=1;i<length;i++{
		  if cur >= 0{
		  	cur = cur+ arr[i]
		  }else {
		  	cur = arr[i]
		  }
		  if max < cur {
		  	max = cur
		  }
	}
	return max
}

func main(){
	arr := []int{1,-2,3,10,-4,7,2,-5}
	fmt.Println(FindGreatestSumOfSubArray(arr))
}