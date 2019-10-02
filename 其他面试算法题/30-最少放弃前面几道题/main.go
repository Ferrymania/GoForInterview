/*
一共有序号1~n的n道题目。小明从序号1的题目开始做，按照序号递增的顺序开始解答。每道题如果
解答需要消耗时间a[i]（如果剩余的时间不足a[i]则无法解答），如果跳过则不消耗时间
因为每道题都可能帮助到期末考试，小明想知道对于每一道题最少需要放弃前面几道题，才能解答完成
每个样例的第一行是n,m,表示n道题目和练习时长m秒
第二行是n个数字a[i],分别表示n道题的解答时耗为a[i]秒
输入
2
5 5
1 2 3 4 5
4 4
4 3 2 1
输出
0 0 1 2 4
0 1 2 2
 */
package main

import "fmt"

func minGiveUp(n int,m int,a []int)[]int{
	res := make([]int,n)
	res[0] = 0
	remain :=0
	for i:=1;i<n;i++{
		remain = m-a[i]
		if remain == 0{
			res[i] = i
			continue
		}
		prev :=[]int{}
		for j:=0;j<i;j++{
			prev = append(prev,a[j])
		}
		prevSort := sort(prev)
		count := i
		sum :=0
		for _,v :=range prevSort {
			sum +=v
			if sum <=remain {
				count--
			}else {
				res[i]=count
				break
			}
		}
	}
	return res
}

func sort(arr []int)[]int{
	if len(arr)<2{
		return arr
	}
	pivot :=arr[0]
	less := []int{}
	greater := []int{}
	for _,num :=range arr[1:]{
		if pivot > num {
			less = append(less,num)
		}else {
			greater = append(greater,num)
		}
	}
	less = append(sort(less),pivot)
	greater = sort(greater)
	return append(less,greater...)
}

func main(){
	//a :=[]int{1,2,3,4,5}
	//a := []int{4,3,2,1}
	//fmt.Println(minGiveUp(4,4,a))
	fmt.Println("begin")
	var num int
	fmt.Scan(&num)
	arr := make([][]int,num)
	line :=make([][]int,num)
	for i:=0;i<num;i++{
		line[i] = make([]int,2)
	}
	for i:=0;i<num;i++{
		var n,m int
		fmt.Scan(&n,&m)
		line[i][0] = n
		line[i][1] = m
		arr[i] = make([]int,n)
		//a :=make([]int,n)
		for j:=0;j<n;j++{
			fmt.Scan(&arr[i][j])
		}
	}
	for i:=0;i<num;i++{
		res := minGiveUp(line[i][0],line[i][1],arr[i])
		for i:=0;i<len(res);i++{
			fmt.Print(res[i]," ")
		}
		fmt.Println()
	}
}