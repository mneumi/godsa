package bst

import (
	"github.com/mneumi/godsa/common"
)

type node[T common.Element[T]] struct {
	element T
	left    *node[T]
	right   *node[T]
}

func newNode[T common.Element[T]](e T) *node[T] {
	return &node[T]{
		element: e,
		left:    nil,
		right:   nil,
	}
}

type BST[T common.Element[T]] struct {
	root *node[T]
	size int32
}

func New[T common.Element[T]]() *BST[T] {
	return &BST[T]{
		root: nil,
		size: 0,
	}
}

func (b *BST[T]) Size() int32 {
	return b.size
}

func (b *BST[T]) IsEmpty() bool {
	return b.size == 0
}

func (b *BST[T]) Add(e T) {
	b.root = b.add(b.root, e)
}

func (b *BST[T]) add(n *node[T], e T) *node[T] {
	if n == nil {
		b.size++
		return newNode(e)
	} else if e.Less(n.element) {
		n.left = b.add(n.left, e)
	} else if !e.Less(n.element) {
		n.right = b.add(n.right, e)
	}

	return n
}

func (b *BST[T]) Contains(e T) bool {
	return b.contains(b.root, e)
}

func (b *BST[T]) contains(n *node[T], e T) bool {
	if n == nil {
		return false
	}

	if n.element.Equal(e) {
		return true
	} else if n.element.Less(e) {
		return b.contains(n.left, e)
	} else {
		return b.contains(n.right, e)
	}
}

func (b *BST[T]) PreTraversal(fn func(e T)) {
	b.preTraversal(b.root, fn)
}

func (b *BST[T]) preTraversal(n *node[T], fn func(e T)) {
	if n == nil {
		return
	}

	fn(n.element)
	b.preTraversal(n.left, fn)
	b.preTraversal(n.right, fn)
}

func (b *BST[T]) InTraversal(fn func(e T)) {
	b.inTraversal(b.root, fn)
}

func (b *BST[T]) inTraversal(n *node[T], fn func(e T)) {
	if n == nil {
		return
	}

	b.inTraversal(n.left, fn)
	fn(n.element)
	b.inTraversal(n.right, fn)
}

func (b *BST[T]) PostTraversal(fn func(e T)) {
	b.postTraversal(b.root, fn)
}

func (b *BST[T]) postTraversal(n *node[T], fn func(e T)) {
	if n == nil {
		return
	}

	b.postTraversal(n.left, fn)
	b.postTraversal(n.right, fn)
	fn(n.element)
}

func (b *BST[T]) PreTraversalNR(fn func(e T)) {
	if b.root == nil {
		return
	}

	var stack []*node[T]
	stack = append(stack, b.root)

	for len(stack) != 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		fn(cur.element)

		if cur.right != nil {
			stack = append(stack, cur.right)
		}
		if cur.left != nil {
			stack = append(stack, cur.left)
		}
	}
}

func (b *BST[T]) InTraversalNR(fn func(e T)) {
	// TODO
}

func (b *BST[T]) PostTraversalNR(fn func(e T)) {
	// TODO
}

func (b *BST[T]) LevelTraversal(fn func(e T)) {
	if b.root == nil {
		return
	}

	var queue []*node[T]
	queue = append(queue, b.root)

	for len(queue) != 0 {
		cur := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		fn(cur.element)

		if cur.left != nil {
			queue = append(queue, cur.left)
		}
		if cur.right != nil {
			queue = append(queue, cur.right)
		}
	}
}
