package main

import "fmt"

type heap struct {
	s []int
}


// insert adds element e to the heap
func (h *heap) insert(e int) {
	l := len(h.s)

	if l == cap(h.s) {
		auxs := make([]int, l*2)
		for i, elem := range h.s {
			auxs[i] = elem
		}
		h.s = auxs
	}
	h.s = h.s[0 : l+1]
	h.s[l] = e
	h.moveUp(len(h.s) - 1)
}

func (h *heap) moveUp(n int) {
	if n != 0 {
		r := n / 2
		if h.s[r] < h.s[n] {
			h.s[r], h.s[n] = h.s[n], h.s[r]
			h.moveUp(r)
		}
	}
}


// pop returns the top element from the heap
func (h *heap) pop() int {
	heaplength := len(h.s)
	top_elem := h.s[0]
	if heaplength != 0 {
		if heaplength == 1 {
			h.s = make([]int, 0, 2)
		} else {
			h.s[0] = h.s[heaplength-1]
			h.s = h.s[0 : heaplength-1]
			h.moveDown(0)
		}
	} else {
		return -1 // If heap is empty
	}

	return top_elem

}
func (h *heap) moveDown(n int) {

	leftChild := n*2 + 1
	rightChild := n*2 + 2

	if len(h.s)-1 < leftChild {
		return
	}
	if len(h.s)-1 < rightChild {
		if h.s[n] < h.s[leftChild] {
			h.s[n], h.s[leftChild] = h.s[leftChild], h.s[n]
			h.moveDown(leftChild)
		}
		return
	}
	var bigChild int
	if h.s[leftChild] > h.s[rightChild] {
		bigChild = leftChild
	} else {
		bigChild = rightChild
	}
	if h.s[n] < h.s[bigChild] {
		h.s[n], h.s[bigChild] = h.s[bigChild], h.s[n]
		h.moveDown(bigChild)
	}
}

// create takes the given slice and creates a heap type that uses the slice for storage
func create(s []int) *heap {
	h := new(heap)
	h.s = make([]int, 0, len(s))
	for _, e := range s {
		h.insert(e)
	}
	return h
}

func main() {

	s := []int{1, 3, 5, 6, 7}
	h := create(s)
	fmt.Printf("Heap : %v\n", h.s)
	fmt.Printf("Heap size: %d Capacity: %d\n", len(h.s), cap(h.s))

	h.insert(10)
	h.insert(15)
	h.insert(9)
	h.insert(50)

	fmt.Printf("Heap : %v\n", h.s)
	fmt.Printf("Heap size: %d Capacity: %d\n", len(h.s), cap(h.s))

	fmt.Printf("Popped : %v\n", h.pop())

	fmt.Printf("Heap : %v\n", h.s)
	fmt.Printf("Heap size: %d Capacity: %d\n", len(h.s), cap(h.s))

}
