package datastruct

import (
	"fmt"
	"sync"
)

type VandPtr struct {
	Value []byte
	Node  *SkipListNode
}
type LRU struct {
	valueMap map[string]*ListNode
	size     int
	cap      int
	list     *List
	mu       sync.Mutex
}

func NewLRU(cap int) *LRU {
	l := &LRU{
		size: 0,
		cap:  cap,
	}

	l.list = NewList()
	l.valueMap = make(map[string]*ListNode)

	return l
}
func (l *LRU) Insert(key []byte, value []byte) {

	l.mu.Lock()
	defer l.mu.Unlock()

	kv := KV{
		key:   key,
		value: value,
	}
	keyString := string(key)

	if l.size < l.cap {
		l.list.PushHead(kv)
		tmpNode := l.list.head()

		l.valueMap[keyString] = tmpNode
		l.size++
		return
	}

	l.list.PushHead(kv)
	l.valueMap[keyString] = l.list.head()
	rkv, err := l.list.PopTail()

	if err != nil {
		panic("wrong logic in lru")
	}

	rKV, ok := rkv.(KV)
	if !ok {
		panic("wrong logic in lru")
	}

	keyString = string(rKV.key)
	delete(l.valueMap, keyString)
}

func (l *LRU) Find(key []byte) ([]byte, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	keyString := string(key)

	nodePtr, ok := l.valueMap[keyString]

	if !ok {
		return nil, fmt.Errorf("no found")
	}

	rkv := nodePtr.val
	rKV, ok := rkv.(KV)

	if !ok {
		return nil, fmt.Errorf("wrong kv struct")
	}

	l.list.PushHead(rKV)
	l.valueMap[keyString] = l.list.head()

	_, err := l.list.PopTail()

	if err != nil {
		panic("wrong logic in lru")
	}

	return rKV.value, nil
}

func (l *LRU) Delete(key []byte) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	keyString := string(key)

	nodePtr, ok := l.valueMap[keyString]
	if !ok {
		return fmt.Errorf("no found")
	}

	l.size--
	l.list.Delete(nodePtr)
	delete(l.valueMap, keyString)
	return nil
}
