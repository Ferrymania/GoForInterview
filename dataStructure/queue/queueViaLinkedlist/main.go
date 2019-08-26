package main

import (
	"fmt"
	"github.com/Ferrymania/GoForInterview/dataStructure/queue/queueViaLinkedlist/linkedqueue"
)

func main(){
	q := &linkedqueue.LinkedQueue{}
	for i:=1;i<10;i++ {
		q.EnQueue(i*2)
	}
	fmt.Println(q)
	q.DeQueue()
	fmt.Println(q)
}
