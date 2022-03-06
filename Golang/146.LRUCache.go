/*
请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
实现 LRUCache 类：
    LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
    int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
    void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。

链接：https://leetcode-cn.com/problems/lru-cache

Note:
输入
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
输出
[null, null, null, 1, null, -1, null, -1, 3, 4]

解释
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // 缓存是 {1=1}
lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
lRUCache.get(1);    // 返回 1
lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
lRUCache.get(2);    // 返回 -1 (未找到)
lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
lRUCache.get(1);    // 返回 -1 (未找到)
lRUCache.get(3);    // 返回 3
lRUCache.get(4);    // 返回 4

*/

package Golang

// define Node
type Node struct {
	key, val   int   // 键、值
	prev, next *Node // 前、后
}

// init Node
func initNode(key, value int) *Node {
	return &Node{
		key: key,
		val: value,
	}
}

// double linked node
type DoubleLinkedList struct {
	head, tail *Node // 虚拟头节点、虚拟尾节点
	size       int   // 链表元素数
}

// 初始化双向链表的数据
func initDoubleLinkedList() *DoubleLinkedList {
	dll := &DoubleLinkedList{
		head: initNode(0, 0),
		tail: initNode(0, 0),
		size: 0,
	}
	dll.head.next = dll.tail
	dll.tail.prev = dll.head

	return dll
}

// 在链表尾部添加节点n，时间 O(1)
func (dll *DoubleLinkedList) addLast(n *Node) {
	n.prev = dll.tail.prev
	n.next = dll.tail

	dll.tail.prev.next = n
	dll.tail.prev = n
	dll.size++
}

// 删除链表中的n节点（n一定存在）
// 由于是双链表且给的是目标 Node 节点，时间 O(1)
func (dll *DoubleLinkedList) remove(n *Node) {
	n.prev.next = n.next
	n.next.prev = n.prev
	n.prev = nil
	n.next = nil
	dll.size--
}

// 删除链表中第一个节点，并返回该节点，时间 O(1)
func (dll *DoubleLinkedList) removeFirst() *Node {
	if dll.head.next == dll.tail {
		return nil
	}
	first := dll.head.next
	dll.remove(first)
	return first
}

// 返回链表长度，时间 O(1)
func (dll *DoubleLinkedList) getSize() int {
	return dll.size
}

// define LRU
type LRUCache struct {
	capacity int               // capacity
	knMap    map[int]*Node     // key -> Node(key, val)
	cache    *DoubleLinkedList // Node(k1, v1) <-> Node(k2, v2)...
}

// constructor LRU
func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		knMap:    map[int]*Node{},
		cache:    initDoubleLinkedList(),
	}
}

// 将某个 key 提升为最近使用的
func (lc *LRUCache) makeRecently(key int) {
	// 查找
	node := lc.knMap[key]
	// 先从链表中删除这个节点
	lc.cache.remove(node)
	// 重新插到队尾
	lc.cache.addLast(node)
}

// 添加最近使用的元素
func (lc *LRUCache) addRecently(key, val int) {
	n := initNode(key, val)
	// 链表尾部就是最近使用的元素
	lc.cache.addLast(n)
	// 在 map 中添加 key 的映射
	lc.knMap[key] = n
}

// 删除某一个 key
func (lc *LRUCache) deleteKey(key int) {
	n := lc.knMap[key]
	// 从链表中删除
	lc.cache.remove(n)
	// 从 map 中删除
	delete(lc.knMap, key)
}

// 删除最久未使用的元素
func (lc *LRUCache) removeLeastRecently() {
	// 链表头部的第一个元素就是最久未使用的
	deletedNode := lc.cache.removeFirst()
	// 同时从 map 中删除它的 key
	delete(lc.knMap, deletedNode.key)
}

func (lc *LRUCache) Get(key int) int {
	if node, exist := lc.knMap[key]; !exist {
		return -1
	} else {
		lc.makeRecently(key)
		return node.val
	}
}

func (lc *LRUCache) Put(key, val int) {
	if _, exist := lc.knMap[key]; exist {
		// 删除旧的数据
		lc.deleteKey(key)
		// 新插入的数据为最近使用的数据
		lc.addRecently(key, val)

		return
	}

	// 如果满了
	if lc.cache.size == lc.capacity {
		// 删除最久未使用的元素
		lc.removeLeastRecently()
	}
	// 添加为最近使用的元素
	lc.addRecently(key, val)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
