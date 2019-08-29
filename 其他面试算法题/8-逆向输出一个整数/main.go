/*
给定一个整数，逆向输出它
 */
package main

import "fmt"

func ReverseNum(num int)int{
	var arr []int

	for num>0{
		arr = append(arr,num%10)
		num /=10
	}
	n := 0
	j := 1
	for i:=len(arr)-1;i>=0;i--{
		n += arr[i]*j
		j *=10
	}
	return n
}

func main(){
	fmt.Println(" 1320 reverse is ",ReverseNum(1320))
}