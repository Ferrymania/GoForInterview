/*
牛客最近来了一个新员工Fish，每天早晨总是会拿着一本英文杂志，写些句子在本子上。同事Cat对Fish写的内
容颇感兴趣，有一天他向Fish借来翻看，但却读不懂它的意思。例如，“student. a am I”。后来才意识到，
这家伙原来把句子单词的顺序翻转了，正确的句子应该是“I am a student.”。Cat对一一翻转这些单词顺序
可不在行，你能帮助他么？
 */
package main

import (
	"fmt"
	"strings"
)

func ReverseSentence(str string)string{
	words := strings.Split(str," ")
	var newStr string
	for i:=len(words)-1;i>=0;i--{
		newStr += words[i]
		if i != 0 {
			newStr += " "
		}
	}
	return newStr
}

func main(){
	fmt.Println(ReverseSentence("student a am I"))
}
