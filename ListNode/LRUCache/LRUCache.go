package main

type DLinkedNode struct {
	key,value int
	prev, next*DLinkedNode
}

type LRUCache struct {
	size int
	capacity int
	cache map[int]*DLinkedNode
	//何意味
	head,tail *DLinkedNode
}

func Constructor(capacity int) LRUCache {
	l:=LRUCache {
		cache: map[int]*DLinkedNode{}
		head: &DLinkedNode{0,0,nil,nil}
		tail: &DLinkedNode{0,0,nil,nil}
		capacity: capacity,
		//????
	}
	l.head.next=l.tail
	l.tail.prev=l.head
	//wei shen me cheng huan
	return  l
}

func (this *LRUCache) Get(key int) int {
//liang ge can shu?????
	if node,ok:=this.cache[key];ok {
		this.moveToHead(node)
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node,ok:=this.cache[key];ok {
		node.value = value
		this.moveToHead(node)
	}else {
		node:=&DLinkedNode{key:key,value: value}
		this.cache[key] = node
		this.addToHead(node)
		this.size++
		if this.size>this.capacity {
			removed:=this.removeTail()
			delete(this.cache,removed.key)
			this.size--
		}
	}
}

func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *DLinkedNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}
