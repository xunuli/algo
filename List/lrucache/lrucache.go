package lrucache

// LRU缓存机制，采用双向链表和哈希表实现
// 哈希表实现缓存cache查询
// 双向链表实现LRU内存淘汰
type LRUCache struct {
	//使用量
	size int
	//总容量
	capacity int
	//缓存
	cache map[int]*DlinkedNode
	//双向链表的首尾哑结点
	head, tail *DlinkedNode
}

// 定义双向链表的结构
type DlinkedNode struct {
	//结点的值
	key, value int
	//前后指针
	prev, next *DlinkedNode
}

// 初始化双向链表的结点
func initNode(key, value int) *DlinkedNode {
	return &DlinkedNode{
		key:   key,
		value: value,
		prev:  nil,
		next:  nil,
	}
}

// 初始化LRU缓存
func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		size:     0,
		capacity: capacity,
		cache:    map[int]*DlinkedNode{},
		head:     initNode(0, 0),
		tail:     initNode(0, 0),
	}
	//链接双向链表
	lru.head.next = lru.tail
	lru.tail.prev = lru.head

	return lru
}

// 查询LRU缓存，查到则返回，并将结点移动至链表的头部，否则返回-1；
func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		//表示未查询到
		return -1
	} else {
		//查到
		node := this.cache[key]
		//将该节点移动至链表的头部
		this.moveToHead(node)
		return node.value
	}
}

// 更新缓存，如果存在则更新，并将其移动至链表头部；
// 不存在则插入至链表的头部，此时判断是否淘汰链表尾部的结点
func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; ok {
		//表示存在
		node := this.cache[key]
		node.value = value
		this.moveToHead(node)
	} else {
		//表示不存在
		newnode := initNode(key, value)
		//插入到链表的头部
		this.addToHead(newnode)
		//加入缓存
		this.cache[key] = newnode
		this.size++
		//判断是否淘汰
		if this.size > this.capacity {
			//需要淘汰尾部结点
			removenode := this.removeTail()
			//删除尾部结点的缓存
			delete(this.cache, removenode.key)
			this.size++
		}
	}
}

// 删除某个结点
func (this *LRUCache) removeNode(node *DlinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// 将结点移动至链表头部，先删除在插入到链表的头部
func (this *LRUCache) moveToHead(node *DlinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

// 将新节点插入至链表头部
func (this *LRUCache) addToHead(node *DlinkedNode) {
	this.head.next.prev = node
	node.next = this.head.next
	this.head.next = node
	node.prev = this.head
}

// 将尾部结点淘汰
func (this *LRUCache) removeTail() *DlinkedNode {
	removenode := this.tail.prev
	this.removeNode(removenode)
	return removenode
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
