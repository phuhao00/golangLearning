package likou

//LRU 双向链表实现

type Node struct {
	key, value int
	pre, next  *Node
}

type LRUCache struct {
	capacity  int
	mapCache  map[int]*Node
	head, end *Node
}

func Constructor(capacity int) LRUCache {
	mapCache := make(map[int]*Node, capacity)
	return LRUCache{
		capacity: capacity,
		mapCache: mapCache,
	}
}
func (this *LRUCache) Get(key int) int {
	if node, exist := this.mapCache[key]; !exist {
		return -1
	} else {
		this.remove(node)
		this.setHeader(node)
		return node.value
	}
}

func (this *LRUCache) Put(key, value int) {
	if node, exist := this.mapCache[key]; exist {
		node.value = value
		this.remove(node)
		this.setHeader(node)
	} else {
		if len(this.mapCache) >= this.capacity {
			delete(this.mapCache, this.end.key)
			this.remove(this.end)
		}
		n := &Node{
			key:   key,
			value: value,
		}
		this.mapCache[key] = n
		this.setHeader(n)
	}
}

func (this *LRUCache) remove(node *Node) {
	if node.pre == nil {
		this.head = node.next
	} else {
		node.pre.next = node.next
	}
	if node.next == nil {
		this.end = node.pre
	} else {
		node.next.pre = node.pre
	}
}

func (this *LRUCache) setHeader(node *Node) {
	node.pre = nil
	if this.head == nil {
		this.head = node
	} else {
		this.head.pre, node.next = node, this.head
		this.head = node
	}
	if this.end == nil {
		this.end = node
	}
}
