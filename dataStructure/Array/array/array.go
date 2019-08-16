/*
This is package about array data structure 
 */
package array

import (
	"bytes"
	"fmt"
)

type Array struct{
	data []int
	size int
}

//constructor with capacity parameters
func Constructor(capacity int) *Array{
	return &Array{
		data:make([]int,capacity),
	}

}
// //constructor without parameters
// func(arr *Array) Array() {
// 	arr.data = make([]int,10)
// }

//get number of elements in array
func(arr *Array) GetSize() int {
	return arr.size
}
//get capacity of array
func(arr *Array) GetCapacity() int {
	return cap(arr.data)
}

func(arr *Array) IsEmpty() bool {
	return arr.size == 0
}

func(arr *Array) Add(index int,e int){
	if arr.size == len(arr.data) {
		panic("Add failed.Array is ful")
	}
	if(index <0 || index > arr.size){
		panic("add failed.Require index >= 0 and index <= size")
	}
	arr.data[index] = e
	arr.size++
}

func(arr *Array) AddFirst(e int){
	arr.Add(0,e)
}

func(arr *Array) AddLast(e int){
	arr.Add(arr.size,e)
}

func(arr *Array) Get(index int) int {
	if index <0 || index >= arr.size {
		panic("Get failed .Index is illegal")
	}

	return arr.data[index]
}

func(arr *Array) Set(index int,e int){
	if index <0 || index >=arr.size{
		panic("Set failed.Index is illegal")
	}

	arr.data[index] = e
}

func(arr *Array) Contains(e int) bool {
	for _,v :=range arr.data{
		if v == e {
			return true
		}
	}

	return false
}


func (arr *Array) Find(e int) int{
	for i,v := range arr.data {
		if v == e {
			return i
		}
	}

	return -1
}

func(arr *Array) Remove(index int) int {
	if index < 0 || index >= arr.size {
		panic("Remove failed.Index is illegal")
	}

	e := arr.data[index]
	for i:= index + 1;i<arr.size;i++{
		arr.data[i-1] = arr.data[i]
	}
	arr.size--
	return e
}

func(arr *Array) String() string {
	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("Array:size = %d,capacity = %d\n",arr.size,len(arr.data)))
	buffer.WriteString("[")
	for i := 0;i<arr.size;i++ {
		buffer.WriteString(fmt.Sprint(arr.data[i]))
		if i != arr.size-1{
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("]")
	return buffer.String()
}


