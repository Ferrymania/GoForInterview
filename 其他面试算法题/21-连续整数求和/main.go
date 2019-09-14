/*
给定一个正整数N,试求有多少组连续 正整数满足所有数字之和为N(1<=N<=10^9)

example
5
2组 {5}  {2}{3}
 */
package main

import "fmt"

/*

 */
func sequence(N int)int{
	count := 1
	pHigh := 2
	pLow :=1
	for pHigh > pLow {
		cur := (pLow+pHigh)*(pHigh-pLow+1)/2
		if cur < N {
			pHigh++
		}
		if cur == N {
			count ++
			pLow ++
		}
		if cur>N{
			pLow++
		}
	}
	return count
}

func main(){
	var N int
	fmt.Scanln(&N)
	fmt.Println(sequence(N))
}