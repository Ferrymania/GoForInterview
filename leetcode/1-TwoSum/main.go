/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
 */
package main

import "fmt"

//Brute Force
// Time complexity O(n^2)  Space complexity O(1)
func twoSum(nums []int, target int) []int {
	for i :=0;i<len(nums)-1;i++ {
		for j:=i+1;j<len(nums);j++{
			if nums[i] + nums[j] ==target{
				return []int{i,j}
			}
		}
	}
	return nil
}

//Two-pass Hash Table
//Time complexity O(n)  Space complexity O(n)
func twoSum2(nums []int, target int) []int {
	m :=make(map[int]int)
	for i,v := range nums  {
		m[v] = i
	}
	for i :=0;i<len(nums)-1;i++ {
		complement := target - nums[i]
		if _,containsKey := m[complement];containsKey == true && m[complement] != i{
			return []int{i,m[complement]}
		}
	}
	return nil
}
//One-pass Hash Table
func twoSum3(nums []int,target int)[]int {
	m :=make(map[int]int)
	for i:=0;i<len(nums);i++ {
		complement := target -nums[i]
		if _,containsKey := m[complement];containsKey == true && m[complement] != i{
			return []int{i,m[complement]}
		}
		m[complement] = i
	}
	return nil
}

func main() {
	arr := []int{2,7,3,4,6,5,9}
	fmt.Println(twoSum(arr,12))
	fmt.Println(twoSum2(arr,12))
	fmt.Println(twoSum3(arr,12))
}