/*
判断字符串是否是回文串
 */
package main

import "fmt"

func IsMirrorString(s string)bool{
	if len(s) < 2{
		return true
	}
	str := []byte(s)
	i:=0
	j:=len(str)-1
	for i<j&&str[i]==str[j]{
		i++
		j--
	}
	return i>=j
}

func main(){
	s :="abcdcbab"
	fmt.Println("Is it mirror string",IsMirrorString(s))
}
