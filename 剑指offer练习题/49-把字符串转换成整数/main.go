/*
将一个字符串转换成一个整数(实现Integer.valueOf(string)的功能，但是string不符合数字要求时返回0)，
要求不能使用字符串转换整数的库函数。 数值为0或者字符串不是一个合法的数值则返回0。
 */
package main

import "fmt"

/*
设置三个标志符分别记录“+/-”、“e/E”和“.”是否出现过。

对于“+/-”：正常来看它们第一次出现的话应该出现在字符串的第一个位置，如果它第一次出现在不是字符串首位，而且它的前面也不是“e/E”，那就不符合规则；如果是第二次出现，那么它就应该出现在“e/E”的后面，如果“+/-”的前面不是“e/E”，那也不符合规则。

对于“e/E”：如果它的后面不接任何数字，就不符合规则；如果出现多个“e/E”也不符合规则。

对于“.”：出现多个“.”是不符合规则的。还有“e/E”的字符串出现“.”也是不符合规则的。
同时，要保证其他字符均为 0-9 之间的数字。
 */
func StrToInt(str string)int{

	mul :=1
	sum :=0
	for i:=len(str)-1;i>=0;i--{
		if i == 0 {
			if  str[i]<='9' && str[i]>='0' {
				n := str[i]-'0'
				sum += int(n)*mul
				mul *=10
				break
			}
			if str[i] == '+' {
				break
			}
			if str[i] == '-' {
				sum = 0-sum
				break
			}
			if  str[i]>'9' || str[i]<'0' {
				return 0
			}

		}else {
			if str[i]>'9' || str[i]<'0' {
				return 0
			}
			n := str[i]-'0'
			sum += int(n)*mul
			mul *=10
		}
	}
	return sum
}

func main(){
	n := StrToInt("-123")
	fmt.Println(n)
}