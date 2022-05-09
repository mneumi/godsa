package bst

import (
	"math/rand"
	"testing"
)

type myInt int

func (m myInt) Less(n myInt) bool {
	return m < n
}

func (m myInt) Equal(n myInt) bool {
	return m == n
}

func (m myInt) Grater(n myInt) bool {
	return m > n
}

func TestRemoveMin(t *testing.T) {
	bst := New[myInt]()

	for i := 0; i < 10000; i++ {
		num := rand.Intn(100000)
		bst.Add(myInt(num))
	}

	var list []myInt
	for !bst.IsEmpty() {
		list = append(list, bst.RemoveMin())
	}

	for i := 1; i < len(list); i++ {
		if list[i-1] > list[i] {
			t.Fatalf("Error")
		}
	}

	t.Log("Pass")
}

func TestRemoveMax(t *testing.T) {
	bst := New[myInt]()

	for i := 0; i < 10000; i++ {
		num := rand.Intn(100000)
		bst.Add(myInt(num))
	}

	var list []myInt
	for !bst.IsEmpty() {
		list = append(list, bst.RemoveMax())
	}

	for i := 1; i < len(list); i++ {
		if list[i-1] > list[i] {
			t.Fatalf("Error")
		}
	}

	t.Log("Pass")
}
