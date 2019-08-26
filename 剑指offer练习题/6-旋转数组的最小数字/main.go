/*
把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。输入一个递增排序的数组的
一个旋转，输出旋转数组的最小元素，例如数组{3,4,5,1,2}为{1,2,3,4,5}的一个旋转，该数组的
最小值为1
*/
package main

import "fmt"

/*
采用二分查找法，存在三种情况
arr[mid] > arr[high]:此时最小数字在mid的右边
arr[mid] == arr[high]:如[1,0,1,1,1],[0,1,1,1,1]等情况
arr[mid]<arr[high]:此时最小数字在mid的左边
 */
func getMinInRotateArray(arr []int)int{
	length := len(arr)
	if length == 0 {
		return 0
	}else if length == 1 {
		return arr[0]
	}
	low :=0
	high := length -1
	for low < high {
		mid := low + (high-low)/2
		if arr[mid] > arr[high]{
			low = mid +1
		}else if arr[mid] == arr[high]{
			high--
		}else {
				high = mid
		}
	}
	return arr[low]
}

func main() {
	arr := []int{3,4,5,1,2}
	min := getMinInRotateArray(arr)
	fmt.Println("min number is ",min)
}
