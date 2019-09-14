/*
给定一个长度为N-1且只包含0和1的序列A1到AN-1，如果一个1到N的排列P1到PN满足对于1≤i<N，当Ai=0时Pi<Pi+1，当Ai=1时Pi>Pi+1，则称该排列符合要求，那么有多少个符合要求的排列？
 */
package main

import "fmt"

const MOD = 1000000007
func maxSeq(N int,arr []int)int{
	dp :=make([][]int,N)
	for i:=0;i<N;i++{
		dp[i] = make([]int,N)
	}
	dp[0][0] = 1
	for i:=0;i<N-1;i++{
		if arr[i] == 0{
			for j:=1;j<=i+1;j++{
				dp[i+1][j] = (dp[i+1][j-1]+dp[i][j-1])%MOD
			}
		}else {
			for j:=i;j>=0;j--{
				dp[i+1][j]=(dp[i+1][j+1]+dp[i][j])%MOD
			}
		}
	}
	result := 0
	for i:=0;i<len(dp);i++{
		result = (result+dp[len(dp)-1][i])%MOD
	}
	return result
}

func main(){
	var N int
	fmt.Scanln(&N)
	arr :=make([]int,N-1)
	var a int
	for i:=0;i<N-1;i++{
		fmt.Scan(&a)
		arr[i] = a
	}
	fmt.Println(N)
	fmt.Println(arr)
	res := maxSeq(N,arr)
	fmt.Println(res)
}