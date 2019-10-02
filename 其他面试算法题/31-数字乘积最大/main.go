/*
给出一个数字n。在1到n之间找到一个数字m，使得m的各个数位的数字乘积最大
 */
package main

import "fmt"

func findMaxM(n int)int{
	if n  <=0 {
		return 0
	}
	if n>0 && n <=10 {
		return n
	}
	max :=0
	for i:=11;i<=n;i++{
		s := calMul(i)
		if s > max {
			max = s
		}
	}
	return max
}

func findMax(n int)int{
	if n  <=0 {
		return 0
	}
	if n>0 && n <=10 {
		return n
	}
}

func calMul(n int)int{
	mul := 1
	for n!= 0{
		mul *= n%10
		n  /= 10
	}
	return mul
}

func main(){
	//fmt.Println(calMul(23))
	fmt.Println(findMaxM(100))
}