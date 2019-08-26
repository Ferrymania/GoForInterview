/*
You are climbing a stair case. It takes n steps to reach to the top.
Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
Input: 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step
 */
package main

// recursive,long time cost
func climbStairs_recursive(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2{
		return 2
	}
	return climbStairs_recursive(n-1) + climbStairs_recursive(n-2)
}


//动态规划，记忆化搜索
func climbStairs(n int) int {
	memo :=make([]int,n+1)
	memo[0] = 1
	memo[1] = 1
	for i:=2;i<=n;i++{
		memo[i] = memo[i-1]+memo[i-2]
	}
	return memo[n]
}