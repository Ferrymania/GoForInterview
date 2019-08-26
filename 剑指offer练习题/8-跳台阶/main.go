/*
一只青蛙一次可以跳上1级台阶，也可以跳上2级。求该青蛙跳上一个n级的台阶总共有多少种跳法（先后次序不同算不同的结果）。
 */
package main

import "fmt"

//跳法记为f(n) f(1) = 1,f(2)=2,.....,f(n)=f(n-1)+f(n-2),即为斐波那契数列
func JumpFloor(n int) int {
	if (n<=2&& n>0){
		return n
	}
	return JumpFloor(n-1)+JumpFloor(n-2)
}

//递推
func JumpFloor2(n int) int {
	if n <=2 {
		return n
	}
	fn1 :=1
	fn2 :=2
	for i:=3;i<=n;i++ {
		fn2 = fn1 + fn2
		fn1 = fn2 - fn1
	}
	return fn2
}

func main() {
	fmt.Println("十级台阶的跳法为：",JumpFloor(10))
	fmt.Println("十级台阶的跳法为：",JumpFloor2(10))
}


