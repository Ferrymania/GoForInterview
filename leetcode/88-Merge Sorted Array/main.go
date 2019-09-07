/*
Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.
Note:
The number of elements initialized in nums1 and nums2 are m and n respectively.
You may assume that nums1 has enough space (size that is greater or equal to m + n) to hold additional elements from nums2.
Example:
Input:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3
Output: [1,2,2,3,5,6]
 */
package main

import "fmt"

//从前往后比较费时，利用从后往前，先申请空间
func merge(nums1 []int, m int, nums2 []int, n int)  {
	if m <=0 || n<=0{
		return
	}
	length := m+n
	for i:=m;i<length;i++{
		nums1 = append(nums1,0)
	}
	fmt.Println(nums1)
	//定义两个指针，分别指向nums1和nums2的末尾元素
	i := length-1
	j := m-1
	k := n-1
	for j>=0&&k>=0 {
		if nums1[j] > nums2[k]{
			nums1[i] = nums1[j]
			j--
			i--
		}else {
				nums1[i] = nums2[k]
				k--
				i--
		}
	}
	//如果第二个数组还有则继续，若是第一个数组有则保持原样即可
	if k>=0 {
		for k>=0 {
			nums1[i] = nums2[k]
			k--
			i--
		}
	}
	fmt.Println(nums1)
}

func main(){
	nums1 := []int{1,2,3}
	nums2 := []int{2,5,6}
	merge(nums1,3,nums2,3)
	fmt.Println(nums1)
}