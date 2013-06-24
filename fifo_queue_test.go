//
// Created by Yaz Saito on 06/15/12.
//

package fifo_queue

import "testing"
import "math/rand"

func TestAssert(t *testing.T, b bool, objs ...interface{}) {
	if !b {
		t.Fatal(objs...)
	}
}

func TestBasic(t *testing.T) {
	q := NewQueue()
	TestAssert(t, q.Len() == 0)
	q.PushBack(10)
	TestAssert(t, q.Len() == 1)
	TestAssert(t, q.PopFront().(int) == 10)
	TestAssert(t, q.Len() == 0)
	q.PushBack(11)
	TestAssert(t, q.Len() == 1)
	TestAssert(t, q.PopFront().(int) == 11)
	TestAssert(t, q.Len() == 0)
}

func TestRandomized(t *testing.T) {
	var first, last int
	q := NewQueue()
	for i := 0; i < 10000; i++ {
		if rand.Intn(2) == 0 {
			count := rand.Intn(128)
			for j := 0; j < count; j++ {
				q.PushBack(last)
				last++
			}
		} else {
			count := rand.Intn(128)
			if count > (last - first) {
				count = last - first
			}
			for i := 0; i < count; i++ {
				TestAssert(t, q.Len() > 0, "len==0", q.Len())
				TestAssert(t, q.PopFront().(int) == first)
				first++
			}
		}
	}
}
