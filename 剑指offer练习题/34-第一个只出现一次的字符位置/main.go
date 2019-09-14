/*
在字符串中找出第一个只出现一次的字符，如输入"abaccdeff"，则输出"b"
 */
package main

import "fmt"

func FirstNotRepeatingChar(s string)string{
	strArr := []rune(s)
	if strArr == nil {
		return ""
	}
	strMap := map[rune]int{}
	for i:=0;i<len(strArr);i++{
		strMap[strArr[i]]++
	}
	for i:=0;i<len(strArr);i++{
		if strMap[strArr[i]] == 1 {
			return string(strArr[i])
		}
	}
	return ""
}

func main(){
	str := "abaccdeff"
	fmt.Println(FirstNotRepeatingChar(str))
}
