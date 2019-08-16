/*
Write a function that reverses a string. The input string is given as an array of characters char[].
Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.
You may assume all the characters consist of printable ascii characters.
 */
package main

import "fmt"

//use temp variable swap
func reverseString(s []byte)string{
	for i,j:=0,len(s)-1;i<j;i,j=i+1,j-1 {
		temp := s[i]
		s[i] = s[j]
		s[j] = temp
	}
	return string(s)
}
//use xor a^a = 0 a^0 = 0 To swap a,b
// a=a^b b=a^b a=a^b
func reverseString2(s []byte)string{
	for i,j:=0,len(s)-1;i<j;i,j=i+1,j-1{
		s[i] = s[i]^s[j]
		s[j] = s[i]^s[j]
		s[i] = s[i]^s[j]
	}
	return string(s)
}
func main() {
	s := "helloWorld"
	fmt.Println(reverseString([]byte(s)))
	fmt.Println(reverseString2([]byte(s)))
}