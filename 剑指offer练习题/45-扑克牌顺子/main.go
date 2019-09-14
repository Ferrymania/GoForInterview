/*
从扑克牌中随机抽5张牌，判断是不是一个顺子，即这5张牌是不是连续的。2~10为数字本身，
A为1，J为11，Q为12，K为13，而大、小王可以看成任意数字
 */
package main

import "fmt"

/*
首先把数组排序吗，再统计数组中0的个数，最后统计排序之后数组中相邻数字之间的空缺总数。
如果空缺的总数小于或者等于0的个数，那么这个数组就是连续的；反之则不连续。
 */
func IsContinous(numbers []int)bool {
	if len(numbers) != 5 {
		return false
	}
	arr := quickSort(numbers)
	fmt.Println(arr)
	zero := 0  //大小王的个数
	gap := 0  //间隔数
	for i:=0;i<5;i++{
		if arr[i] == 0 {
			zero++
		}
	}
	for i:=zero;i<4;i++ {
		if arr[i] == arr[i+1] {
			return false
		}
		gap += arr[i+1]-arr[i]-1
	}
	fmt.Println("zero have ",zero,"gap have ",gap)
	if gap > zero {
		return false
	}
	return true
}

func quickSort(arr []int)[]int{
	//基线条件，数组为空或只包含一个元素的数组是有序的
	if len(arr) < 2{
		return arr
	}
	pivot := arr[0]  //选定首项为基准值
	less := []int{}  //小于基准值的数字组成的数组
	greater :=[]int{} //大于基准值的数字组成的数组
	for _,num :=range arr[1:]{
		if pivot > num{
			less = append(less,num)   //比基准值小，移向左边的数组
		}else {
			greater = append(greater,num)  //比基准值大，移向右边的数组
		}
	}
	less = append(quickSort(less),pivot)  //再对左边的子数组进行排序
	greater = quickSort(greater)  //再对右边的子数组进行排序
	return append(less,greater...)
}

func main() {
	nums := []int {0,2,5,6,3}
	fmt.Println(IsContinous(nums))
}