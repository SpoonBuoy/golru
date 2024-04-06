package lru

import (
	"fmt"
	"sync"
	"time"
)

type Node struct {
	Key       int
	Value     int
	CreatedAt time.Time
	Expiry    time.Duration
	Next      *Node
	Prev      *Node
}

func NewNode(key int, val int, expiry time.Duration) *Node {
	return &Node{
		Key:       key,
		Value:     val,
		Expiry:    expiry,
		CreatedAt: time.Now(),
	}
}

type DoubleLL struct {
	Head *Node
	Tail *Node
}

func NewDoubleLL() *DoubleLL {
	return &DoubleLL{}
}
func (dll *DoubleLL) AddAtFront(key int, val int, expiry time.Duration) {
	node := NewNode(key, val, expiry)
	if dll.Head == nil {
		//empty list
		dll.Head = node
		dll.Tail = node
		return
	}
	dll.Head.Prev = node
	node.Next = dll.Head
	dll.Head = node
}
func (dll *DoubleLL) AddAtBack(key int, val int, expiry time.Duration) {
	node := NewNode(key, val, expiry)
	if dll.Tail == nil {
		//empty
		dll.Head = node
		dll.Tail = node
		return
	}
	dll.Tail.Next = node
	node.Prev = dll.Tail
	dll.Tail = node
}
func (dll *DoubleLL) Delete(node *Node) {
	if node == nil {
		return
	}
	if node.Prev == nil && node.Next == nil {
		//only node exists
		//node
		dll.Head = nil
		dll.Tail = nil
		return
	}

	if node.Prev == nil {
		//node is head
		//node - next
		dll.Head = node.Next
		dll.Head.Prev = nil
		node = nil
		return
	}
	if node.Next == nil {
		//node is tail
		//prev - node
		dll.Tail = node.Prev
		dll.Tail.Next = nil
		node = nil
		return
	}
	//somewhere in between
	//prev - node - next
	//node.prev.next = node.next
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	node = nil
}

type LRU struct {
	Queue    DoubleLL
	Exists   map[int]*Node
	Mu       sync.RWMutex
	Capacity uint
	Count    uint
}

func NewLRU(size uint) *LRU {
	return &LRU{
		Queue:    *NewDoubleLL(),
		Exists:   make(map[int]*Node),
		Capacity: uint(size),
		Count:    0,
	}
}
func (lru *LRU) Get(key int) int {
	//check if the key exists
	lru.Mu.Lock()
	node, exists := lru.Exists[key]
	lru.Mu.Unlock()
	if !exists {
		//does not exist
		return -1
	}
	//we have to move it to top now
	lru.Queue.Delete(node)
	lru.Queue.AddAtFront(node.Key, node.Value, node.Expiry)
	lru.Mu.Lock()
	lru.Exists[key] = lru.Queue.Head
	lru.Mu.Unlock()
	return node.Value
}

func (lru *LRU) Set(key int, val int, expiry time.Duration) {
	lru.Mu.Lock()
	node, exists := lru.Exists[key]
	lru.Mu.Unlock()
	if !exists {
		if lru.Count < lru.Capacity {
			//can still hold more
			lru.Queue.AddAtFront(key, val, expiry)
			lru.Mu.Lock()
			lru.Exists[key] = lru.Queue.Head
			lru.Count++
			lru.Mu.Unlock()
		} else {
			//cant hold more
			tail := lru.Queue.Tail
			lru.Queue.Delete(tail)
			delete(lru.Exists, tail.Key)

			lru.Queue.AddAtFront(key, val, expiry)
			lru.Mu.Lock()
			lru.Exists[key] = lru.Queue.Head
			lru.Mu.Unlock()
			//count will reamin same since we deleted and added
		}
		return
	}
	//otherwise we update it
	lru.Queue.Delete(node)
	lru.Queue.AddAtFront(key, val, expiry)
	lru.Mu.Lock()
	lru.Exists[key] = lru.Queue.Head
	lru.Mu.Unlock()
}
func (lru *LRU) Print() {
	fmt.Printf("\n Printing LRU \n")
	head := lru.Queue.Head
	//tail := lru.Queue.Tail
	for head != nil {
		fmt.Printf("Key : %d, Val : %d, CreatedAt : %s, Expiry : %s \n", head.Key, head.Value, head.CreatedAt, head.Expiry)
		head = head.Next

	}
	fmt.Println()
}

type Entry struct {
	Key    int           `json:"key"`
	Value  int           `json:"value"`
	Expiry time.Duration `json:"expiry"`
}

func (lru *LRU) Top10() []Entry {
	head := lru.Queue.Head
	top10 := make([]Entry, 0)
	tot := 0
	for head != nil {
		if tot >= 10 {
			break
		}
		entry := Entry{
			Key:    head.Key,
			Value:  head.Value,
			Expiry: head.Expiry,
		}
		top10 = append(top10, entry)
		tot++
		head = head.Next
	}
	return top10
}

func (lru *LRU) CleanUpExpired() {
	go func() {
		for {
			//we clean after one second
			time.Sleep(time.Second)
			lru.Mu.Lock()
			for key, node := range lru.Exists {
				if time.Since(node.CreatedAt) > node.Expiry {
					fmt.Printf("Expired key %d val %d \n", node.Key, node.Key)
					//its expired, have to be deleted
					lru.Queue.Delete(node)
					delete(lru.Exists, key)
					lru.Count--
				}
			}
			lru.Mu.Unlock()
		}
	}()

}
