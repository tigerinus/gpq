package gpq_test

import (
	"testing"

	"github.com/tigerinus/gpq"
)

// This is an example struct that is irrelevant to gpq
type Expirable[T any] struct {
	Data           T
	ExpirationTime int64
}

func TestPriorityQueue(t *testing.T) {
	pq := gpq.NewPriorityQueue[*Expirable[string]](func(i, j *Expirable[string]) bool {
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

	pq.Push(&e2)
	{
		e := pq.Peek()
		if e.Data != e2.Data {
			t.Fail()
		}

		if e.ExpirationTime != e2.ExpirationTime {
			t.Fail()
		}
	}

	pq.Push(&e1)
	{
		e := pq.Peek()
		if e.Data != e2.Data {
			t.Fail()
		}

		if e.ExpirationTime != e2.ExpirationTime {
			t.Fail()
		}
	}

	pq.Push(&e3)
	{
		e := pq.Peek()
		if e.Data != e3.Data {
			t.Fail()
		}

		if e.ExpirationTime != e3.ExpirationTime {
			t.Fail()
		}
	}

	for i := 3; i > 0; i-- {
		if i != pq.Len() {
			t.Fail()
		}

		e := pq.Pop()

		if e.ExpirationTime != int64(i) {
			t.Fail()
		}
	}
}
