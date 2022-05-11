package linkedliststack

import (
	"strings"

	"github.com/mneumi/godsa/common"
	"github.com/mneumi/godsa/linkedlist"
)

type LinkedListStack[T common.Element[T]] struct {
	linkedList *linkedlist.LinkedList[T]
}

func New[T common.Element[T]]() *LinkedListStack[T] {
	return &LinkedListStack[T]{
		linkedList: linkedlist.New[T](),
	}
}

func (l *LinkedListStack[T]) Size() int {
	return l.linkedList.Size()
}

func (l *LinkedListStack[T]) IsEmpty() bool {
	return l.linkedList.IsEmpty()
}

func (l *LinkedListStack[T]) Push(e T) {
	l.linkedList.AddFirst(e)
}

func (l *LinkedListStack[T]) Pop() (T, error) {
	return l.linkedList.RemoveFirst()
}

func (l *LinkedListStack[T]) Peek() (T, error) {
	return l.linkedList.GetFirst()
}

func (l *LinkedListStack[T]) String() string {
	sbu := strings.Builder{}

	sbu.WriteString("Stack: top ")
	sbu.WriteString(l.linkedList.String())

	return sbu.String()
}
