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

func (l *List) GetSize() int {
	return l.size
}

func (l *List) GetIterator() ListIterator {
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

func (l *List) PrintList() {

	var ptr *ListNode = l.Head
	for ptr != nil {
		fmt.Println(ptr.val)
		ptr = ptr.next
	}
}

func (l *List) PrintReList() {
	var ptr *ListNode = l.Tail
	for ptr != nil {
		fmt.Println(ptr.val)
		ptr = ptr.pre
	}
}

func (l *List) PushHead(value interface{}) {

	if l.size == 0 {
		tmpNode := NewListNode(value, nil, nil)
		l.Head = tmpNode
		l.Tail = tmpNode
		l.size++
		return
	}

	tmpNode := NewListNode(value, nil, l.Head)
	l.Head.pre = tmpNode
	l.size++
	l.Head = tmpNode
}

func (l *List) PushTail(value interface{}) {

	if l.size == 0 {
		tmpNode := NewListNode(value, nil, nil)
		l.Head = tmpNode
		l.Tail = tmpNode
		l.size++
		return
	}

	tmpNode := NewListNode(value, l.Tail, nil)
	l.Tail.next = tmpNode
	l.size++
	l.Tail = tmpNode
}

func (l *List) PopTail() (interface{}, error) {

	if l.size == 0 {
		return nil, fmt.Errorf("no element")
	}

	fmt.Println(l.size)
	if l.size == 1 {
		v := l.Head.val
		l.Head = nil
		l.Tail = nil
		l.size--
		return v, nil
	}

	v := l.Tail.val
	if l.Tail.pre == nil {
		fmt.Println("AAAAA")
	}
	l.Tail.pre.next = nil
	l.Tail = l.Tail.pre
	l.size--
	return v, nil
}

func (l *List) PopHead() (interface{}, error) {
	if l.size == 0 {
		return nil, fmt.Errorf("no element")
	}

	if l.size == 1 {
		v := l.Head.val
		l.Head = nil
		l.Tail = nil
		l.size--
		return v, nil
	}

	v := l.Head.val
	l.Head.next.pre = nil
	l.Head = l.Head.next
	l.size--
	return v, nil
}

func (l *List) Delete(n *ListNode) {
	if l.size == 0 {
		panic("list erase nil node")
	}

	if l.size == 1 {
		l.Head = nil
		l.Tail = nil
		l.size--
		return 
	}


	if l.Head == n{
		l.Head = n.next
	}
	if l.Tail == n{
		l.Tail = n.pre
	}

	if n.pre != nil{
		n.pre.next = n.next
	}
	if n.next != nil{
		n.next.pre = n.pre
	}
	
	l.size--
}