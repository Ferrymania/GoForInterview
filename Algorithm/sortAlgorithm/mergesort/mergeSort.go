/*
create mergesort
 */
package main

import "fmt"

// merge array,l is first element of array,m is the element array seperated by; r is the last element of array
func Merge(arr []int,L,M,R int){
	leftArray := make([]int,M-L)
	rightArray :=make([]int,R-M+1)

	for i:=L;i<M;i++{
		leftArray[i-L] = arr[i]
	}

	for i:=M;i<=R;i++{
		rightArray[i-M]=arr[i]
	}

	i,j := 0,0
	k :=L

	for i< len(leftArray) && j < len(rightArray) {
		if leftArray[i] < rightArray[j] {
			arr[k] = leftArray[i]
			i++
			k++
		}else{
			arr[k] = rightArray[j]
			j++
			k++
		}

	}

	for i<len(leftArray) {
		arr[k] = leftArray[i]
		i++
		k++
	}

	for j<len(rightArray) {
		arr[k] = rightArray[j]
		k++
		j++
	}
}

//mergesort  arr arrays ,l is index of first elememnt of array,r is latst element of array
func MergeSort(arr []int,l,r int){
	// one element return directly
	if l==r {
		return
	}else{
		m:=(l+r)/2
		MergeSort(arr,l,m)
		MergeSort(arr,m+1,r)
		Merge(arr,l,m+1,r)
	}
}

func main() {
	arr :=[]int{38,65,97,76,13,27,49}
	MergeSort(arr,0,len(arr)-1)
	fmt.Println("after sort",arr)
}
