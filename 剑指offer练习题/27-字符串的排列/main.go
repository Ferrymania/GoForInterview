/*
输入一个字符串，打印出该字符串中字符的所有排列。
 */
package main

import "fmt"

/*
第一步求所有可能出现在第一个位置的字符，即把第一个字符和后面所有的字符交换
第二步，固定第一个字符，求后面所有字符的排列。这个时候我们仍把后面的所有字符分成两部分；
后面字符的第一个字符，以及这个字符之后的所有字符，然后把第一个字符逐一和它后面的字符交换
递归法：时间复杂度O(n) 空间复杂度O(1)
 */

func PermutationStr(s string){
	if s == "" {
		return
	}
	Permutation([]byte(s),0)
}

func Permutation(str []byte,start int){
	if str == nil {
		return
	}
	if start == len(str)-1{
		fmt.Println(string(str))
	}else {
		for i:=start;i<len(str);i++ {
			//判断是否有重复字符
			if !isDuplicate(str ,start,i){
				continue
			}
			swapchar(str,start,i)
			Permutation(str,start+1)
			swapchar(str,start,i)
		}
	}

}

func swapchar(str []byte,i,j int){
	temp := str[i]
	str[i] = str[j]
	str[j] = temp
}

func isDuplicate(str []byte,begin ,end int)bool{
	for i:=begin;i<end;i++{
		if str[i] == str[end]{
			return false
		}
	}
	return true
}

func main() {
	PermutationStr("aba")
}