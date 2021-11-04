package leetcode0146

// LRUCache struct
type LRUCache struct {
	// when size is bigger than capacity, we need to remove tail node
	size     int
	capacity int
	// to get node in O(1), we use hash map to store node
	cache map[int]*DLinkedNode
	// we only insert or delete node from head and tail, so keep the node here
	head, tail *DLinkedNode
}

// DLinkedNode struct
type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

func initLinkedNode(key, value int) *DLinkedNode {
	return &DLinkedNode{
		key:   key,
		value: value,
	}
}

// Constructor LRUCache
func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		cache:    map[int]*DLinkedNode{},
		head:     initLinkedNode(0, 0),
		tail:     initLinkedNode(0, 0),
		capacity: capacity,
	}
	lru.head.next = lru.tail
	lru.tail.prev = lru.head
	return lru
}

// Get LRUCache
func (lru *LRUCache) Get(key int) int {
	if _, ok := lru.cache[key]; !ok {
		return -1
	}
	node := lru.cache[key]
	lru.moveToHead(node)
	return node.value
}

// Put LRUCache
func (lru *LRUCache) Put(key, value int) {
	if _, ok := lru.cache[key]; !ok {
		node := initLinkedNode(key, value)
		lru.cache[key] = node
		lru.addToHead(node)
		lru.size++
		if lru.size > lru.capacity {
			lru.removeTail()
		}
	} else {
		node := lru.cache[key]
		node.value = value
		lru.moveToHead(node)
	}
}

func (lru *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = lru.head
	node.next = lru.head.next
	lru.head.next.prev = node
	lru.head.next = node
}

func (lru *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (lru *LRUCache) moveToHead(node *DLinkedNode) {
	lru.removeNode(node)
	lru.addToHead(node)
}

func (lru *LRUCache) removeTail() {
	node := lru.tail.prev
	lru.removeNode(node)
	delete(lru.cache, node.key)
	lru.size--
}
