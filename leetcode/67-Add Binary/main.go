/*
Given two binary strings, return their sum (also a binary string).
The input strings are both non-empty and contains only characters 1 or 0.
Input: a = "11", b = "1"
Output: "100"
 */
package main

import "fmt"

func addBinary(a string, b string) string {

	aStr := []byte(a)
	bStr := []byte(b)
	aLength := len(aStr)
	bLength := len(bStr)
	var sum []byte
	toHigh := 0
	var temp int
	i := aLength-1
	j := bLength-1
	for i>=0 || j >=0 {
		temp = toHigh
		if i >= 0{
			temp += int(aStr[i]-'0')
		}
		if j>= 0 {
			temp += int(bStr[j]-'0')
		}
		toHigh = temp/2
		sum = append(sum,byte(temp%2))
	}
	if toHigh != 0 {
		sum = append(sum,byte(toHigh))
	}
	return string(sum)
}

func main(){
	a := "11"
	b := "1"
	fmt.Println("11+1=",addBinary(a,b))
}
