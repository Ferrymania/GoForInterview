/*
输入一个整数，输出该数二进制表示中1的个数。其中负数用补码表示。
 */
package main

import "fmt"

//方法一：把整数和1做位与运算看结果是不是0,接着把输入的整数右移一位（右移一位等价于整数除以2，但除法效率低得多）
func NumberOf1Low(n int) int {
	count:=0
	flag :=1
	/*
	n 为负数时会陷入死循环
	for n!=0 {
			if (n & flag)!= 0{
				count ++
			}
			n = n >> 1
	}
	 */


	//把1不断左移做位与运算
	for flag!=0 {
			if (n & flag)!= 0{
				count ++
			}
			flag = flag << 1
		}

	return count
}

//方法二 把一个整数减去1之后再和原理的整数做位与运算，
// 得到的结果相当于是把整数的二进制表示中的最右边的一个1变为0
func NumberOf1(n int) int {
	count :=0
	for n!=0 {
		count++
		n = (n-1)&n
	}
	return count
}
func main() {
	fmt.Println("1 in 10 is ",NumberOf1Low(17))
	fmt.Println("high solution 1 in 10 is ",NumberOf1(17))
}