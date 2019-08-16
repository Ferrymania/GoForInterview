package main

import "fmt"

func bubbleSort(arr []int){
	var temp int
	for i:=0;i<len(arr);i++{
		for j:=len(arr)-1;j>i;j-- {
			fmt.Printf("j is %v a[j] is %v \n",j,arr[j])
			if arr[j]<arr[j-1]{
				temp = arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = temp
			}
		}
	}
}

func main() {
	arr :=[]int{38,65,97,76,13,27,49}
	bubbleSort(arr)
	fmt.Println("after sort",arr)
}
