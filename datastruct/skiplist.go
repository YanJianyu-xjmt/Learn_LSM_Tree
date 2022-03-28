package datastruct

import (
	"bytes"
	"fmt"
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
		size:        0,
		byteSize:    0,
		randSeed:    r,
		Height:      1,
		HeadIndexes: []*SkipListNode{nil},
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
	fmt.Println(height)
	tmpNode.NextIndexes = make([]*SkipListNode, height)

	s.byteSize += (len(key) + len(value))

	for i := 0; i < height; i++ {
		if i > len(s.HeadIndexes) {
			s.HeadIndexes = append(s.HeadIndexes, tmpNode)
			continue
		}

		if s.HeadIndexes[i] == nil {
			s.HeadIndexes[i] = tmpNode
			continue
		}
		ptr := s.HeadIndexes[i]
		isBig := bytes.Compare(ptr.Key, key)

		if isBig == 0 {
			s.byteSize -= len(ptr.Value)
			ptr.Value = value
			return
		}

		if isBig == 1 {
			tmpNode.NextIndexes[i] = ptr
			s.HeadIndexes[i] = tmpNode
			continue
		}

		isOK := false
		for !isOK {
			isOK = false
			if ptr.NextIndexes[i] == nil {
				ptr.NextIndexes[i] = tmpNode
				isOK = true
			}
			nextPtr := ptr.NextIndexes[i]
			isNBig := bytes.Compare(nextPtr.Key, key)
			switch isNBig {
			case 0:
				s.byteSize -= len(nextPtr.Value)
				ptr.NextIndexes[i].Value = value
				isOK = true
			case 1:
				ptr.NextIndexes[i] = tmpNode
				tmpNode.NextIndexes[i] = nextPtr
				isOK = true
			case -1:
				ptr = nextPtr
			}
		}
	}

	s.size++
}

func (s *SkipList) Find(key []byte) ([]byte, error) {
	h := s.Height

	starth := h - 1

	for i := h - 1; i >= 0; i-- {
		if s.HeadIndexes[i] == nil {
			continue
		}

		isBig := bytes.Compare(s.HeadIndexes[i].Key, key)
		if isBig == 0 {
			return s.HeadIndexes[i].Value, nil
		}

		if isBig == 1 {
			starth = i
			break
		}
	}

	ptr := s.HeadIndexes[starth]

	for starth >= 0 {
		if bytes.Equal(ptr.Key, key) {
			return ptr.Value, nil
		}

		nextPtr := ptr.NextIndexes[starth]
		if nextPtr == nil && bytes.Compare(nextPtr.Key, key) == 1 {
			starth--
		}
	}
	return nil, fmt.Errorf("not found")
}

func (s *SkipList) Delete(key []byte) error {
	h := s.Height
	starth := h - 1

	for i := h - 1; i >= 0; i-- {
		if s.HeadIndexes[i] == nil {
			continue
		}

		isBig := bytes.Compare(s.HeadIndexes[i].Key, key)
		if isBig == 0 {
			s.byteSize -= (len(s.HeadIndexes[i].Key) + len(s.HeadIndexes[i].Value))
			s.size--
			return nil
		}

		if isBig == 1 {
			starth = i
			break
		}
	}

	ptr := s.HeadIndexes[starth]

	for starth >= 0 {
		if bytes.Equal(ptr.Key, key) {
			s.byteSize -= (len(ptr.Key) + len(ptr.Value))
			s.size--
			return nil
		}

		nextPtr := ptr.NextIndexes[starth]
		if nextPtr == nil && bytes.Compare(nextPtr.Key, key) == 1 {
			starth--
		}
	}
	return fmt.Errorf("not found")
}

// must be used in insert within mutex Lock
func (s *SkipList) getInsertHeight() int {
	if s.size == 0 {
		return 1
	}

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

func PrintSkipList(s *SkipList) {

	h := s.Height
	for i := 0; i < h; i++ {
		showstring := ""
		ptr := s.HeadIndexes[i]

		for ptr != nil {
			showstring += string(ptr.Key)
			showstring += "----->"
			ptr = ptr.NextIndexes[i]
		}

		fmt.Println(showstring)
	}
}
