问题 ：MySQL 数据库的索引有哪些？ B+ 树了解吗，有哪些特点？为什么用 B+ 树而不用其他树？这个和磁盘预读有什么关系？乐观锁和悲观锁是什么？现在让你实现一个乐观锁说说思路？

主键和索引的区别 聚簇索引和非聚簇索引，为什么要用 B + Tree 一条 SQL 语句执行的流程，尽量说详细 AOP 的实现原理  MySQL 的隔离机制 MySQL 的锁机制 数据库的内连接，左连接，右连接的区别 数据库常见引擎和隔离级别 如何判断 SQL 语句用到了哪些锁;6.MySQL 索引的类型 min B树什么时候的高度会变高； 数据库隔离级别、索引、RR隔离级别下的SQL查询结果 MySQL 主从复制的方式

 查看数据库设计、插入一列、查找排名前10的同学

数据库： mysql的底层数据结构，为什么用b+树（我自己讲了一下b+树与b树、红黑树、数组、链表的区别） 聚簇索引、非聚簇索引、覆盖索引、多列索引（在a、b上建索引，分别查询a和b，都有效么（最左前缀原则）） 事务隔离级别、视图、怎么样找出慢查询、乐观锁与悲观锁、innodb引擎下的mysql锁

事务隔离级别有哪些，MySQL默认事务隔离级别是啥（可重复读），可重复读是什么意思，可重复读的实现方式（除了行锁还有啥？）

数据库的事务和事务的特性、脏读幻读的定义

## SQL语句

* SQL 方案（schema）语句:用于定义存储于数据库中的数据结构
* SQL数据语句：用于操作SQL方案语句所定义的数据结构
* SQL事务语句：用于开始、结束或回滚事务

创建数据库

``` sql
create database bank;
```

关联数据库

``` sql
use bank;
```

查看数据库、表、列、用户、权限等信息

``` sql
show databases;
show tables;
show columns from customers;
show status;
show grants;
show errors;
shoiw warnings;
```

创建表

``` mysql 
create table customers
(
    cust_id int 	NOt NULL AUTO_INCREMENT,
    cust_name char(50) NOT NULL,
    cust_address char(50) NULL,
    cust_city char(50) NULL,
    cust_state char(5) NULL,
    cust_zip char(10) NULL,
    cust_country char(50) NULL,
    cust_contact char(50) NULL,
    cust_email char(255) NULL,
    PRIMARY KEY (cust_id)
)ENGINE=InnoDB;
```

**每个表中只允许一个`AUTO_INCREMENT`列，而且它必须被索引（如通过使它成为主键）。**

指定默认值

``` mysql
create table orderitems
(
    order_num int not null,
    order_item int NOT  NULL,
    quantity int NOT NULL DEFAULT 1,
    item_price decimal(8,2) NOT NULL,
    RPIMARY KEY(order_num,order_item)
)ENGINE=InnoDB;
```

更新表

``` mysql
alter table vendors add vend_phone char(20);

alter table vendors drop column vend_phone;
```

删除表

``` mysql
drop table customer2;
```

重命名表

``` mysql
rename table customers2 to customers;
```



### select 语句

子句顺序

| 子句     | 说明 | 是否必须使用 |
| -------- | ---- | ------------ |
| SELECT   |      |              |
| FROM     |      |              |
| WHERE    |      |              |
| GROUP BY |      |              |
| HAVING   |      |              |
| ORDER BY |      |              |
| LIMIT    |      |              |

``` sql
select prod_name from products;  //检索单个列
select prod_id,prod_name,prod_price from products；  //检索多个列
select * from products //检索所有列
```

检索不同的行，使用`DISTINCT`关键字，必须直接放在列名的前面

``` sql 
select distinct  vend_id from products;
```

返回第一行或前几行，使用LIMIT子句。带一个值的LIMIT总是从第一行开始，给出的书为返回的行数。带两个值的LIMIT可以指定从行号为第一个值得位置开始。

``` sql
select prod_name from products limit 1;
select prod_name from products limit 7,5;
```

#### 排序检索数据

ORDER BY 子句取一个或多个列得名字，据此对输出进行排序.在按多个列排序时，排序完全按所规定得顺序进行，只有第一个排序序列值不唯一时，才会使用之后得排序序列。

``` mysql
select prod_name from products order by prod_name;
select prod_id,prod_price,prod_name from products order by prod_price,prod_name;
```

数据排序默认为升序(ASC)，降序排序须指定`DESC`关键字。DESC 关键字只应用到直接位于其前面得列名。

