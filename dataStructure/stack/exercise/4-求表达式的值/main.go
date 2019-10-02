/*
计算四则表达式的值
 */
package main

import "fmt"

type stack struct {
	arr []byte
	size int
}
func constructor()*stack{
	return &stack{arr :[]byte{}}
}
func(s *stack)isEmpty()bool{
	return s.size == 0
}
func(s *stack)push(char byte){
	s.arr = append(s.arr,char)
	s.size++
}
func(s *stack)pop()byte {
	if s.isEmpty() {
		panic("stack is empty")
	}
	s.size--
	char := s.arr[s.size]
	s.arr = s.arr[:s.size]
	return char
}
func(s *stack)peek()byte{
	if s.isEmpty(){
		panic("stack is empty")
	}
	return s.arr[s.size-1]
}


/*
计算四则表达式 首先把中缀表达式转成后缀表达式
初始化两个栈：运算符栈S1和储存中间结果的栈S2
过程如下：从左向右扫描，
遇到数字时，将其压入S2
遇到运算符时：
	a 若S1为空，或栈顶元素为左括号'('，入栈
	b 否则，若优先级比S1栈顶运算符的高，也将运算符压入S1
	c 否则，将S1栈顶的运算符弹出并压入S2,再次转到a，与S1中新的栈顶元素
遇到括号时：
	a 如果是左括号"(",则直接压入s1
	b 如果是右括号')',则依次弹出S1栈顶运算符，直到遇到左括号为止，此时将这一对括号丢弃
重复上述步骤，直到表达式最右边
将S1剩余的运算符依次弹出并压入S2
依次弹出S2中的元素并输出，结果的逆序即为后缀表达式
 */

func convertPost(str string)*stack{
	s1 :=constructor()
	s2 := constructor()
	for i:=0;i<len(str);i++{
		if str[i]>='0' && str[i]<='9'{
			s2.push(str[i])
			continue
		}
		if s1.isEmpty() || s2.peek() == '(' || str[i] == '('{
			s1.push(str[i])
		}else{}
	}
	return s2
}
func main(){
	s :=constructor()
	str := "213*234-12+324"
	for i:=0;i<len(str);i++{
		s.push(str[i])
	}
	fmt.Println(s.arr)
	for i:=0;i<len(str);i++{
		char := s.pop()
		fmt.Println(char)
	}
	fmt.Println(s.isEmpty())
}