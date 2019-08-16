/*
用一条语句判断一个整数是不是2的整数次方。
 */
package main

import "fmt"

//一个整数如果是2的整数次方，那么它的二进制表示中有且只有一位i是1，而其他所有位都是0
func is2Pow(n int) bool {
	return n &(n-1) == 0
}

func main() {
	fmt.Println("Is 8 is 2 pow",is2Pow(8))
	fmt.Println("Is 15 is 2 pow",is2Pow(15))
}
