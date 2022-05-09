package trie

type node struct {
	isWord bool
	next   map[rune]*node
}

func newNode() *node {
	return &node{
		isWord: false,
		next:   make(map[rune]*node),
	}
}

type Trie struct {
	root *node
	size int32
}

func New() *Trie {
	return &Trie{
		root: newNode(),
		size: 0,
	}
}

func (t *Trie) Size() int32 {
	return t.size
}

func (t *Trie) Add(word string) {
	cur := t.root

	for _, c := range word {
		if _, ok := cur.next[c]; !ok {
			cur.next[c] = newNode()
		}
		cur = cur.next[c]
	}

	if !cur.isWord {
		cur.isWord = true
		t.size++
	}
}

func (t *Trie) Contains(word string) bool {
	cur := t.root

	for _, c := range word {
		if _, ok := cur.next[c]; !ok {
			return false
		}
		cur = cur.next[c]
	}

	return cur.isWord
}

func (t *Trie) IsPrefix(prefix string) bool {
	cur := t.root

	for _, c := range prefix {
		if _, ok := cur.next[c]; !ok {
			return false
		}
		cur = cur.next[c]
	}

	return true
}

func (t *Trie) Remove(word string) {
	// TODO
}
