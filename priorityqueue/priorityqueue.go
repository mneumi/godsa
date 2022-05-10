package priorityqueue

import (
	"github.com/mneumi/godsa/common"
	"github.com/mneumi/godsa/maxheap"
)

type PriorityQueue[T common.Element[T]] struct {
	maxHeap *maxheap.MaxHeap[T]
}

func New[T common.Element[T]]() *PriorityQueue[T] {
	return &PriorityQueue[T]{
		maxHeap: maxheap.New[T](),
	}
}

func (p *PriorityQueue[T]) Size() int {
	return p.maxHeap.Size()
}

func (p *PriorityQueue[T]) IsEmpty() bool {
	return p.maxHeap.IsEmpty()
}

func (p *PriorityQueue[T]) GetFront() (T, error) {
	return p.maxHeap.FindMax()
}

func (p *PriorityQueue[T]) Enqueue(e T) {
	p.maxHeap.Add(e)
}

func (p *PriorityQueue[T]) Dequeue() (T, error) {
	return p.maxHeap.ExtractMax()
}
