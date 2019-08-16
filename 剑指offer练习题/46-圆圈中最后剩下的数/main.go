/*
约瑟夫环问题，0，1，...,n-1这n个数字排成一个圆圈，从数字0开始每次从这个圆圈里删除第m个数字。求出这个圆圈里剩下的最后一个数字
 */
package main

import "fmt"

//用环形链表模拟圆圈，创建一个总共有n个结点，然后每次在这个链表中删除第m个结点
//时间复杂度O(mn)，空间复杂度O(n)
type node struct {
	num int
	next *node
}
func lastRemaining(n,m int) int {
	if n<1 || m<1{
		return -1
	}
	head := &node{}
	(*head).num = 0
	pre := head
	temp :=&node{}
	for i:=1;i<n;i++{
		(*temp).num = i
		(*pre).next = temp
		pre = temp
	}
	(*temp).next = head
	temp2 := &node{}
	for n!=1 {
		temp2 = head
		for i:=1;i<m-1;i++{
			temp2 = temp2.next
		}
		temp2.next = temp2.next.next
		//head = temp2.next
		n--
	}
	return (*head).num
}
//f（n,m）表示每次在n个数字0,1...n-1中每次删除第m个数字后剩下的数字
//f(n,m)={ 0    n=1
//       { [f(n-1,m)+m]%n    n>1
//时间复杂度 O(n) 空间复杂度O(1)
func lastRemaining_recursive(n,m int) int {
	if n <1 || m<1 {
		return -1
	}
	last := 0
	for i:=2;i<=n;i++{
		last = (last+m)%i
	}
	return last
}

func main() {
	fmt.Println(lastRemaining(5,3))
	fmt.Println(lastRemaining_recursive(5,3))
}