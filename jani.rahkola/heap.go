package main

import (
	"fmt"
	"rand"
)

type heap struct {
	storage []int
}

func (h *heap) insert(e int) {
	if len(h.storage) == cap(h.storage) { // array full, cant insert
		return
	}
	fstFree := len(h.storage)
	h.storage = h.storage[0 : fstFree+1]
	h.storage[fstFree] = e
	h.upHeap(fstFree)
}

func (h *heap) upHeap(i int) {
	if i == 0 {
		return
	}
	parent := int((i - 1) / 2)
	if h.storage[i] >= h.storage[parent] {
		return
	}
	h.storage[i], h.storage[parent] = h.storage[parent], h.storage[i]
	h.upHeap(parent)
}

func (h *heap) pop() int {
	if len(h.storage) == 0 { // heap empty
		return -1
	}
	popped := h.storage[0]
	h.storage[0] = h.storage[len(h.storage)-1]
	h.storage = h.storage[0 : len(h.storage)-1]
	h.downHeap(0)
	return popped
}

func (h *heap) downHeap(i int) {
	leftIndex := 2*i + 1
	rightIndex := 2*i + 2
	l := len(h.storage)
	var smaller int

	if leftIndex >= l { // no left child
		if rightIndex >= l { // no childs
			return
		}
		if h.storage[i] <= h.storage[rightIndex] {
			return
		}
		smaller = rightIndex
	} else if rightIndex >= l { // no right child
		if h.storage[i] <= h.storage[leftIndex] {
			return
		}
		smaller = leftIndex
	} else { // both childs exist
		if h.storage[leftIndex] <= h.storage[rightIndex] {
			smaller = leftIndex
		} else {
			smaller = rightIndex
		}
	}

	if h.storage[i] > h.storage[smaller] {
		h.storage[i], h.storage[smaller] = h.storage[smaller], h.storage[i]
	}
	h.downHeap(smaller)
}

func create(s []int) *heap {
	s = s[0:0]
	return &heap{s}
}

func main() {
	s := make([]int, 50)
	heap := create(s)
	for i := 0; i < 50; i++ {
		heap.insert(rand.Int() % 31)
	}
	for _, _ = range s {
		fmt.Println(heap.pop())
	}
}
