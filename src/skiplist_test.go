package src

import (
	"testing"
)

func TestSkipListInsert(t *testing.T) {
	s := NewSkipList()

	a := "a"
	b := "b"
	c := "c"
	d := "d"

	s.Insert([]byte(c), []byte(d))
	s.Insert([]byte(a), []byte(b))

	PrintSkipList(s)

	t.Errorf("TTTT")

}
