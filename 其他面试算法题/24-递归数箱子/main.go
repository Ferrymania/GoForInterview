/*
第一个箱子比第二个多2个 第二个比第三个多3个，第三个比第四个多7个，第4个比第5个多5个
第五个比第六个多121个
输入箱子数，以及最后一个箱子的物品数
求所有箱子物品总数
 */
package main

func sum(M int,min int)int{
	s := 0
	for i:=1;i<=M;i++{
		s += f(i,M,min)
	}
	return s
}

func f(n int ,M int,min int) int{
	if n == M {
		return min
	}
	return f(n+1,M,min) +  2
}

func main(){

}