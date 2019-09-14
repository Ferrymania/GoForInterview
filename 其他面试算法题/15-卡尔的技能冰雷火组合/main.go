/*
dota游戏里面，召唤师可以控制冰雷火三种元素，并通过元素组合产生新的技能。现在我们
修改了一张新的地图，地图中他能够控制n种元素，并且将m个元素围成一个圈组成一个新的
技能（）
 */
package main

import "fmt"

const p = 1000000007
func fac(n int)int{
	if n == 0 {
		return 1
	}
	f := 1
	for i:=1;i<=n;i++ {
		f *=i
	}
	return f
}
func combineSum(n,m int)int {
	sum := fac(n+m-1)/(fac(n-1)*fac(m))
	return sum
}

func pow(a,b int)int {
	res := 1
	for b!=0{
		if b&1 != 0 {
			res = (res%p)*(a%p)%p
		}
		a  = (a%p)*(a%p)%p
		b = b>>1
	}
	return  res
}
func C(a,b int)int {
	if a < b {
		return 0
	}
	res:=1
	for i:=1;i<=b;i++{
		j:=(a-b+i)%p
		k :=i%p
		res = res * (j*pow(k,p-2)%p)%p
	}
	return res%p
}
func combine(a,b int) int {
	if b == 0 {
		return 1
	}
	return (C(a%p,b%p)*combine(a/p,b/p))%p
}
func main(){

	n:=0
	m:=0

	fmt.Scanln(&n,&m)

	fmt.Println(fac(5))
	fmt.Println(combineSum(n,m))
	mod := combineSum(n,m)%1000000007
	fmt.Println(mod)
	fmt.Println(combine(n+m-1,m))
}