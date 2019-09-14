/*
给定一个数组A[0,1,…,n-1],请构建一个数组B[0,1,…,n-1],
其中B中的元素B[i]=A[0]*A[1]*...*A[i-1]*A[i+1]*...*A[n-1]。不能使用除法。
 */
package main

import "fmt"

/*
B0    1   A1  A2  ...  A(n-2)  A(n-1)
B1    A0  1   A2  ...  A(n-2)  A(n-1)
...
B(n-2) A0  A1  A2        1     A(n-1)
b(n-1) A0  A1  A2   ...   A(n-1)  1
 */
func multiply(A []int)[]int{
	B := make([]int,len(A))
	for i:=0 ;i<len(A);i++{
		mul := 1
		for j :=0;j<len(A);j++{
			if i == j {
				mul *=1
				continue
			}
			mul *= A[j]
		}
		B[i] = mul
	}
	return B
}

func main() {
	A := []int{1,2,3,4,5}
	fmt.Println(multiply(A))
}