/*
刚开始有一行不同颜色的方块，每次可以消掉相邻的、颜色相同的若干个（k 个），并获得 k*k 分。
现在给出一个游戏的起始局面,直到所有的豆子都消掉为止，问最多能获得多少分？
输入一行n个正整数，代表这一行中豆子的颜色及排列
输入 1 4 2 2 3 3 2 4 1
输出 21
 */
package main

import "fmt"


func dfs(arr []int,dp [][][]int,l,r,x int)int{
	if dp[l][r][x] != 0 {
		return dp[l][r][x]
	}
	if l == r{
		dp[l][r][x] = (x+1)*(x+1)
		return dp[l][r][x]
	}else if l<r{
		dp[l][r][x] = dfs(arr,dp,l,r-1,0)+(x+1)*(x+1)
		for i :=l;i<r;i++{
			if arr[i] == arr[r]{
				dp[l][r][x] = max(dp[l][r][x],dfs(arr,dp,i+1,r-1,0)+dfs(arr,dp,l,i,x+1))
			}
		}
		return dp[l][r][x]
	}
	return 0
}

func max(a,b int)int {
	if a> b{
		return a
	}
	return b
}

func main(){
	arr :=[]int{1,4,2,2,3,3,2,4,1}
	//arr :=[]int{1,2,2,1,1}
	length := len(arr)
	dp :=make([][][]int,length)
	for i:=0;i<length;i++{
		dp[i] = make([][]int,length)
	}
	for i:=0;i<length;i++{
		for j:=0;j<length;j++{
			dp[i][j] = make([]int,length)
		}
	}
	fmt.Println(dp)
	fmt.Println(dfs(arr,dp,0,len(arr)-1,0))
}