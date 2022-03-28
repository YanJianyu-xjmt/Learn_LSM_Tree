package datastruct

import "testing"

func TestListSize(t *testing.T) {
	a := "a"
	b := "b"
	c := "c"

	l := NewList()

	_, err := l.PopHead()
	if err == nil || l.GetSize() != 0 {
		t.Errorf("empty pop logic is wrong")
	}

	l.PushHead(a)

	if l.GetSize() != 1 {
		t.Errorf("wrong logic in count size")
	}

	l.PushHead(b)

	l.PrintList()
	l.PrintReList()
	y, err := l.PopTail()
	if u, ok := y.(string); err != nil || !ok || u != "a" {
		t.Errorf("wrong pop logic")
	}

	y, err = l.PopTail()
	if u, ok := y.(string); err != nil || !ok || u != "b" {
		t.Errorf("wrong pop logic")
	}

	_, err = l.PopTail()
	if err == nil || l.GetSize() != 0 {
		t.Errorf("wrong pop logic")
	}

	l.PushTail(a)
	l.PushHead(b)
	l.PushTail(c)

	// b->a->c
	l.PrintList()
	l.PrintReList()

	y, err = l.PopTail()
	if x, ok := y.(string); err != nil || !ok || x != "c" {
		t.Errorf("wrong pop tail")
	}

	y, err = l.PopHead()
	if x, ok := y.(string); err != nil || !ok || x != "b" {
		t.Errorf("wrong pop tail")
	}

	y, err = l.PopHead()
	if x, ok := y.(string); err != nil || !ok || x != "a" {
		t.Errorf("wrong pop tail")
	}
}
