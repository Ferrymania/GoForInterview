## 队列

队列也是一种线性结构

相比数组，队列对应的操作是数组的子集

只能从一端（队尾）添加元素，只能从另一端（队首）取出元素

### 数组实现

用front来记录队列首元素的位置，用rear来记录队列尾元素往后一个位置。入队列的时候只需要将待入队列的元素放到数组下标为rear的位置，同时执行rear++,出队列的时候只需要执行front++即可

#### 顺序队列

初始状态：`Q.front=Q.rear=0`

进队操作：队不满时，先送值到队尾元素，再将队尾指针加1

出队操作：队不空时，先取对头元素值，再将队头指针加1

#### 循环队列

把顺序队列想成一个环状的空间，即把存储队列元素的表从逻辑上看成一个环，称为循环队列。当队首指针`Q.front=Maxsize-1`,再前进一个位置就自动到0，这可以利用除法取余运算（%）来实现

初始时：Q.front=Q.rear=0

队首指针进1:`Q.front=(Q.front+1)%Maxsize`
队尾指针进1：`Q.rear = (Q.rear+1)%Maxsize`

队列长度：`(Q.rear+Maxsize-Q.front)%Maxsize`

#### 缺点

出队列后数组前半部分的空间不能够充分的利用，解决方法是把数组看成一个环状的空间（循环队列）。当数组最后一个位置被占用后，可以从数组首位置开始循环利用。

### 链表实现

