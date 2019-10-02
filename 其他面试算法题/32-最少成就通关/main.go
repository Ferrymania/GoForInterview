package main

import "fmt"

func minTask(n int,taskList [][]int)[]int{
	taskComplete := make(map[int]bool,n)
	res := make([]int,n)
	for i:=1;i<=n;i++ {
		taskComplete[i] = false
	}
	fmt.Println(taskComplete)

	for i:=0;i<n;i++{
		count :=0
		if taskList[i][0]!= 0 {
			for j:=1;j<=taskList[i][0];j++{
				count += dfs(taskComplete,taskList[i][j])
			}
		}else {
			res[i]=0
		}
	}

	return res
}

func dfs(taskComplete map[int]bool,task int)int{
	if taskComplete[task] == true {
		return 1
	}
	count :=0
	for j:=0;j<len(taskList[i]);j++{
		 dfs(taskComplete,taskList[i][j])
	}

}

func main(){
	arr := [][]int{
		{1,2},
		{0},
		{1,2},
	}
	minTask(3,arr)
}
