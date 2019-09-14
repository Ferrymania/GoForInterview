/*
从标准输入读取字符串，按照指定的两层分隔符切分成多对key-value，
列之间以' '分割，第一列表示key_value_delimiter
第二列表示key-value_delimiter ，第三列表示待切分的字符串。
分割符''不会出现列内容中

example
# : a:3#b:8#c:9
输出
3
a 3
b 8
c 9
 */
package main

import (
	"fmt"
	"strings"
)

func split(s string){
	//var kvpd []byte
	//var kvd []byte
	//sLen := len(s)
	//index :=0
	strArr :=strings.Split(s," ")

	kvpd := strArr[0]
	kvd := strArr[1]
	source :=strArr[2]
	sourceArr := strings.Split(source,kvpd)

	count :=0
	var dst []string
	for i:=0;i<len(sourceArr);i++{
		once := strings.Split(sourceArr[i],kvd)
		if once[0] != "" && once[1]!=""{
			count++
			temp := once[0] +" "+once[1]
			dst = append(dst,temp)
		}
	}
	fmt.Println(count)
	if count > 0{
		for i:=0;i<count;i++ {
			fmt.Println(dst[i])
		}
	}

}

func main(){
	str := "# : #a:#b:8#c:9"
	split(str)
}