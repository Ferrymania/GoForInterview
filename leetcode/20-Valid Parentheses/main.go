/*
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Note that an empty string is also considered valid.
 */
package main

import (
	"fmt"
	"github.com/Ferrymania/GoForInterview/dataStructure/stack/stackViaLinkedlist/linkedstack"
)

func isValid(s string) bool {
	if s == ""{
		return true
	}
	stack := linkedstack.NewLinkedStack()
	for i:=0;i<len(s);i++ {
		if s[i] == '(' || s[i] =='['|| s[i] =='{' {
			stack.Push(s[i])
			fmt.Println("insert paren")
			fmt.Println(stack)
		}else {
			topChar := stack.Pop()
			if s[i] == ')' &&  topChar != uint8('(') {
				fmt.Println("1 not match")
				return false
			}
			if  s[i] == ']' &&  topChar != uint8('[') {
				fmt.Println("top char ",topChar)
				fmt.Println("2 not match")
				return false
			}
			if s[i] == '}' &&  topChar != uint8('{') {

				fmt.Println("3 not match")
				return false
			}
		}
	}
	fmt.Println(stack.IsEmpty())
	return stack.IsEmpty()
}

func main() {
	s := "{{[123]}}"
	fmt.Println("Is string valid parentheses",isValid(s))
}

