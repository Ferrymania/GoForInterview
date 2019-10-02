/*
输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。
 */
package main

import (
	"fmt"
)

func PrintMatrixClockwisely(arr [][]int,rows int,cols int){
	if arr == nil || rows <=0 || cols <= 0{
		return
	}
	start :=0
	for rows > start*2 && rows > start*2 {
		PrintMatrixCircle(arr,rows,cols,start)
		start++
	}
}
//每圈打印
func PrintMatrixCircle(arr [][]int,rows ,cols,start int){
	endX := cols -1 -start
	endY := rows -1 -start

	for i:=start;i<=endX;i++{
		number := arr[start][i]
		fmt.Print(number,"->")
	}

	if start<endY {
		for i:=start+1;i<=endY;i++{
			number:=arr[i][endX]
			fmt.Print(number,"->")
		}
	}

	if start <endX && start <endY {
		for i:=endX-1;i>=start;i--{
			number := arr[endY][i]
			fmt.Print(number,"->")
		}
	}

	if start <endX && start<endY-1{
		for i :=endY-1;i>=start+1;i--{
			number := arr[i][start]
			fmt.Print(number,"->")
		}
	}
}


func main(){
	arr :=[][]int{
		[]int{1,2,3,4},
		[]int{5,6,7,8},
		[]int{9,10,11,12},
		[]int{13,14,15,16},
	}
	PrintMatrixClockwisely(arr,4,4)
}