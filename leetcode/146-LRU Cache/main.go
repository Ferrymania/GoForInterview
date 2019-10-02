/*
Design and implement a data structure for Least Recently Used (LRU) cache. It should support the following operations: get and put.

get(key) - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
put(key, value) - Set or insert the value if the key is not already present. When the cache reached its capacity, it should invalidate the least recently used item before inserting a new item.

The cache is initialized with a positive capacity.
 */
package main

type node struct {
	prev *node
	next *node
	key int
	val int
}
type LRUCache struct {
	head *node
	tail *node
	store map[int]*node
	cap int
}


func Constructor(capacity int) LRUCache {
	return LRUCache{cap:capacity,store:make(map[int]*node)}
}

func(this *LRUCache)removeFromChain(n *node){
	if n == this.head{
		this.head = n.next
	}
	if n== this.tail {
		this.tail = n.prev
	}
	if n.next != nil {
		n.next.prev = n.prev
	}

	if n.prev != nil {
		n.prev.next = n.next
	}
}

func(this *LRUCache)addToChain(n *node){
	n.prev = nil
	n.next = this.head
	if n.next != nil {
		n.next.prev = n
	}
	this.head = n
	if this.tail == nil {
		this.tail =n
	}
}


func (this *LRUCache) Get(key int) int {
	n,ok := this.store[key]
	if !ok {
		return -1
	}
	this.removeFromChain(n)
	this.addToChain(n)
	return n.val
}


func (this *LRUCache) Put(key int, value int)  {
	n,ok := this.store[key]
	if !ok {
		n = &node{val:value,key:key}
		this.store[key] = n
	}else {
		n.val = value
		this.removeFromChain(n)
	}

	this.addToChain(n)
	if len(this.store) > this.cap {
		n = this.tail
		if n != nil {
			this.removeFromChain(n)
			delete(this.store,n.key)
		}
	}
}