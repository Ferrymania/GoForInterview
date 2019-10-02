package main

import "fmt"

func bubbleSort(arr []int){
	var temp int
	flag := false
	for i:=0;i<len(arr);i++{
		for j:=len(arr)-1;j>i;j-- {
			if arr[j]<arr[j-1]{
				temp = arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = temp
				//如果没有交换 说明已经有序
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}

func main() {
	arr :=[]int{38,65,97,76,13,27,49}
	bubbleSort(arr)
	fmt.Println("after sort",arr)
}
