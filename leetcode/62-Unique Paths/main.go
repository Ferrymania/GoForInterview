/*
A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).
The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).
How many possible unique paths are there?
Input: m = 3, n = 2
Output: 3
Explanation:
| start|  |   |
|      |  |end|
From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
1. Right -> Right -> Down
2. Right -> Down -> Right
3. Down -> Right -> Right
 */
package main

import "fmt"

//设opt[i][j]为到位置[i,j]路径数，opt[i][j] = opt[i-1][j] + opt[i,j-1]
//此外，opt[0][j]=opt[i][0]=1
func uniquePaths(m int, n int) int {
	opt :=make([][]int,m)
	for i:=range opt{
		opt[i] = make([]int,n)
	}
	for i:=range opt {
		for j:=range opt[i]{
			opt[i][j] = 1
		}
	}
	for i:=1;i<m;i++ {
		for j:=1;j<n;j++{
			opt[i][j] = opt[i-1][j] + opt[i][j-1]
		}
	}
	return opt[m-1][n-1]
}

func main(){
	fmt.Println("3 row  2 cols solution is ",uniquePaths(3,2))
}
