package main

import "fmt"

// why not just type heap []int ?
type heap struct {
	data []int
}

// insert adds element e to the heap
func (h *heap) insert(e int) {
	if cap(h.data) == len(h.data) {		
		data := make([]int, cap(h.data)*2)[0:len(h.data)]
		copy(data, h.data)
		h.data = data
	}	

	h.data = h.data[0:len(h.data)+1]
	h.data[len(h.data)-1] = e

	this := len(h.data)-1

	for this > 0 && h.data[this] < h.data[up(this)] {
		h.data[this], h.data[up(this)] = h.data[up(this)], h.data[this]
	}
}


// pop returns the top element from the heap
func (h *heap) pop() int {
	val := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[0:len(h.data)-1]
	// if > 1 element left, maintain heap invariant
	if len(h.data) > 1 {
		this := 0
		min := -1
		for this != min {
			min = this
			if left(this) < len(h.data) && h.data[left(this)] < h.data[min] {
				min = left(this)
			}
			if right(this) < len(h.data) && h.data[right(this)] < h.data[min] {
				min = right(this)
			}
			h.data[min], h.data[this] = h.data[this], h.data[min]
		}
	}

	return val
}


// create takes the given slice and creates a heap type that uses the slice for storage
func create(s []int) *heap {
	h := new(heap)
	// could use ceil(log_2(len(s))) or something as size
	h.data = make([]int, len(s))[0:0]
	for _,v := range(s) {
		h.insert(v)
	}
	return h
}

func up(i int) int { return (i-1)/2 }
func left(i int) int { return i*2+1 }
func right(i int) int { return i*2+2 }

func main() {
	h := create([]int{99,5,2,3,11,9,21834})
	h.insert(33)

	fmt.Println(h)
	// try heap sorting
	for i := 0; i < 8; i++ {
		fmt.Println(h.pop())
	}
	fmt.Println(h)
}
