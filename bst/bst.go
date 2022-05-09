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
	} else if e.Grater(n.element) {
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

func (b *BST[T]) Minimum() *node[T] {
	if b.root == nil {
		return nil
	}
	return b.minimum(b.root)
}

func (b *BST[T]) minimum(n *node[T]) *node[T] {
	if n.left == nil {
		return n
	}
	return b.minimum(n.left)
}

func (b *BST[T]) MinimumNR() *node[T] {
	if b.root == nil {
		return nil
	}

	cur := b.root
	for cur.left != nil {
		cur = cur.left
	}

	return cur
}

func (b *BST[T]) Maximum() *node[T] {
	if b.root == nil {
		return nil
	}
	return b.maximum(b.root)
}

func (b *BST[T]) maximum(n *node[T]) *node[T] {
	if n.right == nil {
		return n
	}
	return b.maximum(n.right)
}

func (b *BST[T]) MaximumNR() *node[T] {
	if b.root == nil {
		return nil
	}

	cur := b.root
	for cur.right != nil {
		cur = cur.right
	}

	return cur
}

func (b *BST[T]) RemoveMin() T {
	var ret T

	minNode := b.Minimum()
	if minNode == nil {
		return ret
	}

	ret = minNode.element

	b.root = b.removeMin(b.root)

	return ret
}

func (b *BST[T]) removeMin(n *node[T]) *node[T] {
	if n.left == nil {
		rightNode := n.right
		n.right = nil
		b.size--
		return rightNode
	}

	n.left = b.removeMin(n.left)
	return n
}

func (b *BST[T]) RemoveMax() T {
	var ret T

	maxNode := b.Maximum()
	if maxNode == nil {
		return ret
	}

	ret = maxNode.element

	b.root = b.removeMax(b.root)

	return ret
}

func (b *BST[T]) removeMax(n *node[T]) *node[T] {
	if n.right == nil {
		leftNode := n.left
		n.left = nil
		b.size--
		return leftNode
	}

	n.right = b.removeMax(n.right)
	return n
}

func (b *BST[T]) Remove(e T) {
	b.root = b.remove(b.root, e)
}

func (b *BST[T]) remove(n *node[T], e T) *node[T] {
	if n == nil {
		return nil
	}

	if e.Less(n.element) {
		n.left = b.remove(n.left, e)
		return n
	} else if e.Grater(n.element) {
		n.right = b.remove(n.right, e)
		return n
	} else {
		if n.left == nil {
			rightNode := n.right
			n.right = nil
			b.size--
			return rightNode
		} else if n.right == nil {
			leftNode := n.left
			n.left = nil
			b.size--
			return leftNode
		}

		successor := b.minimum(n.right)
		successor.right = b.removeMin(n.right)
		successor.left = n.left

		n.left = nil
		n.right = nil

		return successor
	}
}

func (b *BST[T]) Floor(e T) T {
	var ret T

	// TODO

	return ret
}

func (b *BST[T]) Ceil(e T) T {
	var ret T

	// TODO

	return ret
}

func (b *BST[T]) Rank(e T) int {
	// TODO

	return 0
}

func (b *BST[T]) Select(index int) T {
	var ret T

	// TODO

	return ret
}
