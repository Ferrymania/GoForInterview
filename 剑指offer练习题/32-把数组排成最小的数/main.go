/*
输入一个正整数数组，把数组里所有数字拼接起来排成一个数
打印能拼接出的所有数字中最小的一个。
例如 输入数组{3，32，321}则打印出这3个数字能排成的最小数字321323
 */
package main

import "strconv"

/*
n个数字共有n!个排列

先将数组转换成字符串数组，然后对字符串数组按照规则排序，最后将排好序的字符串数组拼接出来。
关键就是制定排序规则：
若ab > ba 则 a > b
若ab < ba 则 a < b
若ab = ba 则 a = b
 */

func PrintMinNumber(number []int)string{
	if number == nil {
		return ""
	}
	length := len(number)
	if length == 1 {
		return strconv.Itoa(number[0])
	}
	strArr := make([]string,length)
	for i:=0;i<length;i++{
		strArr[i] = strconv.Itoa(number[i])
	}


}



