/*
小明想知道 n! ，最后一位不为 0 的数字，你能告诉他吗？ n! = n*(n-1)*(n-2)*....*3*2*1
样例输入
7
样例输出
4
提示
7！=5040 ，最后一位不为0的数字为4
*/
package main

import "fmt"

func postion(n int)int{
	if n <= 0{
		return 0
	}
	f :=1
	for i:=1;i<=n;i++{
		f *=i
		for f%10==0 {
			f /=10
		}
		f = f%1000000
	}
	return f%10
}

func main(){
	fmt.Println(postion(7))
}