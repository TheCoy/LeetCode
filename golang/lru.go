package golang

type LRUCache struct {
	size       int
	capactity  int
	hmap       map[int]*LruNode
	head, tail *LruNode
}

type LruNode struct {
	prev, next *LruNode
	key, value int
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		capactity: capacity,
		hmap:      map[int]*LruNode{},
	}
	l.head.next = l.tail
	l.tail.prev = l.head

	return l
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.hmap[key]; ok {
		this.moveToHead(node)
		return node.value
	}

	return -1
}

func (this *LRUCache) Set(key, value int) {
	if node, ok := this.hmap[key]; ok {
		node.value = value
		this.moveToHead(node)
		return
	}
	newNode := &LruNode{
		key:   key,
		value: value,
	}
	this.hmap[key] = newNode
	this.addToHead(newNode)
	this.size++
	if this.size > this.capactity {
		removed := this.removeTail()
		delete(this.hmap, removed.key)
		this.size--
	}

}

func (this *LRUCache) moveToHead(node *LruNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) addToHead(node *LruNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *LruNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) removeTail() *LruNode {
	node := this.tail.prev
	this.removeNode(node)

	return node
}
