/*
地上有一个m行和n列的方格。一个机器人从坐标0,0的格子开始移动，每一次只能向左，右，上，下四个方向
移动一格，但是不能进入行坐标和列坐标的数位之和大于k的格子。 例如，当k为18时，机器人能够进入方格（
35,37），因为3+5+3+7 = 18。但是，它不能进入方格（35,38），因为3+5+3+8 = 19。请问该机器人能够达
到多少个格子？

 */
package main

import "fmt"

/*
回溯法
机器人从坐标（0，0）开始移动。当它准备进入座标为（i，j）的格子时，通过检查坐标的数位和 来判断机器人是否能够进入。如果机器人能够进入坐标为（i,j）的格子，则再判断它能否进入4个相邻的格子
(i,j-1),(i-1,j)(i,j+1)(i+1,j)
 */
func movingCount(threshold ,rows,cols int)int{
	if threshold < 0 || rows <=0||cols<=0{
		return 0
	}
	visited := make([]bool,rows*cols)
	count:=movingCountCore(threshold,rows,cols,0,0,visited)
	return count
}

func movingCountCore(threshold int,rows int,cols int,row int,col int,visited []bool)int{
	count :=0
	if check(threshold,rows,cols,row,col,visited){
		fmt.Println("step fowards")
		visited[row*cols+col] = true
		count = 1 + movingCountCore(threshold,rows,cols,row-1,col,visited)+
			movingCountCore(threshold,rows,cols,row,col-1,visited)+
			movingCountCore(threshold,rows,cols,row+1,col,visited)+
			movingCountCore(threshold,rows,cols,row,col+1,visited)
	}
	return count
}

func check(threshold int,rows int,cols int,row int, col int,visited []bool)bool{
	if row>=0 && row < rows && col >=0 && col <cols&&
		calSum(row)+calSum(col)<=threshold &&
		!visited[row*cols+col]{
		return true
	}
	return false
}

func calSum(n int)int{
	s :=0
	for n!=0 {
		s += n%10
		n = n/10
	}
	return s
}

func main(){
	fmt.Println(movingCount(10,3,4))
}