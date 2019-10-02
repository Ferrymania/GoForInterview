/*
给定N个数字，每次从中取走2个数，然后计算出它们的乘积，而这个过程重复M次，求
最后这M个乘积和的最小值
 */
package main

import "fmt"

func getMinSum(n int,m int,arr []int)int{
	if 2*m > n {
		return -1
	}
	sum :=0
	sortArr := quickSort(arr)
	head :=0
	end := 2*m-1
	for i:=0;i<m;i++{
		sum += sortArr[head]*sortArr[end]
		head++
		end--
	}
	return sum
}

func quickSort(arr []int)[]int{
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0]
	less := []int{}
	greater := []int{}

	for i:=1;i<len(arr);i++{
		if pivot > arr[i] {
			less =append(less,arr[i])
		}else{
			greater = append(greater,arr[i])
		}
	}

	less = append(quickSort(less),pivot)
	greater = quickSort(greater)
	return append(less,greater...)
}

func main(){
	n:=4
	m:=2
	arr := []int{4,3,2,1}
	fmt.Println(quickSort(arr))
	fmt.Println(getMinSum(n,m,arr))
}