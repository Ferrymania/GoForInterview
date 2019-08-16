package main

import (
	"fmt"
	"github.com/Ferrymania/GoForInterview/dataStructure/queue/queueViaArray/arrayqueue"
)
func main() {
	q := arrayqueue.CreateQueue(10)

	for i:=0;i<6;i++ {
		q.EnQueue(i*2)
	}
	fmt.Println(q)
	q.DeQueue()
	fmt.Println("after dequeue",q)
}
