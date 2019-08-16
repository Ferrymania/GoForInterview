/*
Given an unsorted array of integers, find the length of the longest consecutive elements sequence.
Your algorithm should run in O(n) complexity.
example:
Input: [100, 4, 200, 1, 3, 2]
Output: 4
Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
 */
package main

import "fmt"

//Brute Force
//Time complexity O(n^3) space complexity O(1)
func arrayContains(arr []int,num int)bool{
	for _,v := range arr{
		if v == num {
			return true
		}
	}
	return false
}
func longestConsecutive(nums []int) int {
	longestStreak := 0
	for _,v := range nums{
		currentNum := v
		currentStreak := 1

		for arrayContains(nums,currentNum+1){
			currentNum +=1
			currentStreak +=1
		}
		if currentStreak > longestStreak {
			longestStreak = currentStreak
		}
	}
	return longestStreak
}

//HashSet and Intelligent Sequence Building
// Time complexity O(n) Space complexity O(n)
func longestConsecutive2(nums []int)int {
	intSet := make(map[int]bool)
	for _,v :=range nums {
		intSet[v] = true
	}
	longestStreak :=0
	for k := range intSet{
		if !intSet[k-1]{
			currentNum := k
			currentStreak :=1
			for intSet[currentNum+1]{
				currentNum++
				currentStreak++
			}
			if currentStreak > longestStreak {
				longestStreak = currentStreak
			}
		}
	}
	return longestStreak
}
func main(){
	a :=[]int{100,1,200,2,4,3}
	fmt.Println(longestConsecutive(a))
	fmt.Println(longestConsecutive2(a))
}