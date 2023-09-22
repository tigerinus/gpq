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

func (pq *PriorityQueue[T]) Push(item T) {
	heap.Push(&pq.ipq, item)
}

func (pq *PriorityQueue[T]) Pop() T {
	item := heap.Pop(&pq.ipq)
	return item.(T)
}

func NewPriorityQueue[T any](less func(i, j T) bool) PriorityQueue[T] {
	return PriorityQueue[T]{
		ipq: internal.NewPriorityQueue[T](less),
	}
}