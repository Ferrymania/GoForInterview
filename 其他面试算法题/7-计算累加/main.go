/*
输入n,如果是奇数就1/2+1/4+...+1/n,如果是偶数就1/1+1/3+...+1/n
 */
package main

import (
	"fmt"
	"math"
)

func cal(num int)float64{
	if num == 0{
		return 0.0000000000000000
	}
	var sum float64
	if num %2 ==0{
		for i:=2;i<=num;i*=2{
			sum += 1.0/float64(i)
		}
	}else {
		for i:=1;i<=num;i=2*i+1{
			sum +=1.0/float64(i)
		}
	}
	return Decimal(sum)
}


func Decimal(value float64) float64 {
	//value, _ = strconv.ParseFloat(fmt.Sprintf("%.3f", value), 64)
	//return value
	return math.Trunc(value*1e16+0.5) * 1e-16
}
func main(){
	fmt.Printf("cal(4) : %.16f",cal(1))
}


