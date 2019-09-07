/*
现有一组版本号字符串的数组，请找出里面版本号最小的值
版本号格式形如 major.minor.patch.ext,其中 major minor patch ext等
都是数字
每个版本号可能点分个数不一致，大小比较按照点分数字依次从前往后，
比如 7.10.2 > 7.3 ,3.1 > 2.3.4
 */
package main

import (
	"fmt"
	"strconv"
	"strings"
)

//输入一组string 版本号的数组，返回最小的版本号
func minVersion(versions []string)string{
	versionMap := make(map[int][]int)
	for i:=0;i<len(versions);i++{
		versionStr := strings.Split(versions[i],".")
		versionInt := make([]int,4)
		for i:=0;i<len(versionStr);i++{
			versionInt[i],_ = strconv.Atoi(versionStr[i])
		}
		fmt.Println(versionInt)
		versionMap[i] = append(versionMap[i],versionInt...)
	}
	for k,v :=range versionMap {
		for i:=0;i<4;i++{

		}
	}
	return "miss"
}

func main(){
	versions := []string{"3","4.3.5.4","2.10.3","2.4"}
	minVersion(versions)
}

func compare(arr ...int)[]int{
	min := arr[0]
	for i:=1;i<len(arr);i++{
		if arr[1] < min {
			min = arr[i]
		}
	}
	return nil
}