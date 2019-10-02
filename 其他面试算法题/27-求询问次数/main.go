/*
Alice现在有一个数x，在范围[0,2n)内。你需要询问Alice一些问题。每次询问有一个数t，Alice会回答你
t&x 是否等于t。你不能根据已有的回答改变接下来的询问（即询问需要提前想好）。你需要最少询问多少次
来保证你能确定x的值。输出答案mod106+3
&的意思是位与运算
当且仅当两种方案中有一个询问不同，我们就认为两种方法是不同的。例如询问[1,2]和[1,3]是不同的，
[1,2][2,1]也是不同的（尽管他们得到了相同的结果）

样例输入
3
样例输出
6

提示
询问的t为1,2,4三个数，有6种排列方式（[1,2,4][1,4,2][2,1,4][2,4,1][4,1,2][4,2,1]）
 */
package main

import "fmt"

func minTimes(n int)int{
	count :=1
	for i:=1;i<n;i*=2{
		count++
	}
	fmt.Println(count)
	methods := 1
	for count >=1{
		methods *=count
		count--
	}
	return methods
}

func main(){
	var n int
	//fmt.Scanln(&n)
	n = 3
	fmt.Println(minTimes(n))
}