``` mysql
select prod_id,pro_price from products order by prod_price desc,prod_name;
```

#### 过滤数据

在select语句中，数据根据WHERE子句中指定得搜索条件进行过滤。在同时使用ORDER BY 和 WHERE子句时，应该让ORDER BY 位于 WHERE之后，否则将会产生错误。

``` mysql
select prod_name,prod_price from products where prod_price = 2.50;
```

| 操作符              | 说明                                        |
| ------------------- | ------------------------------------------- |
| =                   |                                             |
| <>                  | 不等于                                      |
| !=                  | 不等于                                      |
| <                   |                                             |
| <=                  |                                             |
| >                   |                                             |
| >=                  |                                             |
| BETWEEN ... AND ... | 在指定得两个值之间,包括指定得开始值和结束值 |

##### 空值检查

在一个列不包含值时，称其为包含控制NULL.

``` mysql
select prod_name from products where prod_price is null;
```

##### 组合WHERE 子句 

以AND子句的方式或OR子句的方式使用

``` mysql
select prod_id,prod_price,prod_name frome products where vend_id = 1003 and prod_price <= 10;
select prod_name,prod_price from products where vend_id =1002 or vend_id =1003;
```

##### IN 操作符

IN 操作符用来指定条件范围，范围中的每个条件都可以进行匹配。IN取合法值的由逗号分隔的清单，全都括在圆括号。

``` mysql
select prod_name,prod_price from products where vend_id in (1002,1003) order by prod_name;
```

IN 与 OR具有相同的功能，相比优点：

* 在使用长的合法选项清单时，IN操作符的语法更清楚且更直观。
* 在使用IN时，计算的次序更容易管理
* IN 操作符一般比OR操作符清单执行更快
* IN的最大优点是可以包含其他select语句，使得能够更动态地建立WHERE子句。

##### NOT 操作符

否定它之后所跟地任何条件

``` mysql
select prod_name ,prod_price from products where vend_id NOT IN (1002,1003) order by prod_name;
```

##### LIKE 操作符

使用通配符进行搜索

%通配符：%表示任何字符出现任意次数

_通配符：匹配单个字符,不能多也不能少

``` mysql
select prod_name from products where prod_name LIKE 's%e';
select prod_id,prod_name from products where prod_name like '_ ton anvil';
```

#### 计算字段

``` mysql
select concat(Rtrim(vend_name),'(',RTrim(vend_country),')') as vend_title from vendors order by vend_name;

+------------------------+
| vend_title             |
+------------------------+
| ACME(USA)              |
| Anvils R Us(USA)       |
| Furball Inc.(USA)      |
| Jet Set(England)       |
| Jouets Et Ours(France) |
| LT Supplies(USA)       |
+------------------------+

select prod_id,quantity,item_price,quantity*item_price as expanded_price from orderitems where order_num = 20005;

+---------+----------+------------+----------------+
| prod_id | quantity | item_price | expanded_price |
+---------+----------+------------+----------------+
| ANV01   |       10 |       5.99 |          59.90 |
| ANV02   |        3 |       9.99 |          29.97 |
| TNT2    |        5 |      10.00 |          50.00 |
| FB      |        1 |      10.00 |          10.00 |
+---------+----------+------------+----------------+
```

#### 聚集函数

运行在行组上，计算和返回单个值的函数。

| 函数    | 说明             |
| ------- | ---------------- |
| AVG()   | 返回某列的平均值 |
| COUNT() | 返回某列的函数   |
| MAX()   | 返回某列的最大值 |
| MIN()   | 返回某列的最小值 |
| SUM()   | 返回某列值之和   |

* avg() 只能用来确定单个列的平均值，并且忽略值位NULL的行

``` mysql
select AVG(prod_price) as avg_price from products;
select AVG(prod_price) as avg_price from products where vend_id = 1003;
```

* count()有两种使用方式
  * 使用`count(*)`对表中行的数目进行计数，不管表列中包含的是空值还是非空值。
  * 使用`COUNT(column)`对特定列中具有值得行进行计数，忽略NULL值
* Max（）返回指定列中得最大值，忽略列值为NULL的行。

``` mysql
select max(prod_price) as max_price from products;
```

* MIN()，返回指定列的最小值。

* SUM函数，返回指定列值的和（总计）

#### 数据分组

分组是在SELECT语句的GROUP BY子句中建立的。

