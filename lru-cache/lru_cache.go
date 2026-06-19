package lrucache

import "sync"

type Node struct {
	Key   string
	Value any
	Prev  *Node
	Next  *Node
}

type LRUCache struct {
	Capacity int
	HeadNode *Node
	TailNode *Node
	HashMap  map[string]*Node
	mu       sync.Mutex
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		HeadNode: nil,
		TailNode: nil,
		HashMap:  make(map[string]*Node),
		Capacity: capacity,
	}
}

func (l *LRUCache) removeNode(node *Node) {
	if node.Prev == nil && node.Next == nil {
		l.HeadNode = nil
		l.TailNode = nil
	} else if node.Prev == nil {
		l.HeadNode = node.Next
		node.Next.Prev = nil
	} else if node.Next == nil {
		l.TailNode = node.Prev
		node.Prev.Next = nil
	} else {
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev
	}
	node.Prev = nil
	node.Next = nil
}

func (l *LRUCache) insertFront(node *Node) {
	if l.HeadNode == nil {
		l.HeadNode = node
		l.TailNode = node
		return
	}
	node.Next = l.HeadNode
	l.HeadNode.Prev = node
	l.HeadNode = node
	node.Prev = nil
}

func (l *LRUCache) Get(key string) any {
	l.mu.Lock()
	defer l.mu.Unlock()
	node, ok := l.HashMap[key]
	if !ok {
		return nil
	}

	if node.Prev == nil {
		return node.Value
	}
	l.removeNode(node)
	l.insertFront(node)
	return node.Value
}

func (l *LRUCache) Put(key string, value any) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if node, ok := l.HashMap[key]; ok {
		node.Value = value
		l.removeNode(node)
		l.insertFront(node)
		return
	}
	if len(l.HashMap) >= l.Capacity {
		tail := l.TailNode
		l.removeNode(tail)
		delete(l.HashMap, tail.Key)
	}

	node := &Node{Key: key, Value: value}
	l.insertFront(node)
	l.HashMap[key] = node
}
