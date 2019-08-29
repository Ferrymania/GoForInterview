/*

 */
package main

import "fmt"

//func math(str []string,length int){
//	var nums []int
//	var op []string
//	for i:=0;i<2*length-1;i++{
//		if i%2 == 0{
//			nums = append(nums,int(str[i]))
//		}else {
//			op = append(op,str[i])
//		}
//	}
//	fmt.Println(nums)
//	fmt.Println(op)
//	for i:=0;i<length-1;i++{
//		if i+1 <length-1{
//			for op[i] == "+" && op[i+1]
//		}
//
//	}
//}

func sort(arr []string,start,end int){
	var temp string
	for i:=start;i<=end;i++{
		for j:=end;j>i;j-- {
			if arr[j]<arr[j-1]{
				temp = arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = temp
			}
		}
	}
}

func main() {
	str := []string{"3","+","2","+","1","+","-4","*","-5","+","1"}
	math(str,6)
	arr := []string{"3","2","1","4"}
	sort(arr,0,2)
	fmt.Println(arr)
}

