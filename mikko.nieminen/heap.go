package main

import (
	"fmt"
	"os"
)

const (
	HEAP_MAX_SIZE int = 64
)

type heap struct {
//	data *[] int // pointer to a slice
	data [] int // a slice
}

func (h *heap) insert(e int) {
        if len(h.data)==cap(h.data) {
		fmt.Printf("stack overflow\n");
		os.Exit(1)
	}
	length := len(h.data)
	h.data = h.data [0:length+1] // extend the slice by one
	h.data[length] = e // add a new element to the bottom
	h.fixHeapFromBottom(length) // fix heap order (bottom to top)
}

func (h *heap) pop() int {
	if len(h.data) == 0 {
		fmt.Printf("stack underflow\n");
		os.Exit(1)
	}

	value := h.data[0] // value to be returned
	length := len(h.data)
	h.data[0] = h.data[length-1] // move last element to the top
	h.data = h.data [0:length-1] // shrink slice size by one
	h.fixHeapFromTop(0) // fix heap order (top to bottom)
	return value
}

func create(s []int) *heap {
	h := new(heap)
	h.data = s
	return h
}

func (h *heap) fixHeapFromBottom(index int) {
	// check if the element is on the top already
	if index == 0 {return} // nothing to sort

	parentIndex := (index - 1) / 2
	if h.data[parentIndex] < h.data[index] {
		// swap element up
		h.data[parentIndex], h.data[index] = h.data[index], h.data[parentIndex]

		// repeat recursively until correct position is found
		h.fixHeapFromBottom(parentIndex)
	}
}

func (h *heap) fixHeapFromTop(index int) {
	length := len(h.data)
	if index >= length-1 || length==0 {return} // nothing to sort
	leftChildIndex := 2*index+1
	rightChildIndex := 2*index+2

//	fmt.Printf("\ndata before swap=%v\n", h.data)
	if rightChildIndex<length { // both left and right children exist
		leftChildValue := h.data[leftChildIndex]
		rightChildValue := h.data[rightChildIndex]
		if h.data[index]>leftChildValue && h.data[index]>rightChildValue {
			return; // element is in correct position
		}
		if leftChildValue>rightChildValue {
//			fmt.Printf("swap to left (%d<-->%d)\n", h.data[index], h.data[leftChildIndex])
			h.data[index], h.data[leftChildIndex] = h.data[leftChildIndex], h.data[index]
			h.fixHeapFromTop(leftChildIndex)
		} else {
//			fmt.Printf("swap to right (%d<-->%d)\n", h.data[index], h.data[rightChildIndex])
			h.data[index], h.data[rightChildIndex] = h.data[rightChildIndex], h.data[index]
			h.fixHeapFromTop(rightChildIndex)
		}
	} else if leftChildIndex<length { // only left child exists
		if h.data[leftChildIndex]>h.data[index] {
//			fmt.Printf("swap to left (%d<-->%d)\n", h.data[index], h.data[leftChildIndex])
			h.data[index], h.data[leftChildIndex] = h.data[leftChildIndex], h.data[index]
			h.fixHeapFromTop(leftChildIndex)
		}
	} else {
		// element is in the bottom, do nothing
//		fmt.Println("element is in the bottom")
	}

}


func main() {
	slice := make([]int, 0, HEAP_MAX_SIZE)
	h := create(slice)
	h.insert(10)
	h.insert(12)
	h.insert(200)
	h.insert(190)
	h.insert(3)
	h.insert(1443)
	h.insert(23)
	h.insert(33)
	h.insert(400)
	h.insert(1000)
	h.insert(5)

	fmt.Printf("data=%v\n", h.data)
	for i := 0; len(h.data)>0; i++ {
		fmt.Printf("%d: pop=%d, data=%v\n", i, h.pop(), h.data)
	}
}
