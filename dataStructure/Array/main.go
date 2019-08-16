package main

import (
	"fmt"
	"github.com/Ferrymania/GoForInterview/dataStructure/Array/array"
)
func main() {
	arr := array.Constructor(10)
	fmt.Println(arr)

	for i:=0 ;i<6;i++ {
		arr.AddLast(i*2)
	}
	fmt.Println(arr)

	fmt.Println(arr.Contains(7))

	arr.Remove(3)
	fmt.Println(arr)
}
