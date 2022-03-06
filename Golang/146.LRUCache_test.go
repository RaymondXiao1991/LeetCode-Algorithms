package Golang

import (
	"testing"
)

func TestLRUCache(t *testing.T) {
	lruCache := Constructor(2)

	for _, test := range []struct {
		Key  int
		Val  int
		find int
		want int
	}{
		{1, 1, 1, 1},
		{2, 2, 2, 2},
		{3, 3, 2, 2},
		{4, 4, 1, -1},
		//{4, 4, 3, 3},
		//{4, 4, 4, 4},
	} {
		lruCache.Put(test.Key, test.Val)
		got := lruCache.Get(test.find)
		if got != test.want {
			t.Errorf("FAIL::LRUCache(%v, %v) = %d, want: %d", test.Key, test.Val, got, test.want)
		} else {
			t.Logf("SUCCESS::LRUCache(%v, %v) = %d", test.Key, test.Val, test.want)
		}
	}
	t.Logf("%v", lruCache.Get(3))
	t.Logf("%v", lruCache.Get(4))
}
