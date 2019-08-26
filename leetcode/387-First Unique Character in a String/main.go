/*
Given a string, find the first non-repeating character in it and return it's index. If it doesn't exist, return -1.
Examples:

s = "leetcode"
return 0.

s = "loveleetcode",
return 2.
 */
package main

import "fmt"

func firstUniqChar(s string) int {
	var freq [26]int
	for _,v :=range s {
		freq[v-'a'] ++
	}
	for i := range s{
		if freq[s[i]-'a'] == 1{
			return i
		}
	}
	return -1
}

func main() {
	s:= "leetcode"
	fmt.Println(firstUniqChar(s))
}
