/*
找出两个字符串的最长公共字串。
 */
package main

import (
	"bytes"
	"fmt"
)

/*
把字符串分别以行和列组成一个二维矩阵
比较二维矩阵中每个点对应行列字符是否相等，相等的话值设置为1，否则设置为0
通过查找出值为1的最长对角线就能找到最长公共子串
 		d  g  c  a  d  d  e
	a	0  0  0  1  0  0  0
	b   0  0  0  0  0  0  0
	c   0  0  1  0  0  0  0
	c   0  0  1  0  0  0  0
	a   0  0  0  1  0  0  0
	d   1  0  0  0  1  1  0
	e   0  0  0  0  0  0  1
 */

func getMaxSubStr(s1 ,s2 string)string{

	// 假设str1长度 >= str2长度
	var str1,str2 []byte
	if len(s1)>=len(s2){
		str1 = []byte(s1)
		str2 = []byte(s2)
	}else{
		str1 = []byte(s2)
		str2 = []byte(s1)
	}

	temp := make([][]int,len(str1))
	for i:=range temp{
		temp[i] = make([]int,len(str2))
	}
	//最长公共子串长度
	maxLength := 0
	//记录最长子串最后一个位置
	maxLocation := 0
	for i:=0;i<len(str1);i++{
		for j:=0;j<len(str2);j++{
			if str1[i] == str2[j] {
				if i>0 && j>0 {
					temp[i][j] = temp[i-1][j-1]+1
				}else {
					temp[i][j] = 1
				}

				if temp[i][j] > maxLength {
					maxLength = temp[i][j]
					maxLocation = i
				}
			}else{
				temp[i][j] = 0
			}
		}
	}
	fmt.Println("max sub string length is",maxLength,"location is ",maxLocation)
	//拼接公共子串
	var buf bytes.Buffer
	for i:= maxLocation-maxLength+1;i<=maxLocation;i++{
		buf.WriteByte(str1[i])
	}
	return buf.String()
}
func main(){
	s1 := "dgcadedegfadfa"
	s2 := "abccadedasf"
	fmt.Println("max length is ",getMaxSubStr(s1,s2))
}
