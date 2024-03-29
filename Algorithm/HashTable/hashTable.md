## 哈希表（散列表）

**散列表**（**Hash table**，也叫**哈希表**），是根据键（Key）而直接访问在内存存储位置的数据结构。也就是说，它通过计算一个关于键值的函数，将所需查询的数据映射到表中一个位置来访问记录，这加快了查找速度。这个映射函数称做散列函数，存放记录的数组称做**散列表**

散列表适用于：

* 模拟映射关系
* 防止重复
* 缓存/记住数据，以免服务器再通过处理来生成它们

### 散列函数

一个把查找表中的关键字映射成该关键字对应的地址的函数，记为Hash(key)=Addr

### 哈希函数的设计

“键”通过哈希函数得到的“索引”分布**越均匀越好**

原则：

1. 一致性：如果a ==b ，则has(a) == hash(b)
2. 高效性：计算高效简便
3. 均匀性：哈希值均匀分布

* 整型
  * 小范围正整数直接使用
  * 小范围负整数进行偏移
  * 大整数 取模（分布不均匀）=> [模一个素数](https://planetmath.org/goodhashtableprimes)
* 浮点型
  * 转成整型处理
* 字符串
  * 转成整型处理
* 复合类型

#### 直接定址法

#### 除留余数法

最简单、最常用的方法，假定散列表长度为m,取一个不大于m但最接近或等于m的质数p，散列函数为

H(key)=key % P​

#### 数字分析法

#### 平方取中法

#### 折叠法

### 哈希冲突的处理

#### 开放定址法

##### 线性探测法

##### 平方探测法

##### 再散列法

##### 伪随机序列法

#### 拉链法

散列函数可能会把两个或两个以上的不同关键字映射到同一地址

如果给两个键分配的位置相同，就会发生冲突（collision）

避免冲突，需要有：

* 较低的填装因子（散列表包含的元素数/位置总数，超过0.7就该调整散列表的长度）
* 良好的散列函数

####

* 链地址法（seperate chaining）:如果两个键映射到了同一个位置，就在这个位置存储一个链表

### 散列函数的性能

取决于三个因素

* 散列函数
* 处理冲突的方法
* 装填因子：一个表的装满程度，a=表中记录数/散列表长度，a越大，表明装填的记录越满，发生冲突的可能性就越大，反之发生冲突的可能性越小。