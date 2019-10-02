/*
第一行输入两个整数N和M，N表示植树点数，M表示小伙伴的数量
第二行共M行，每行有两个整数Li和Ri,表示第i个小伙伴负责种树的区间
输出 共M行 每行一个整数，表示第i个小伙伴种树后，路上一共有多少棵树
input
4 3
1 2
3 3
4 4
output
2
3
4
 */
package main

import "fmt"

func calTreesum(n int,m int,arr [][]int){
	flag := make(map[int]bool,n)
	for i:=1;i<=n;i++{
		flag[i] = false
	}
	fmt.Println(flag)

	for i:=0;i<m;i++{
		for j:=arr[i][0];j<=arr[i][1];j++{
			flag[j] = true
		}
		fmt.Println(calSum(flag))
	}
}

func calSum(flag map[int]bool)int{
	s := 0
	for _,v := range flag{
		if v ==true {
			s++
		}
	}
	return s
}

func main(){
	var n,m int
	fmt.Scanln(&n,&m)
	arr :=make([][]int,m)
	for i:=0;i<m;i++{
		arr[i] = make([]int,2)
	}
	for i:=0;i<m;i++{
		fmt.Scanln(&arr[i][0],&arr[i][1])
	}
	//n:=4
	//m :=3
	//arr := [][]int{
	//	{1,2},
	//	{3,3},
	//	{4,4},
	//}
	calTreesum(n,m,arr)
}
