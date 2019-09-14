/*
求两个整数之和，不能用+、-、*、/四则运算符号。
 */
package main

import "fmt"

/*
进行异或运算(两个值不相同，则为1)，计算两个数各个位置上的相加，不考虑进位；
进行位与运算，然后左移一位，计算进位值；
把异或运算的结果赋给 num1，把进位值赋给 num2，依此循环，进位值为空的时候结束循环
num1就是两数之和。
 */

func add(a,b int) int {
	if b == 0 {
		return a
	}
	sum := 0
	carry :=0
	for b!= 0 {
		sum = a ^ b
		carry =(a& b)<<1
		a = sum
		b = carry
	}
	return a
}

func main(){
	fmt.Println(add(3,4))
	fmt.Println(add(4,0))
}