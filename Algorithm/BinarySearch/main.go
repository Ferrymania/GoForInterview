package main

import "fmt"

func binarySearch(list []int,target int) bool {
	if list == nil {
		return false
	}
	low := 0
	high := len(list)-1
	for low <= high {
		mid := (low+high)/2
		fmt.Println("low is ",low," high is ",high," mid is ",mid)
		if list[mid] == target {
			return true
		}
		if list[mid] < target {
			low = mid+1
		}else{
			high = mid-1
		}
	}
	return false
}

func main() {
	fmt.Println(binarySearch([]int{2,5,7,8,9,11,13},8))
	fmt.Println(binarySearch([]int{2,4,6,7,8,9,10},1))
}
