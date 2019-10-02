/*
改进在于：每当一趟匹配过程中出现字符比较不等时，不需回溯i指针，而是利用已经得到的“部分匹配“的结果
将模式向右滑动尽可能远的以端距离后，继续进行比较
 */
package main


//next[i]表示当模式串匹配到T[i]遇到失配时，在模式串中需要重新和主串匹配的位置
//next数组的求解实际是对每个位置找到最长的公共前缀
/*
			0   当i=1时
next[i] =   Max{k|1<k<i且'P1..Pk-1'='Pj-k+1...Pi-1'
			1 其他情况
 */
func get_next(s string)[]int{
	next := []int{}
	i:=1

}
func kmpSearch(s,t string)int{

}
