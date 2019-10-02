/*

第一行包含一个正整数，代表街上总共的商铺数量
第二行包含一个字符串，由N个字符组成，每个字符为'o',代表有厕所，要么为'x',代表无厕所
输入
9
XXOXOOXXX
输出
2 1 0 1 0 0 1 2 3
 */
package main

import "fmt"

func minDistance(N int,s string){
	res := make([]int,N)
	distance := N+1
	for i:=0;i<N;i++{
		if s[i] == 'O'{
			res[i] =0
			distance =1
		}else {
			if res[i]<distance {
				res[i] = distance
				distance +=1
			}
		}
	}
	distance = N+1
	for i:=N-1;i>=0;i--{
		if s[i] == 'O'{
			distance = 1
		}else {
			if res[i]>distance {
				res[i] = distance
				distance +=1
			}
		}
	}
	for i:=0;i<N;i++{
		fmt.Print(res[i]," ")
	}
}
func main(){
	//N :=9
	//str:="XXOXOOXXX"
	var N int
	var str string
	fmt.Scanln(&N)
	fmt.Scanln(&str)
	minDistance(N,str)
}
