package lru

import (
	"testing"
	"time"
)

func TestDLL(t *testing.T) {
	dll := NewDoubleLL()
	exp := time.Second * 20
	for i := 0; i < 5; i++ {
		dll.AddAtFront(i, i, exp)
		dll.AddAtBack(i, i, exp)
	}

	head := dll.Head
	tail := dll.Tail
	for head != tail {
		if head.Value != tail.Value {
			t.Errorf("Fail")
			return
		}
		head = head.Next
		tail = tail.Prev
	}
}
func TestLRU(t *testing.T) {
	lru := NewLRU(6)
	exp := time.Minute * 5
	for i := 0; i < 15; i++ {
		lru.Set(i, i, exp)
		if i == 8 {
			lru.Get(5)
		}
		lru.Print()
	}
}
