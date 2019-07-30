/*在一个二维数组中（每个一维数组的长度相同），每一行都按照从左到右递增的顺序排序，
每一列都按照从上到下递增的顺序排序。
请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。*/

package main

import (
	"fmt"
)

//首先选取数组中右上角的数字。如果该数组等于要查找的数字，查找过程结束
//如果该数字大于要查找的数字，剔除这个数字所在的列
//如果该数字小于要查找的数字，剔除这个数字所在的行
func find(a [][]int, number int) bool {
	found := false
	var rows, cols int
	if len(a) != 0 {
		rows = len(a)    //二维数组行数
		cols = len(a[0]) //二维数组列数
	} else {
		return found
	}
	row := 0
	col := cols - 1
	for row < rows && col >= 0 {
		if a[row][col] == number {    //从右上角的数字开始，若相等则返回
			found = true
			break
		} else if a[row][col] > number { //如果大于目标数，则移向左边一列
			col--
		} else {                         //如果小于目标数，则移向下一行
			row++
		}
	}
	return found
}


//首先选取数组中左下角的数字。如果该数组等于要查找的数字，查找过程结束
//如果该数字大于要查找的数字，剔除这个数字所在的行
//如果该数字小于要查找的数字，剔除这个数字所在的列
func find2(a [][]int, number int) bool {
	found := false
	var rows, cols int
	if len(a) != 0 {
		rows = len(a)    //二维数组行数
		cols = len(a[0]) //二维数组列数
	} else {
		return found
	}
	row := rows-1
	col := 0
	for row >= 0 && col < cols {
		if a[row][col] == number {    //从右上角的数字开始，若相等则返回
			found = true
			break
		} else if a[row][col] > number { //如果大于目标数，则移向上一行
			row--
		} else {                         //如果小于目标数，则移向右边一列
			col++
		}
	}
	return found
}

func main() {
	a := [][]int{
		[]int{1, 2, 8, 9},
		[]int{2, 4, 9, 12},
		[]int{4, 7, 10, 13},
		[]int{6, 8, 11, 15},
	}

	var b [][]int
	fmt.Println(len(b))
	// fmt.Println(a)
	fmt.Println(find(a, 1))
	fmt.Println(find(a, 6))
	fmt.Println(find(a, 15))
	fmt.Println(find(a, 22))

	fmt.Println(find2(a, 1))
	fmt.Println(find2(a, 6))
	fmt.Println(find2(a, 15))
	fmt.Println(find2(a, 22))
	// fmt.Println(find(b,1))
}
