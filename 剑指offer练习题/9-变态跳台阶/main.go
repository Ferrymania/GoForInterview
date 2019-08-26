/*
一只青蛙一次可以跳上1级台阶，也可以跳上2级……它也可以跳上n级。求该青蛙跳上一个n级的台阶总共有多少种跳法。
 */
package main

import "fmt"

//f(n)记为跳法，f(1) = 1,f(2) = 2,f(3)=f(2)+f(1),f(n)=f(n-1)+f(n-2)+..+f(1)
//f(n) = pow(2,(n-1))
//或者f(n)=2*f(n-1)
func jumpFloor2(n int) int {
	if n <=2 {
		return 2
	}
	fn2 :=2
	for i:=3;i<=n;i++ {
		fn2 = fn2*2
	}
	return fn2
}

func main() {
	steps := jumpFloor2(5)
	fmt.Println("5 floors steps are ",steps)
}