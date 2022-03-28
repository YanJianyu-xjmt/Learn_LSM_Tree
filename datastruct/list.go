package datastruct

import "fmt"

type ListNode struct {
	val  interface{}
	next *ListNode
	pre  *ListNode
}

func NewListNode(val interface{}, pre *ListNode, next *ListNode) *ListNode {
	return &ListNode{
		val:  val,
		next: next,
		pre:  pre,
	}
}

type ListIterator = *ListNode

func (l ListIterator) Next() (*ListNode, error) {

	if l == nil {
		return nil, fmt.Errorf("no next")
	}

	l = l.next
	return l, nil
}

func (l ListIterator) Pre() (*ListNode, error) {
	if l == nil {
		return nil, fmt.Errorf("no pre")
	}

	l = l.pre
	return l, nil
}

type List struct {
	Head     *ListNode
	Tail     *ListNode
	size     int
	capacity int
}

func (l *List) getIterator() ListIterator {
	return l.Head
}

func NewList() *List {
	return &List{
		Head:     nil,
		Tail:     nil,
		size:     0,
		capacity: 0,
	}
}
func (l *List) head() *ListNode {
	return l.Head
}

func (l *List) tail() *ListNode {
	return l.Tail
}

func (l *List) PushHead(value interface{}) {

	l.size++
	if l.Head == nil {
		tmpNode := NewListNode(value, nil, nil)
		l.Head = tmpNode
		l.Tail = tmpNode
		return
	}

	tmpNode := NewListNode(value, nil, l.Head)
	l.Head = tmpNode
}

func (l *List) PushTail(value interface{}) {

	l.size++
	if l.Tail == nil {
		tmpNode := NewListNode(value, nil, nil)
		l.Head = tmpNode
		l.Tail = tmpNode
		return
	}

	tmpNode := NewListNode(value, l.Tail, nil)
	l.Tail = tmpNode
}

func (l *List) PopTail() (interface{}, error) {

	if l.size == 0 {
		return nil, fmt.Errorf("no element")
	}

	l.size--
	if l.size == 1 {
		v := l.Head.val
		l.Head = nil
		l.Tail = nil
		return v, nil
	}

	v := l.Tail.val
	l.Tail.pre.next = nil
	l.Tail = nil
	return v, nil
}

func (l *List) PopHead() (interface{}, error) {
	if l.size == 0 {
		return nil, fmt.Errorf("no element")
	}

	l.size--
	if l.size == 1 {
		v := l.Head.val
		l.Head = nil
		l.Tail = nil
		return v, nil
	}

	v := l.Head.val
	l.Head.next.pre = nil
	l.Head = nil
	return v, nil
}
