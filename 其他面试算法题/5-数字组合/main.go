/*
有1,2,3,4个数，能组成多少个不相同无重复的三位数，分别是多少？


*/
package main

import "fmt"

func combination(arr []int)int{
	count := 0
	for i:=0;i<len(arr);i++{
		x:=arr[i]
		for j:=0;j<len(arr);j++{
			y:=arr[j]
			for k:=0;k<len(arr);k++{
				z:=arr[k]
				if x!=y && x!= z &&y!=z{
					count++
					fmt.Println(x*100+y*10+z)
				}
			}
		}
	}
	return count
}

func main(){
	arr:=[]int{1,2,3,4}
	fmt.Println(combination(arr))
}