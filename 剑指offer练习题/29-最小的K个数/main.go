/*
输入n个整数，找出其中最小的k个数。例如，4 5 1 6 2 7 3 8
最小的4个数字是 1 2 3 4
 */
package main

import "fmt"

/*
方法1 基于排序算法找出钱n个 时间复杂度O(nlogn)
 */

/*
方法2 基于partion。如果基于数组的第k个数字来调整，则使得比第k个数字小的所有数字都位于数组
左边，比第k个数字大的所有数字都为数组的右边。这样调整之后，位于数组中左边的k个数字就是最小的k个数（不一定有序）
 */

func GetLeastKNumber(arr []int,k int)[]int{
	length := len(arr)
	if length < k {
		return nil
	}
	start :=0
	end := length-1
	index :=Partition(arr,start,end)
	for index!=k-1{
		if index>k-1{
			end = index-1
			index = Partition(arr,start,end)
		}else{
			start = index +1
			index = Partition(arr,start,end)
		}
	}
	return arr[:k]
}

func Partition(arr []int,start,end int)int{
	low:=start
	high := end
	pivot := arr[low]
	for low<high{
		for low<high && pivot<=arr[high]{
			high--
		}
		arr[low] = arr[high]
		for low<high && pivot>=arr[low]{
			low++
		}
		arr[high] = arr[low]
	}
	arr[low]=pivot
	return low
}


func main(){
	arr :=[]int{4,5,1,6,2,7,3,8}
	fmt.Println(GetLeastKNumber(arr,4))
}