* 除聚集计算语句外，SELECT语句中的每个列都必须在group by 子句中给出。
* 如果分组列中具有NULL值，则NULL将作为一个分组返回。如果列中有多行NULL值，它们将分为一组。
* GROUP BY子句必须出现在WHERE 子句之后，ORDER BY子句之前。

``` mysql
select vend_id,count(*) AS num_prods from products group by vend_id;

+---------+-----------+
| vend_id | num_prods |
+---------+-----------+
|    1001 |         3 |
|    1002 |         2 |
|    1003 |         7 |
|    1005 |         2 |
+---------+-----------+
```

#### 过滤分组

having 操作符

``` mysql
select cust_id,count(*) ass orders from orders group by cust_id having count(*)>=2;
```

#### 组合查询

UNION

* UNION必须由两条或两条以上的SELECT语句组成，语句之间用关键字UNION隔离



### Insert语句

#### 插入完整的行

省略列必须满足以下某个条件：

* 该列定义为允许NULL值
* 在表定义中给出默认值

``` mysql
insert into Customers values (NULL,'pep E.lapew','100 main street','Los angeles','ca','90046','USA',NULL,NULL);

//更安全的方法
insert into customers(cust_name,cust_address,cust_city,cust_state,cust_zip,cust_country,cust_contact,cust_email)values('pep E.lapew','100 main street','Los angeles','ca','90046','USA',NULL,NULL)
```

#### 插入检索出的数据

insert + select

``` mysql
insert into customers(cust_id,cust_contact,cust_email,cust_name,cust_address,cust_city,_cust_state,cust_zip,cust_country) select cust_id,cust_contact,cust_email,cust_name,cust_address,cust_city,_cust_state,cust_zip,cust_country from custnew;
```

### update 语句

``` mysql
update customers set cust_email = 'esfsdfsdf@asd.com' where cust_id = 10005;
```

### delete 语句

``` mysql
delete from customers where cust_id = 10006;
```



## 联结

### 内部联结

### 自联结

### 自然联结

### 外部联结

## 存储过程

简单来说就是为以后的使用而保存的一条或多条MYSQL的集合。3个主要的好处简单、安全、高性能。

### 执行存储过程

``` mysql
call productpricing(@pricelow,@pricehigh,@priceaverage);
```

### 创建存储过程

``` mysql
create procedure productpricing()
begin 
	select avg(prod_price) as priceaverage
	from products;
end;
```

### 删除存储过程

``` mysql
drop procedure productpricing; //仅当存在时删除，如果不存在返回错误

drop procedure if exists
```

## 事务处理（transaction processing）

* 事务（transaction）指一组SQL语句
* 回退（rollback）指撤销指定SQL语句的过程
* 提交（commit）指将为存储的SQL语句结果写入数据库表
* 保留点（save point）指事务处理中设置的临时占位符（place-holder），可以对它发布回退（与回退整个事务处理不同）

事务处理是一种的机制，用来管理必须成批执行的mysql操作，以保证数据库不包含不完整的操作结果。利用事务处理，可以保证一组操作不会中途停止，它们或者作为整体执行，或者完全不执行（除非明确指示）。如果没有错误发生，整组语句提交给（写到）数据库表。如果发生错误，则进行回退（撤销）以恢复数据库到某个已知且安全的状态。

### 事务的隔离级别

| 隔离级别                     | 脏读（Dirty Read） | 不可重复读(Non repeatable read) | 幻读(Phantom Read) |
| ---------------------------- | ------------------ | ------------------------------- | ------------------ |
| 未提交读（Read uncommitted） | 可能               | 可能                            | 可能               |
| 已提交读（Read committed）   | 不可能             | 可能                            | 可能               |
| 可重复读（Repeatable read）  | 不可能             | 不可能                          | 可能               |
| 可串行化（Serializable）     | 不可能             | 不可能                          | 可能               |



## 索引

便捷化检索表中行和列的子集，而不需要检测表中的每行。

#### 唯一索引

#### 组合索引

#### 索引类型

##### B树索引

##### 位图索引

##### 文本索引

#### 索引失效的情况

* 条件中用or，即使其中有条件带索引，也不会使用索引查询
* 对于多列索引，不是使用的第一部分，则不会使用索引
* like的模糊查询以%开头，索引失效
* 如果列类型是字符串，那一定要在条件中将数据使用引号引用起来，否则不会使用索引
* 如果mysql预计使用全表扫描比使用索引快，则不使用索引

#### 查看索引使用情况

## innodb 索引

### 底层原理

## 数据库调优方式



