/*
数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字
 */
package main

import "fmt"

/*
数组中有一个数字出现的次数超过数组长度的一般，也就是说它出现的次数比其他所有数字出现的次数的和还要多。
在遍历数组的时候保存两个值：一个是数组中的一个数字，一个是次数。当我们遍历到下一个数字的时候，
如果下一个数字和我们之前保存的数字相同，则次数加1；如果下一个数字和我们之前保存的数字不同
则次数减1.由于要找的数字出现的次数比其他所有数字出现的次数之和还要多，那么要找的数字肯定是
最后一次把次数设为1时对应的数字
 */

func MoreThanHalfNum(arr []int,length int) int{
	if arr == nil || length <=0{
		panic("input is NOT valid")
	}
	result := arr[0]
	times := 1
	for i:=1;i<length;i++{
		if times == 0 {
			result = arr[i]
			times = 1
		}else if arr[i] == result{
			times++
		}else{
			times--
		}
	}
	if !CheckMoreThanHalf(arr,length,result){
		return 0
	}
	return result
}

func CheckMoreThanHalf(arr []int,length,num int) bool{
	times := 0
	for i := 0;i<length;i++{
		if arr[i] == num{
			times++
		}
	}
	if times*2 <= length{
		return false
	}
	return true
}

func main() {
	arr := []int{1,2,3,2,2,2,5,4,2}
	fmt.Println(" array's max times number is ",MoreThanHalfNum(arr,9))
}