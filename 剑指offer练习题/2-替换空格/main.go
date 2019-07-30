/*
请实现一个函数，将一个字符串中的每个空格替换成“%20”。
例如，当字符串为We Are Happy.则经过替换之后的字符串为We%20Are%20Happy。 
 */

 package main
 
 import (
	 "fmt"
 )
 //计算出替换之后字符串的总长度，从字符串的后面开始复制和替换。所有的字符只要移动一次 O(n)
 
 func replaceSpace(s string)string{
	 b := []byte(s)
	 originLength := len(b)
	 var blankAmount int    //字符串空格数
	 for _,v :=range b{
		 if v == ([]byte(" "))[0]{
			 blankAmount++
		 }
	 }
	 if blankAmount == 0{
		return s
	 }
	 newLength := originLength+2*blankAmount  //替换之后的字符串长度

	 newB :=make([]byte,newLength) 
	 
	 indexOfOriginal := originLength -1   //指向原字符串末尾
	 indexOfNew := newLength -1           //指向新字符串末尾

	 for indexOfOriginal >=0 && indexOfNew >= indexOfOriginal {
		 
		 
		 if b[indexOfOriginal] == ([]byte(" "))[0]{    //原字符串位置为空格,则替换
			 newB[indexOfNew] = ([]byte("0"))[0]       // "0"
			 indexOfNew--
			 newB[indexOfNew] = ([]byte("2"))[0]       // "2"
			 indexOfNew--
			 newB[indexOfNew] = ([]byte("%"))[0]       // "%"
			 indexOfNew--
		 }else {			 
			 newB[indexOfNew] = b[indexOfOriginal]
			 indexOfNew--
		 }
		 indexOfOriginal--		
	 }
	 
	 return string(newB)
 }
 func main() {
	s := "We are happy. "
	fmt.Println(replaceSpace(s),len(replaceSpace(s)))
 }