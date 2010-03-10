package main

import (
	"rand"
	"fmt"
	"sort"
	"os"
	"time"
)

type heap struct {
	heapSize  int    // Current size of the heap.
	heapSlice *[]int // Pointer to slice used for storage.
}

// Let's try out the heap with a little heap sorting.
func main() {
	const slicesize = 100000

	h := create(make([]int, slicesize))

	s := make([]int, slicesize)

	// Initialize s with random values.
	for i := 0; i < len(s); i++ {
		s[i] = rand.Int()
	}

	// Verify disorder.
	if sort.IntsAreSorted(s) {
		fmt.Println("Elements already in order, aborting.")
		os.Exit(1)
	}

	// Perform the sort and time it.
	startTime := time.Nanoseconds()
	h.heapsort(s)
	endTime := time.Nanoseconds()

	// Check that s is in fact sorted.
	if sort.IntsAreSorted(s) {
		fmt.Println("Sorted", slicesize, "elements succesfully in", (endTime-startTime)/1000000, "ms.")
	} else {
		fmt.Println("Sort failed.")
	}
}

// heapsorts sorts s using heap h.
func (h *heap) heapsort(s []int) {
	for _, e := range s {
		h.insert(e)
	}
	for i := len(s) - 1; i >= 0; i-- {
		s[i] = h.pop()
	}
}

// Returns index of parent of element at index i.
func (h *heap) parent(i int) int { return i / 2 }

// Returns index of left child of element at index i.
func (h *heap) left(i int) int {
	if 2*i > h.heapSize {
		return h.heapSize + 1
	}
	return 2 * i
}

// Returns index of right child of element at index i.
func (h *heap) right(i int) int {
	if 2*i+1 > h.heapSize {
		return h.heapSize + 1
	}
	return 2*i + 1
}

// Maintains heap property for node at index i and it's children.
func (h *heap) heapify(i int) {
	var next int
	leftChild := h.left(i)
	rightChild := h.right(i)
	if leftChild <= h.heapSize && (*h.heapSlice)[leftChild] > (*h.heapSlice)[i] {
		next = leftChild
	} else {
		next = i
	}
	if rightChild <= h.heapSize-1 && (*h.heapSlice)[rightChild] > (*h.heapSlice)[next] {
		next = rightChild
	}
	if next != i {
		(*h.heapSlice)[i], (*h.heapSlice)[next] = (*h.heapSlice)[next], (*h.heapSlice)[i]
		h.heapify(next)
	}
}

// insert adds element e to the heap.
func (h *heap) insert(e int) {
	if h.heapSize <= len(*h.heapSlice) {
		var i = h.heapSize
		h.heapSize++
		for i > 0 && (*h.heapSlice)[h.parent(i)] < e {
			(*h.heapSlice)[i] = (*h.heapSlice)[h.parent(i)]
			i = h.parent(i)
		}
		(*h.heapSlice)[i] = e
	}
}

// pop returns the top element from the heap.
func (h *heap) pop() int {
	max := (*h.heapSlice)[0]
	(*h.heapSlice)[0] = (*h.heapSlice)[h.heapSize-1]
	h.heapSize--
	h.heapify(0)
	return max
}

// create takes the given slice and creates a heap type that uses the slice for storage.
func create(s []int) *heap { return &heap{0, &s} }
