package maxheap

import (
	"errors"

	"github.com/mneumi/godsa/common"
)

const (
	defaultCapacity = 100
)

type MaxHeap[T common.Element[T]] struct {
	data []T
}

func New[T common.Element[T]]() *MaxHeap[T] {
	return &MaxHeap[T]{
		data: make([]T, defaultCapacity),
	}
}

func (m *MaxHeap[T]) Size() int {
	return len(m.data)
}

func (m *MaxHeap[T]) IsEmpty() bool {
	return len(m.data) == 0
}

func (m *MaxHeap[T]) parent(index int) int {
	if index == 0 {
		return 0
	}
	return (index - 1) / 2
}

func (m *MaxHeap[T]) leftChild(index int) int {
	return index*2 + 1
}

func (m *MaxHeap[T]) rightChild(index int) int {
	return index*2 + 2
}

func (m *MaxHeap[T]) Add(e T) {
	m.data = append(m.data, e)
	m.siftUp(len(m.data) - 1)
}

func (m *MaxHeap[T]) siftUp(index int) {
	parentIndex := m.parent(index)

	for index > 0 && m.data[parentIndex].Less(m.data[index]) {
		swap(m.data, index, parentIndex)
		index = m.parent(index)
	}
}

func swap[T common.Element[T]](data []T, k int, j int) {
	if k < 0 || k >= len(data) || j < 0 || j >= len(data) {
		return
	}
	data[k], data[j] = data[j], data[k]
}

func (m *MaxHeap[T]) FindMax() (T, error) {
	var ret T
	if len(m.data) <= 0 {
		return ret, errors.New("can not find max when heap is empty")
	}
	return m.data[0], nil
}

func (m *MaxHeap[T]) ExtractMax() (T, error) {
	ret, err := m.FindMax()
	if err != nil {
		return ret, err
	}

	swap(m.data, 0, len(m.data)-1)
	m.data = m.data[:len(m.data)-1]
	m.siftDown(0)

	return ret, nil
}

func (m *MaxHeap[T]) siftDown(index int) {
	for m.leftChild(index) < len(m.data) {
		left := m.leftChild(index)
		if left+1 < len(m.data) && m.data[left+1].Grater(m.data[left]) {
			left = m.rightChild(index) // 此时，变量 `left`` 是子树根节点的值
		}
		// m.data[left] 是 leftChild 和 rightChild 中的最大值

		if m.data[index].Grater(m.data[left]) {
			break
		}

		swap(m.data, index, left)
		index = left
	}
}

func (m *MaxHeap[T]) Replace(e T) (T, error) {
	ret, err := m.FindMax()
	if err != nil {
		return ret, err
	}
	m.data[0] = e
	m.siftDown(0)
	return ret, err
}

func Heapify[T common.Element[T]](list []T) *MaxHeap[T] {
	ret := &MaxHeap[T]{
		data: list,
	}

	for i := ret.parent(len(ret.data) - 1); i >= 0; i-- {
		ret.siftDown(i)
	}

	return ret
}
