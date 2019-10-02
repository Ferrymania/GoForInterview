/*
从主串S指定的字符开始（一般为第一个）和模式串T的第一个字符开始，若相等，则继续逐个比较后续字符，、
直到T中的每个字符依次和S中的一个连续的字符序列相等，则称匹配成功；如果比较过程中有某对字符不相等
则从主串S的下一个字符起再重新和T的第一个字符比较。如果S中的字符都比完了仍然没有匹配成功
则称匹配不成功
时间复杂度最坏为O(m*n)
 */
package main

import "fmt"

//S 原字符串 T 模板串
func match(S ,T string)int{
	i:=0
	j:=0
	for i<len(S)&&j<len(T){
		if S[i] == T[j] {
			//继续比较后续字符
			i++
			j++
		}else{
				i=i-j+1
				j=0
		}
	}
	//匹配成功
	if j >=len(T) {
		index := i-j
		return index
	}else {
		return -1
	}
}

func main(){
	s :="ababcabcacbab"
	t :="abcac"
	fmt.Println(match(s,t))
}