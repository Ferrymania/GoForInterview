package main

import (
	"fmt"
)

func insertSort(arr []int){
	var  key ,j int
	for i:=1;i<len(arr);i++ {
		key = arr[i]
		j = i-1
		for j>=0 && arr[j] > key {
			arr[j+1]=arr[j]
			j=j-1
		}
		arr[j+1] = key
	}
}

func main() {
	arr :=[]int{38,65,97,76,13,27,49}
	insertSort(arr)
	fmt.Println("after select",arr)
}
