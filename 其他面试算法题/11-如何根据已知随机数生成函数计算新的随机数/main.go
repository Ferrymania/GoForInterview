/*
已知随机数生成函数rand7()能产生的随机数是整数1-7的均匀分布，
如何构造rand10()函数，使其产生的随机数是整数1-10的均匀分布
 */
package main

import (
	"fmt"
	"math"
	"math/rand"
)

/*
rand7() -> {1,2,3,4,5,6,7} 集合B的每个出现概率都是1/7
rand7()-1  -> {0,1,2,3,4,5,6} 集合A的每个出现概率都是1/7
[rand7()-1]*7  -> {0,7,14,21,28,35,42} 集合B的每个出现概率都是1/7
A B 独立 P(AB) = P(A) * P (B)
{rand7()} + {[rand7()-1]*7} -> {1,2,3,4,...,49}每个数出现的概率都是49
截断至{1,2,3,...40}
 */

func rand7() int {
	//rand.Seed(time.Now().UnixNano())
	return rand.Intn(7)+1
}

func rand10() int{
	a := math.MaxInt32
	for a > 40 {
		a = (rand7()-1)*7 + rand7() //{1,2,3..,49}
	}
	return a%10+1
}



func main() {
	for i:=0;i<200;i++{
		fmt.Println(rand10())
	}

}
