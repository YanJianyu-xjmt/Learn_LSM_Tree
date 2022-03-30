package datastruct

import (
	"testing"
)

func TestLru1(t *testing.T) {
	lru := NewLRU(2)

	ak := []byte("a1")
	av := []byte("a2")
	bk := []byte("b1")
	bv := []byte("b2")
	ck := []byte("c1")
	cv := []byte("c2")

	lru.Insert(ak, av)
	lru.Insert(bk, bv)

	v, err := lru.Find(ak)
	vs := string(v)

	if err != nil || vs != "a2" {
		t.Errorf("no found")
	}

	v, err = lru.Find(bk)
	vs = string(v)

	if err != nil || vs != "b2" {
		t.Errorf("no found")
	}

	// 测试

	lru.Insert(ck, cv)

	v, err = lru.Find(ck)
	vs = string(v)

	if err != nil || vs != "c2" {
		t.Errorf("no found %s", vs)
	}

	v, err = lru.Find(ak)
	vs = string(v)

	if err == nil {
		t.Errorf("no found %s,%v", vs, err)
	}

}

// 测试查询会不会更新
func TestLru2(t *testing.T) {
	lru := NewLRU(2)

	ak := []byte("a1")
	av := []byte("a2")
	bk := []byte("b1")
	bv := []byte("b2")
	ck := []byte("c1")
	cv := []byte("c2")

	lru.Insert(ak, av)
	lru.Insert(bk, bv)

	v, err := lru.Find(ak)
	vs := string(v)

	if err != nil || vs != "a2" {
		t.Errorf("no found")
	}

	lru.Insert(ck, cv)

	v, err = lru.Find(bk)
	vs = string(v)

	if err == nil || len(vs) > 0 {
		t.Errorf("no found")
	}

	v, err = lru.Find(ak)
	vs = string(v)

	if err != nil || vs != "a2" {
		t.Errorf("no found")
	}

	v, err = lru.Find(ck)
	vs = string(v)

	if err != nil || vs != "c2" {
		t.Errorf("no found")
	}
}
