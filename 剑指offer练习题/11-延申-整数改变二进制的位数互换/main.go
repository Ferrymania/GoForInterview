/*
输入两个整数m和n，计算需要改变m的二进制表示中的多少位才能得到n
例如：10-》1010
13->1101 需要改变三位才能得到
 */
package main

import "fmt"

/*
第一步求这两个数的异或
第二步统计异或结果中的1的位数
 */

func exchange(m ,n int)int{
	res := m^n
	count :=0
	for res!=0 {
		count++
		res = res&(res-1)
	}
	return count
}

func main(){
	fmt.Println(exchange(10,13))
}
