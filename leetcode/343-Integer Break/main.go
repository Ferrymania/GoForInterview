/*
Given a positive integer n, break it into the sum of at least two positive integers and maximize the product of those integers. Return the maximum product you can get.
Input: 10
Output: 36
Explanation: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36.
 */
package main

import "fmt"

// 递归+记忆化搜索
var memo []int
func integerBreak(n int) int {

	return breakInteger(n)
}

func breakInteger(n int)int{
	if n== 1{
		return 1
	}
	if memo[n] != 0{
		return memo[n]
	}
	res := -1
	for i:=1;i<=n-1;i++{
		res = max(res,i*(n-i),i*breakInteger(n-i))
	}
	memo[n] = res
	return res
}

func max(a,b,c int)int{
	max := a
	if b>max {
		max=b
	}
	if c>max{
		max = c
	}
	return max
}


func main(){
	fmt.Println("10 break max is ",integerBreak(10))
}
