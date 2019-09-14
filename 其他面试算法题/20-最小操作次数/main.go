/*
给出两个单词word1和word2,计算出将word1转换为word2的最少操作次数
总共有三种操作方法
插入一个字符
删除一个字符
替换一个字符
 */
package main

import "fmt"

func minOperation(str1,str2 string)int{
	len1 := len(str1)
	len2 := len(str2)
	dp := make([][]int,len1+1)
	for i:=0;i<len1+1;i++{
		dp[i] = make([]int,len2+1)
	}

	for i:=0;i<=len1;i++{
		for j:=0;j<=len2;j++{
			if i==0 {
				dp[i][j] = j
			}else if j==0 {
				dp[i][j] = i
			}else if str1[i-1] ==str2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			}else {
				dp[i][j] = 1 + min(dp[i][j-1],min(dp[i-1][j],dp[i-1][j-1]))
			}
		}
	}
	return dp[len1][len2]
}

func min(a,b int)int{
	if a<b {
		return a
	}
	return b
}

func main(){
	var str1 string
	var str2 string
	fmt.Scanln(&str1)
	fmt.Scanln(&str2)
	//str1 := "abc"
	//str2 := "abd"
	fmt.Println(minOperation(str1,str2))
}