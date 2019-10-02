/*
给出一个布尔表达式的字符串，比如：true or false and false，表达式只包含true，false，and和or，现
在要对这个表达式进行布尔求值，计算结果为真时输出true、为假时输出false，不合法的表达时输出error（
比如：true true）。表达式求值是注意and 的优先级比 or 要高，比如：true or false and
false，等价于 true or (false and false)，计算结果是 true

单个测试用例可能包含多组数据

输入每行包含布尔表达式字符串s，s只包含 true、false、and、or 几个单词（不会出现其它的任何单词），且单词之间用空格分隔。 (1 ≤ |s| ≤ 10^5)
输出true、false 或 error，true表示布尔表达式计算为真，false 表示布尔表达式计算为假，error 表示一个不合法的表达式
 */
package main

import "fmt"

func isPrio(top,newTop string)bool{
	if top == "or" && newTop == "and" {
		return true
	}
	return false
}
func postConvert(inop []string) error{
	stack := []string{}
	for i:=0;i<len(inop);i++{
		if i%2 == 0 {
			if inop[i] == "true" || inop[i] == "false"{
				stack = append(stack,inop[i])
			}else {
				fmt.Println("error")
				//return fmt.Errorf("error")
			}
		}else {
			if inop[i] == "or" {
				stack = append(stack,inop[i])
			}else if inop[i] == "and" {
				for len(stack)!= 0 {
					top := stack[len(stack)-1]
					if isPrio(top,inop[i]) {
						break
					}
					stack = append(stack,top)
				}
				stack = append(stack,inop[i])
			}else {
				fmt.Println("error")
				//return fmt.Errorf("error")
			}
		}
	}
	return nil
}
func calBool(op []string){
	stackBool := []string{}
	for i:=0;i<len(op);i++{
		if i%2 ==0 {
			if op[i] == " true" || op[i] == "false" {
				stackBool = append(stackBool)
			}else {
				fmt.Println("error")
			}
		}else {
			if op[i] == "and" || op[i] == "or" {
				bool1 := stackBool[len(stackBool)-1]
				bool2 := stackBool[len(stackBool)-2]
				if op[i] == "and" {
					if bool1 == "true" && bool2 == "true"{
						stackBool = append(stackBool,"true")
					}else {
						stackBool = append(stackBool,"false")
					}
				}
				if op[i] == "or" {
					if bool1 == "false" && bool2 == "false" {
						stackBool = append(stackBool,"true")
					}else {
						stackBool = append(stackBool,"true")
					}
				}
			}else {
				fmt.Println("error")
			}
		}
	}
}

func main(){
	str :=[]string{"true","or","false","and","false"}
	postConvert(str)
	fmt.Println(str)
}