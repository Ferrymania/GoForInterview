/*
求1+2+....+n 要求不能使用乘除法，for while if else switch case等guan'ni
 */
package main

import "fmt"

func cal(n int) int{
	if n  <=1 {
		return 1
	}
	return n + cal(n-1)
}

func main (){
	fmt.Println(cal(3))
}