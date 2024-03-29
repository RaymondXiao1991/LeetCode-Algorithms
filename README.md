# LeetCode-Algorithms
The code I submitted in LeetCode.

## 一、数组

### 理论

> 数组是存放在连续内存空间上的相同类型数据的集合。

### 需要注意的是

1. 数组下标都是从0开始的。
2. 数组内存空间的地址是连续的
3. 数组的元素是不能删的，只能覆盖。

### 方法

1. 二分法
2. 双指针法
3. 滑动窗口
4. 模拟行为

## 二、链表

### 理论

链表是一种通过指针串联在一起的线性结构，每一个节点由两部分组成，一个是数据域一个是指针域（存放指向下一个节点的指针），最后一个节点的指针域指向null（空指针的意思）。

### 类型

1. 单链表
2. 双链表
  双链表：每一个节点有两个指针域，一个指向下一个节点，一个指向上一个节点。
  双链表 既可以向前查询也可以向后查询。
3. 循环链表
  链表首尾相连。

### 存储方式

链表在内存中不是连续分布的。
链表是通过指针域的指针链接在内存中各个节点。
所以链表中的节点在内存中不是连续分布的 ，而是散乱分布在内存中的某地址上，分配机制取决于操作系统的内存管理。

### 考点

1. 虚拟头结点的技巧
2. 链表的增删改查
3. 反转一个链表
4. 删除倒数第N个节点
5. 链表相交
6. 有否环形，以及环的入口

## 三、哈希表

### 理论

> 哈希表是根据关键码的值而直接进行访问的数据结构。

一般哈希表都是用来快速判断一个元素是否出现集合里。

### 常见的三种哈希结构

* 数组
* set （集合）
* map(映射)

### Array or Map?

什么时候用数组，什么时候用set。

## 四、字符串

### 理论

> 字符串是若干字符组成的有限序列，也可以理解为是一个字符数组

### 常用方法

1. 双指针法
2. 反转
3. KMP算法

## 五、栈与队列

### 栈

> 栈是以底层容器完成其所有的工作，对外提供统一的接口，底层容器是可插拔的（也就是说我们可以控制使用哪种容器来实现栈的功能）。
