package main

import (
	"fmt"
	"github.com/Ferrymania/GoForInterview/dataStructure/Linkedlist/singlylinkedlist"
)

func main() {
	singlylinkedlist := singlylinkedlist.CreateLinkList()

	for i := 0; i < 5; i++ {
		singlylinkedlist.AddToHead(i)
		fmt.Println(singlylinkedlist)
	}

	singlylinkedlist.Add(2, 666)
	fmt.Println(singlylinkedlist)
	singlylinkedlist.PrintLinkedList()
	fmt.Println("Does linkedlist contanis 666", singlylinkedlist.Contains(666))
	singlylinkedlist.Delete(4)
	singlylinkedlist.PrintLinkedList()
}
