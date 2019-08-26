/*
LRU是Least Recently Used 的缩写，缓存一定量的数据，当超过设定的阈值时就把一些过期的数据删除掉。
常用于页面置换方法，是虚拟页式存储管理中常用的算法。
 */
package main

import "github.com/Ferrymania/GoForInterview/dataStructure/queue/queueViaArray/arrayqueue"

/*
使用两个数据结构实现一个LRU缓存
(1)使用双向链表实现的队列，队列的最大容量为缓存的大小。在使用的过程中
把最近使用的页面移动到队列首，最近没有使用的页面将放在队列尾的位置
（2）使用一个哈希表，把页号作为键，把缓存再队列中的结点的地址作为值
 */

type LRU struct {
	cacheSize int
	queue *arrayqueue.Queue
}