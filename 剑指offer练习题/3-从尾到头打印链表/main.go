/*
输入一个链表，按链表值从尾到头的顺序返回一个ArrayList
思路：遍历链表是从头到尾--->从尾到头是后进先出----->栈的结构----->递归结构
 */
package main

import (
	"fmt"
	"github.com/Ferrymania/GoForInterview/dataStructure/Linkedlist/singlylinkedlist"
)
func printListReversingly_Recursive(n *singlylinkedlist.Node){
	if(n != nil){
		if(n.Next != nil) {
			printListReversingly_Recursive(n.Next)
		}
		fmt.Printf("%v\t",n.Value);
	}
}

func main() {
	fmt.Println("Hello link")
	ll :=singlylinkedlist.CreateLinkList()

	for i :=0;i<6;i++ {
		ll.AddToTail(i*2)
	}
	ll.PrintLinkedList()
	printListReversingly_Recursive(ll.Head)

} 
