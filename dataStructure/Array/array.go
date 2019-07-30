/*
This is package about array data structure 
 */
package Array

type Array struct{
	data []int
	size int
}

//constructor with capacity parameters
func(arr *Array) Array(capacity int){
	arr.data = make([]int,0,capacity)
}
// //constructor without parameters
// func(arr *Array) Array() {
// 	arr.data = make([]int,10)
// }

//get number of elements in array
func(arr *Array) getSize() int {
	return arr.size
}
//get capacity of array
func(arr *Array) getCapacity() int {
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

	
}
func main(){

}