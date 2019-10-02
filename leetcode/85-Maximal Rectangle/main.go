/*
Given a 2D binary matrix filled with 0's and 1's, find the largest rectangle containing only 1's and return its area.
Input:
[
  ["1","0","1","0","0"],
  ["1","0","1","1","1"],
  ["1","1","1","1","1"],
  ["1","0","0","1","0"]
]
Output: 6
 */
package main

import "fmt"

func maximalRectangle(matrix [][]string)int{
	if len(matrix)<=0 {
		return 0
	}
	m := len(matrix)  //rows
	n := len(matrix[0])  //cols
	left := make([]int,n)
	right := make([]int,n)
	height :=make([]int,n)
	maxArea :=0
	for i:=0;i<m;i++{
		cur_left :=0
		cur_right :=n

		for j:=0;j<n;j++{
			if matrix[i][j] == "1"{
				height[j]++
			}else {
					height[j]=0
			}
		}

		for j:=0;j<n;j++{
			if matrix[i][j] == "1"{
				left[j]=max(left[j],cur_left)
			}else{
				left[j]=0
				cur_left=j+1
			}
		}

		for j:=n-1;j>=0;j--{
			if matrix[i][j]=="1" {
				right[j]=min(right[j],cur_right)
			}else{
				right[j]=n
				cur_right=j
			}
		}

		for j:=0;j<n;j++ {
			maxArea = max(maxArea,(right[j]-left[j])*height[j])
		}

	}
	return maxArea
}

func max(a,b int)int{
	if a >b{
		return a
	}
	return b
}

func min(a,b int)int{
	if a <b {
		return a
	}
	return b
}

func main(){
	matrix :=[][]string{
		{"1","0","1","0","0"},
		{"1","0","1","1","1"},
		{"1","1","1","1","1"},
		{"1","0","0","1","0"},
	}

	fmt.Println(maximalRectangle(matrix))
}