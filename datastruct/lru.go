package datastruct

type VandPtr struct {
	Value []byte
	Node  *SkipListNode
}
type Lru struct {
	valueMap map[string]VandPtr
	size     int
	cap      int
	s        *SkipList
}
