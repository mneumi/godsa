package linkedlistset

import (
	"github.com/mneumi/godsa/common"
	"github.com/mneumi/godsa/linkedlist"
)

type LinkedListSet[T common.Element[T]] struct {
	list *linkedlist.LinkedList[T]
}

func New[T common.Element[T]]() *LinkedListSet[T] {
	return &LinkedListSet[T]{
		list: linkedlist.New[T](),
	}
}

func (l *LinkedListSet[T]) GetSize() int {
	return l.list.Size()
}

func (l *LinkedListSet[T]) IsEmtpy() bool {
	return l.list.IsEmpty()
}

func (l *LinkedListSet[T]) Contains(e T) bool {
	return l.list.Contains(e)
}

func (l *LinkedListSet[T]) Add(e T) {
	if !l.list.Contains(e) {
		l.list.AddFirst(e)
	}
}

func (l *LinkedListSet[T]) Remove(e T) {
	l.list.RemoveElement(e)
}
