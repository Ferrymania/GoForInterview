/*
在字符串中找出第一个只出现一次的字符。如输入"abaccdeff",则输出'b'
 */
package main

import "fmt"

/*
brute force
从头开始扫描这个字符串中的每个字符，当访问到某字符时，拿这个字符和后面的每个字符相比较，如果没有发现重复的字符，则该字符就是只出现一次的字符。
时间复杂度O(n)

利用哈希表的特性，遍历两次字符串，第一次获得计算每个字符串的次数
第二次从头开始找出次数为1的字符
时间复杂度O(n),空间复杂度O(1)
 */

func FirstNotRepeatingChar(s string)byte{
	if len(s)==0 {
		panic("There is no repeat char because of the empty string")
	}
	m := map[byte]int{}
	for i:=0;i<len(s);i++{
		m[s[i]] ++
	}
	for i:=0;i<len(s);i++{
		if m[s[i]] == 1{
			return s[i]
		}
	}
	return ' '
}

func main(){
	str := "abaccdeff"
	b := FirstNotRepeatingChar(str)
	fmt.Println(string(b))
}