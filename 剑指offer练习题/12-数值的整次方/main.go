/*
给定一个double类型的浮点数base和int类型的整数exponent。求base的exponent次方。
 */
package main

import "fmt"

//判断两个浮点数是否近似相等
func equal(num1,num2 float64) bool {
		if (num1-num2 > -0.00000001)&&(num1-num2 <0.0000001){
			return true
		}else {
			return false
		}
}

func Power(base float64,exponent int) float64 {
	if(equal(base,0.0) && exponent < 0){
		return 0.0
	}
	absExponent := exponent
	if exponent < 0{
		absExponent = - exponent
	}
	result := PowerWithPositiveExponent(base,absExponent)
	if exponent<0 {
		result = 1.0 /result
	}

	return  result
}

func PowerWithPositiveExponent(base float64,exponent int) float64{
	result := 1.0
	for i:=1;i<=exponent;i++ {
		result *= base
	}
	return result
}
func main() {
	fmt.Println("pow(2.0,-2)",Power(2,-2))
}

