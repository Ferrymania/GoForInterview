/*
 robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).
The robot can only move either down or right at any point in time. The robot is
trying to reach the bottom-right corner of the grid (marked 'Finish' in the
diagram below).
Now consider if some obstacles are added to the grids. How many unique paths
would there be?

An obstacle and empty space is marked as 1 and 0 respectively in the grid.
Note: m and n will be at most 100.
Input:
[
  [0,0,0],
  [0,1,0],
  [0,0,0]
]
Output: 2
Explanation:
There is one obstacle in the middle of the 3x3 grid above.
There are two ways to reach the bottom-right corner:
1. Right -> Right -> Down -> Down
2. Down -> Down -> Right -> Right
 */
package main

/*
假设opt[i][j] 为 到点（i，j）的路径数，则
opt[i][j] = opt[i-1][j] + opt[i][j-1]
如果该位置是障碍的话，则
opt[i][j] = 0
 */
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid == nil {
		return 0
	}
	// no paths to the destination.
	if (obstacleGrid[0][0] == 1) {
		return 0;
	}

	m := len(obstacleGrid)  // rows
	n := len(obstacleGrid[0])  //cols

	opt := make([][]int,m)
	for i:=0;i<m;i++{
		opt[i] = make([]int,n)
	}

	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			if obstacleGrid[i][j] != 1{
				opt[i][j]
			}
		}
	}

}