/*
要求输入一个整数n,求出斐波那契数列的第n项
 */
package main

import "fmt"

//利用传统的递归方法进行计算，效率低下
func Fibonacci(n int) int{
	if n<=0 {
		return 0
	}
	if n==1 {
		return 1
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}

//从下往上计算，不重复计算数列中间项 时间复杂度O(n)
func Fibonacci2(n int)int {
	arr := [2]int{0,1}
	if n <2{
		return arr[n]
	}
	fibNMinusOne :=1
	fibNMinusTwo :=0
	fibN := 0

	for i :=2 ;i<=n;i++{
		fibN = fibNMinusOne + fibNMinusTwo
		fibNMinusTwo = fibNMinusOne
		fibNMinusOne = fibN
	}
	return fibN
}

//利用数学公式 线性代数解法 时间复杂度O(logn)
//   |f(n)   f(n-1) | = |1 1|(n-1)
//   |f(n-1) f(n-2) |   |1 0|
func main(){
	fmt.Println("first solution is ",Fibonacci(10))

	fmt.Println("second solution is ",Fibonacci2(10))
}
