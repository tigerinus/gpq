package internal

import (
	"container/heap"
	"sync"
)

type PriorityQueue[T any] struct {
	queue []T
	less  func(i, j T) bool
	mutex *sync.Mutex
}

func (pq PriorityQueue[T]) Len() int {
	return len(pq.queue)
}

func (pq PriorityQueue[T]) Peek() any {
	if pq.Len() == 0 {
		return nil
	}

	return (pq.queue)[0]
}

func (pq PriorityQueue[T]) Less(i, j int) bool {
	return pq.less((pq.queue)[i], (pq.queue)[j])
}

func (pq *PriorityQueue[T]) Pop() any {
	pq.mutex.Lock()
	defer pq.mutex.Unlock()

	if pq.Len() == 0 {
		return nil
	}

	item := (pq.queue)[pq.Len()-1]

	pq.queue = (pq.queue)[:pq.Len()-1]

	return item
}

func (pq *PriorityQueue[T]) Push(item any) {
	_item := item.(T)

	pq.mutex.Lock()
	defer pq.mutex.Unlock()

	pq.queue = append(pq.queue, _item)
}

func (pq *PriorityQueue[T]) Swap(i, j int) {
	pq.mutex.Lock()
	defer pq.mutex.Unlock()

	if i >= pq.Len() || j >= pq.Len() {
		return
	}
	(pq.queue)[i], (pq.queue)[j] = (pq.queue)[j], (pq.queue)[i]
}

func NewPriorityQueue[T any](less func(i, j T) bool) PriorityQueue[T] {
	pq := PriorityQueue[T]{
		queue: make([]T, 0),
		less:  less,
		mutex: &sync.Mutex{},
	}

	heap.Init(&pq)

	return pq
}
