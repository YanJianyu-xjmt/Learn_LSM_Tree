package src

import (
	"bytes"
	"math/rand"
	"sync"
	"time"
)

type SkipListNode struct {
	Key   []byte
	Value []byte

	NextIndexes []*SkipListNode
}

type SkipList struct {
	size        int
	byteSize    int
	Height      int
	HeadIndexes []*SkipListNode
	randSeed    *rand.Rand
	Mu          sync.Mutex
}

func NewSkipList() *SkipList {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	return &SkipList{
		size:     0,
		byteSize: 0,
		randSeed: r,
		Height:   1,
	}
}

func NewSkipListNode(key []byte, value []byte) *SkipListNode {

	return &SkipListNode{
		Key:   key,
		Value: value,
	}
}

func (s *SkipList) Insert(key []byte, value []byte) {

	tmpNode := NewSkipListNode(key, value)
	s.Mu.Lock()
	defer s.Mu.Unlock()

	height := s.getInsertHeight()
	tmpNode.NextIndexes = make([]*SkipListNode, height)
	for i := height; i > 0; i-- {

		if s.HeadIndexes[i-1] == nil {
			s.HeadIndexes[i-1] = tmpNode
			continue
		}

		if s.HeadIndexes[i-1].NextIndexes[i-1] == nil {
			s.HeadIndexes[i-1].NextIndexes[i-1] = tmpNode
			continue
		}

		pre := s.HeadIndexes[i-1]
		isBig := bytes.Compare(pre.Key, key)
		for pre != nil {
			switch isBig {
			case 0:
				pre.Value = value
				return
			case 1:
				s.HeadIndexes[i-1] = tmpNode
				tmpNode.NextIndexes[i-1] = pre
				break
			case -1:
				if pre.NextIndexes[i-1] == nil {
					pre.NextIndexes[i-1] = tmpNode
					break
				}

				isBig = bytes.Compare(pre.NextIndexes[i-1].Key, key)
				if isBig == 1 {
					tmpNode.NextIndexes[i-1] = pre.NextIndexes[i-1]
					pre.NextIndexes[i-1] = tmpNode
					break
				}
				pre = pre.NextIndexes[i-1]
			}
		}
	}
}

// must be used in insert within mutex Lock
func (s *SkipList) getInsertHeight() int {
	for i := 0; i < s.Height; i++ {
		num := s.randSeed.Int()

		if num%2 == 0 {
			return i + 1
		}
	}

	s.Height++
	s.HeadIndexes = append(s.HeadIndexes, nil)
	return s.Height
}
