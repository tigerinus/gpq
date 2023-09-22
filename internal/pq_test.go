package internal_test

import (
	"container/heap"
	"testing"

	"github.com/tigerinus/gpq/internal"
)

type Expirable[T any] struct {
	Data           T
	ExpirationTime int64
}

func TestPriorityQueue(t *testing.T) {
	pq := internal.NewPriorityQueue[*Expirable[string]](func(i, j *Expirable[string]) bool {
		return i.ExpirationTime > j.ExpirationTime
	})

	e1 := Expirable[string]{
		Data:           "test1",
		ExpirationTime: 1,
	}

	e2 := Expirable[string]{
		Data:           "test2",
		ExpirationTime: 2,
	}

	e3 := Expirable[string]{
		Data:           "test3",
		ExpirationTime: 3,
	}

	heap.Push(&pq, &e2)
	heap.Push(&pq, &e1)
	heap.Push(&pq, &e3)

	for i := 3; i > 0; i-- {
		if i != pq.Len() {
			t.Fail()
		}

		item := heap.Pop(&pq)
		e, ok := item.(*Expirable[string])
		if !ok {
			t.Fail()
		}

		if e.ExpirationTime != int64(i) {
			t.Fail()
		}
	}
}
