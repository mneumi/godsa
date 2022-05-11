package linkedlist

import (
	"errors"
	"fmt"
	"strings"

	"github.com/mneumi/godsa/common"
)

type node[T common.Element[T]] struct {
	e    T
	next *node[T]
}

func newNode[T common.Element[T]](e T) *node[T] {
	return &node[T]{
		e:    e,
		next: nil,
	}
}

type LinkedList[T common.Element[T]] struct {
	dummyHead *node[T]
	size      int
}

func New[T common.Element[T]]() *LinkedList[T] {
	return &LinkedList[T]{
		dummyHead: newNode(*new(T)),
		size:      0,
	}
}

func (l *LinkedList[T]) Size() int {
	return l.size
}

func (l *LinkedList[T]) IsEmpty() bool {
	return l.size == 0
}

func (l *LinkedList[T]) Add(index int, e T) error {
	if index < 0 || index > l.size {
		return errors.New("add failed. illegal index")
	}

	prev := l.dummyHead
	for i := 0; i < index; i++ {
		prev = prev.next
	}

	n := newNode(e)
	n.next = prev.next
	prev.next = n
	l.size++

	return nil
}

func (l *LinkedList[T]) AddFirst(e T) {
	l.Add(0, e)
}

func (l *LinkedList[T]) AddLast(e T) {
	l.Add(l.size, e)
}

func (l *LinkedList[T]) Get(index int) (T, error) {
	var ret T
	if index < 0 || index > l.size {
		return ret, errors.New("add failed. illegal index")
	}

	cur := l.dummyHead.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}

	return cur.e, nil
}

func (l *LinkedList[T]) GetFirst() (T, error) {
	return l.Get(0)
}

func (l *LinkedList[T]) GetLast() (T, error) {
	return l.Get(l.size - 1)
}

func (l *LinkedList[T]) Set(index int, e T) error {
	if index < 0 || index > l.size {
		return errors.New("add failed. illegal index")
	}

	cur := l.dummyHead.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}

	cur.e = e

	return nil
}

func (l *LinkedList[T]) Contains(e T) bool {
	cur := l.dummyHead.next

	for cur.next != nil {
		if cur.e.Equal(e) {
			return true
		}
		cur = cur.next
	}

	return false
}

func (l *LinkedList[T]) String() string {
	sbu := strings.Builder{}

	cur := l.dummyHead.next

	for cur != nil {
		sbu.WriteString(fmt.Sprintf("%v", cur))
		cur = cur.next
	}
	sbu.WriteString("nil")

	return sbu.String()
}
