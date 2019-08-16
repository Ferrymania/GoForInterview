/*
我们可以用2*1的小矩形横着或者竖着去覆盖更大的矩形。请问用n个2*1的小矩形无重叠地覆盖一个2*n的大矩形，总共有多少种方法？
 */
package main

import "fmt"

//覆盖方法记为f（n） ,仍然为斐波那契数列

func rectCover(number int) int {
	if number <1 {
		return 0
	}

	if number < 2 {
		return 1
	}
	return rectCover(number-1) + rectCover(number-2)
}

func main () {
	fmt.Println("10 2*1 rect cover solution ",rectCover(10))
}