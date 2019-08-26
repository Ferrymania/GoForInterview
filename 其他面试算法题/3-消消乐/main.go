package main

import "fmt"

func main(){
	arr :=make([][]int,5)
	for i:= range arr{
		arr[i] = make([]int,5)
	}
	for i:=0;i<5;i++{
		for j:=0;j<5;j++{
			arr[i][j] = i+j
		}
	}
	fmt.Println(arr)
}