package gpq

import (
	"container/heap"

	"github.com/tigerinus/gpq/internal"
)

type PriorityQueue[T any] struct {
	ipq internal.PriorityQueue[T]
}

func (pq PriorityQueue[T]) Len() int {
	return pq.ipq.Len()
}

func (pq PriorityQueue[T]) Peek() *T {
	item := pq.ipq.Peek()

	if item == nil {
		return nil
	}

	typed := item.(T)
	return &typed
}

func (pq *PriorityQueue[T]) Push(item T) {
	heap.Push(&pq.ipq, item)
}

func (pq *PriorityQueue[T]) Pop() *T {
	item := heap.Pop(&pq.ipq)

	if item == nil {
		return nil
	}

	typed := item.(T)
	return &typed
}

func NewPriorityQueue[T any](less func(i, j T) bool) PriorityQueue[T] {
	return PriorityQueue[T]{
		ipq: internal.NewPriorityQueue[T](less),
	}
}
