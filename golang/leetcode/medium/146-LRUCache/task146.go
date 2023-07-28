package main

func main() {

}

type Node struct {
	Next  *Node
	Prev  *Node
	Value int
	Key   int
}

type LRUCache struct {
	Head  *Node
	Tail  *Node
	Cap   int
	Nodes map[int]*Node
}

func Constructor(capacity int) LRUCache {
	nodes := make(map[int]*Node)
	cache := LRUCache{
		Cap:   capacity,
		Nodes: nodes,
		Head:  &Node{},
		Tail:  &Node{},
	}

	cache.Head.Next, cache.Tail.Prev = cache.Tail, cache.Head

	return cache
}

func (this *LRUCache) remove(node *Node) {
	prev, next := node.Prev, node.Next

	prev.Next = next
	next.Prev = prev
}

func (this *LRUCache) insert(node *Node) {
	next, prev := this.Tail, this.Tail.Prev

	prev.Next = node
	next.Prev = node

	node.Prev = prev
	node.Next = next
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.Nodes[key]
	if !ok {
		return -1
	}

	this.remove(node)
	this.insert(node)
	this.Nodes[key] = node

	return node.Value
}

func (this *LRUCache) Put(key int, value int) {
	_, ok := this.Nodes[key]

	if ok {
		this.remove(this.Nodes[key])
	}

	node := &Node{Key: key, Value: value}
	this.insert(node)
	this.Nodes[key] = node

	if len(this.Nodes) > this.Cap {
		nodeToDelete := this.Head.Next
		this.remove(nodeToDelete)
		delete(this.Nodes, nodeToDelete.Key)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
