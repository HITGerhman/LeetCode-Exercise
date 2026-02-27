package main

type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

type LRUCache struct {
	capacity   int
	cache      map[int]*DLinkedNode
	head, tail *DLinkedNode
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		cache:    map[int]*DLinkedNode{},
		head:     &DLinkedNode{},
		tail:     &DLinkedNode{},
		capacity: capacity,
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (l *LRUCache) Get(key int) int {
	if node, ok := l.cache[key]; ok {
		l.moveToHead(node)
		return node.value
	}
	return -1
}

func (l *LRUCache) Put(key int, value int) {
	if node, ok := l.cache[key]; ok {
		node.value = value
		l.moveToHead(node)
	} else {
		node := &DLinkedNode{key: key, value: value}
		l.cache[key] = node
		l.addToHead(node)
		if len(l.cache) > l.capacity {
			removed := l.removeTail()
			delete(l.cache, removed.key)
		}
	}
}

func (l *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = l.head
	node.next = l.head.next
	l.head.next.prev = node
	l.head.next = node
}

func (l *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (l *LRUCache) moveToHead(node *DLinkedNode) {
	l.removeNode(node)
	l.addToHead(node)
}

func (l *LRUCache) removeTail() *DLinkedNode {
	node := l.tail.prev
	l.removeNode(node)
	return node
}
