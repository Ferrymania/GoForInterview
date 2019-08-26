package main

import "fmt"

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

func main() {
	arr :=[]int{5,4,9,8,7,6,0,1,3,2}
	fmt.Println(quickSort(arr))
}

