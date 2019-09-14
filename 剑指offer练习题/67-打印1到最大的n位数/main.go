/*
输入数字n，按顺序打印出从1到最大的n位十进制数。比如输入3，则打印出1，2，3
一直到最大的3位数999
 */
package main

import "fmt"

func PrintToMaxDigits(n int){
		max := 1
		for i:=0;i<n;i++{
			max *=10
		}
		max--
		for i:=1;i<=max;i++{
			fmt.Println(i)
		}
}

func main(){
	PrintToMaxDigits(10000000000000)
}