/*
Given a triangle, find the minimum path sum from top to bottom. Each step you may move to adjacent numbers on the row below.
For example, given the following triangle
[
     [2],
    [3,4],
   [6,5,7],
  [4,1,8,3]
]
The minimum path sum from top to bottom is 11 (i.e., 2 + 3 + 5 + 1 = 11).
 */
package main

import "fmt"

// Top-Down recursion
func minimumTotal_recursive(triangle [][]int) int {
	if triangle == nil || len(triangle) ==0 {
		return 0
	}
	return dfs(0,0,triangle)
}

func dfs(row int,pos int,triangle [][]int) int{
	if(row+1>len(triangle)){
		return triangle[row][pos]
	}
	return triangle[row][pos]+ min(dfs(row+1,pos,triangle),dfs(row+1,pos+1,triangle))

}

func min(a,b int)int{
	if a <=b {
		return a
	}else{
		return b
	}
}

func main() {
	triangle := [][]int{{2},
						{3,4},
						{6,5,7},
						{4,1,8,3}}
	fmt.Println(minimumTotal_recursive(triangle))
}