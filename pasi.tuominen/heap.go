package main

import (
	"fmt"
)

type heap struct {
	storage []int
}

// insert adds element e to the heap
func (h *heap) insert(e int) {
	fmt.Printf("Inserted: %d\n", e)
	h.storage = Append(h.storage, e)
	h.bubbleUp(len(h.storage) - 1)
}

func (h *heap) bubbleUp(n int) {
	if n != 0 {
		p := n / 2
		if h.storage[p] < h.storage[n] {
			t := h.storage[p]
			h.storage[p] = h.storage[n]
			h.storage[n] = t
			h.bubbleUp(p)
		}
	}
}


// pop returns the top element from the heap
func (h *heap) pop() int {
	length := len(h.storage)
	if length == 0 {
		panic()
	}
	popValue := h.storage[0]
	if length == 1 {
		h.storage = make([]int, 0, 2)
	} else {
		h.storage[0] = h.storage[length-1]
		h.storage = h.storage[0 : length-1]
		h.bubbleDown(0)
	}
	return popValue

}


func (h *heap) bubbleDown(n int) {

	lli := n*2 + 1 //left leaf index
	rli := n*2 + 2 //right leaf index

	if len(h.storage)-1 < lli { //no leafs
		return
	}
	if len(h.storage)-1 < rli { // only left leaf
		if h.storage[n] < h.storage[lli] {

			h.storage[n], h.storage[lli] = h.storage[lli], h.storage[n]
			h.bubbleDown(lli)

		}
		return
	}
	// 2 leafs
	var bigger int
	if h.storage[lli] > h.storage[rli] {
		bigger = lli
	} else {
		bigger = rli
	}

	if h.storage[n] < h.storage[bigger] {

		h.storage[n], h.storage[bigger] = h.storage[bigger], h.storage[n]
		h.bubbleDown(bigger)
	}

}

// create takes the given slice and creates a heap type that uses the slice for storage
func create(s []int) *heap {
	h := new(heap)

	if len(s) > 0 {
		h.storage = make([]int, 0, 2*len(s))
	} else {
		h.storage = make([]int, 0, 4)
	}
	for _, value := range s {
		h.insert(value)
	}
	return h
}


func (h *heap) printsize() { fmt.Printf("size: %d/%d\n", len(h.storage), cap(h.storage)) }

func (h *heap) printvalues() {
	fmt.Printf("Values in heap-array: ")
	for i := 0; i < len(h.storage); i++ {
		fmt.Printf("%d ", h.storage[i])
	}
	fmt.Printf("\n")

}

func Append(slice []int, value int) []int {
	l := len(slice)

	if len(slice) == cap(slice) {
		newSlice := make([]int, l*2)
		for i, c := range slice {
			newSlice[i] = c
		}
		slice = newSlice
	}
	slice = slice[0 : l+1]
	slice[l] = value
	return slice
}


func main() {

	slice := []int{1}

	h := create(slice)

	h.printvalues()
	h.printsize()
	h.insert(4)
	h.insert(2)
	h.insert(53)
	h.insert(31)
	h.insert(32)
	h.insert(33)
	h.insert(6)
	h.insert(33)

	for len(h.storage) != 0 {
		fmt.Printf("\n")
		h.printvalues()
		h.printsize()
		fmt.Printf("Popped %d\n", h.pop())

	}

}
