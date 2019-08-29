package main

import "fmt"

func selectSort(arr []int){
	var min int
	for i:=0;i<len(arr);i++{
		min = i
		for j:=i+1;j<len(arr);j++{
			if arr[j]<arr[min] {
				min = j
			}
		}
		if(min != i){
			temp := arr[min]
			arr[min] = arr[i]
			arr[i] = temp
		}
	}
}

func main() {
	//arr :=[]int{38,65,97,76,13,27,49}
	arr := []int{1, 4, 2, 6, 3}
	selectSort(arr)
	fmt.Println("after select",arr)
}
