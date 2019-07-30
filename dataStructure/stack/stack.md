## 栈

栈是一种线性结构，相比数组，栈对应的操作是数组的子集，是一种后进先出（Last in first out）的数据结构

栈的实现有两种方法，分别采用数组来实现和采用链表实现。

### 数组实现栈

采用数组实现栈的时候，栈空间是一段连续的空间。

复杂度：

* push O(1)
* pop O(1)
* peek() O(1)
* getSize() O(1)
* isEmpty() O(1)

### 链表实现栈
