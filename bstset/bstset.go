package bstset

import (
	"github.com/mneumi/godsa/bst"
	"github.com/mneumi/godsa/common"
)

type BSTSet[T common.Element[T]] struct {
	bst *bst.BST[T]
}

func New[T common.Element[T]]() *BSTSet[T] {
	return &BSTSet[T]{
		bst: bst.New[T](),
	}
}

func (b *BSTSet[T]) GetSize() int {
	return b.bst.Size()
}

func (b *BSTSet[T]) IsEmpty() bool {
	return b.bst.IsEmpty()
}

func (b *BSTSet[T]) Add(e T) {
	b.bst.Add(e)
}

func (b *BSTSet[T]) Contains(e T) bool {
	return b.bst.Contains(e)
}

func (b *BSTSet[T]) Remove(e T) {
	b.bst.Remove(e)
}
