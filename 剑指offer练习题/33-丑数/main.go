/*
把只包含质因子2、3和5的数称作丑数（Ugly Number）。例如6、8都是丑数，但14不是，因为它包含质因子7。 习惯上我们把1当做是第一个丑数。求按从小到大的顺序的第N个丑数。
 */
package main

import "fmt"

func IsUgly(number int)bool{
	for number%2==0 {
		number /= 2
	}
	for number%3 == 0{
		number /=3
	}
	for number%5 == 0{
		number /=5
	}
	return number == 1
}
func GetUglyNumber(index int) int {
	if index <= 0 {
		return 0
	}
	number :=0
	uglyFound :=0
	for uglyFound<index {
		number++
		if IsUgly(number){
			uglyFound++
		}
	}
	return number
}
func main(){
	fmt.Println(GetUglyNumber(5))
}

