package main

type LRUCache struct {
	size int
	cache map[int]*LinkNode
	capacity int
	head, tail *LinkNode
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		cache: map[int]*LinkNode{},
		head: initLinkNode(0,0),
		tail: initLinkNode(0,0),
		capacity: capacity,
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}


func (this *LRUCache) Get(key int) int {
	if _,ok := (*this).cache[key]; !ok {
		return -1
	}
	node := (*this).cache[key]
	this.moveToHead(node)
	return (*node).value
}


func (this *LRUCache) Put(key int, value int)  {
	if _,ok := (*this).cache[key]; !ok {
		node := initLinkNode(key, value)
		this.cache[key] = node
		this.addToHead(node)
		this.size++
		if this.size > this.capacity {
			node := this.removeTail()
			delete(this.cache, node.key)
			this.size--
		}

	} else {
		node := this.cache[key]
		node.value = value
		this.moveToHead(node)
	}
}
type LinkNode struct {
	key, value int
	prev, next *LinkNode
}

func (this *LRUCache) removeNode(node *LinkNode)  {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) addToHead(node *LinkNode)  {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}
func (this *LRUCache) moveToHead(node *LinkNode) {
	this.removeNode(node)
	this.addToHead(node)
}
func (this *LRUCache) removeTail() *LinkNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}

func initLinkNode(key, value int) *LinkNode {
	return &LinkNode{key: key, value: value}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */