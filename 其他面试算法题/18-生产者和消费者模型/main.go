/*
生产者消费者问题是一个著名的线程同步问题，该问题描述如下：有一个生产者在生产产品，
这些产品将提供给若干个消费者去消费，为了使生产者和消费者能并发执行，在两者之间设置
一个具有多个缓冲区的缓冲池，生产者将它生产的产品放入一个缓冲区中，消费者可以从缓冲
区中取走产品进行消费，显然生产者和消费者之间必须保持同步，即不允许消费者到一个空的
缓冲区中取产品，也不允许生产者向一个已经放入产品的缓冲区中再次投放产品。
golang 的channel天生具有这种特性，即
①缓冲区满时写，缓冲区空时读，都会阻塞。
②channel 本身就是并发安全的。
golang实现多生产者多消费者：
 */
package main

import (
	"fmt"
	"time"
)

func consumer(cname string,ch chan int){
	for i:= range ch {
		fmt.Println("consumer--",cname,":",i)
	}
	fmt.Println("ch closed")
}

func producer(pname string ,ch chan int){
	for i:=0;i<4;i++{
		fmt.Println("producer--",pname,":",i)
		ch <-i
	}

}

func main(){
	data := make(chan int)
	go producer("生产者1", data)
	go producer("生产者2", data)
	go consumer("消费者1", data)
	go consumer("消费者2", data)

	time.Sleep(10 * time.Second)
	close(data)
	time.Sleep(10 * time.Second)
}