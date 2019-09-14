/*
给定一个正整数数组，最大为100个成员，从第一个成员开始，走到数组最后一个成员最少的步骤数，
第一步必须从第一元素开始，1<=步长<len/2,第二步开始以所在成员的数字走相应的步数，如果目标不可达返回-1
只输出最少的步骤数量
 */
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func max(a,b int)int {
	if a > b{
		return a
	}else {
		return b
	}
}
func GetMinSteps(nums []int)int{
	count,right,next := 0,0,0
	for i,v := range nums {
		if i > right {
			right = next
			count++
		}
		next = max(next,i+v)
	}
	if right != nums[len(nums)-1] {
		return -1
	}
	return count
}
func numbers(s []string)[]int{
	var arr []int
	for _,f := range s {
		i,err := strconv.Atoi(f)
		if err  == nil {
			arr = append(arr,i)
		}
	}
	return arr
}
func main(){
	arr := []int{}
	scanner := bufio.NewScanner(os.Stdin)
 	scanner.Scan()
	parts := strings.Split(scanner.Text()," ")
	arr = numbers(parts)

	//fmt.Println(arr)
	//nums := []int{7,5,9,4,2,6,8,3,5,4,3,9}
	step := GetMinSteps(arr)
	fmt.Println(step)
}