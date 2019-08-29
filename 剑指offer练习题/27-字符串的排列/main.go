/*
输入一个字符串，打印出该字符串中字符的所有排列。
 */
package main

import "fmt"

/*
第一步求所有可能出现在第一个位置的字符，即把第一个字符和后面所有的字符交换
第二步，固定第一个字符，求后面所有字符的排列。这个时候我们仍把后面的所有字符分成
两部分；后面字符的第一个字符，以及这个字符之后的所有字符，然后把第一个字符逐一和它后面的
字符交换
 */

func PermutationStr(s string){
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

func main() {
	PermutationStr("abc")
}