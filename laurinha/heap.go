package main

import "fmt"

type heap struct {
	data []int
	size int
}

func parent(i int) int {
	return (i-1)/2
}


func left(i int) int {
	return i*2 + 1
}

func right(i int) int {
	return i*2 + 2
}

// moves data[i] upwards in the heap if needed
func reverseHeapify(data []int, i int) {
	p := parent(i)
	for i > 0 && data[p] < data[i] {
		data[p], data[i] = data[i], data[p]
		i = p
		p = parent(i)
	}
}

// moves h.data[i] downwards in the heap as long as needed
func (h *heap) heapify(i int) {
	l := left(i)
	r := right(i)
	smallest := i
	if l < h.size && h.data[l] < h.data[i] {
		smallest = l
	}
	if r < h.size && h.data[r] < h.data[smallest] {
		smallest = r
	}
	if smallest != i {
		h.data[i], h.data[smallest] = h.data[smallest], h.data[i]
		h.heapify(smallest)
	}
}

// inserts e into h if h has room
// returns true if e was inserted succesfully,
// otherwise false (there was no room for e)
func (h *heap) insert(e int) bool {
	if len(h.data) == h.size {
		return false
	}

	h.data[h.size] = e
	reverseHeapify(h.data, h.size)
	h.size++

	return true
}


// removes the smallest element from h and returns it
func (h *heap) pop() int {
	min := h.data[0]
	h.data[0] = h.data[h.size-1]
	h.size--
	h.heapify(0)

	return min
}

// returns true if h is empty, otherwise false
func (h *heap) empty() bool {
	return h.size == 0
}

//creates a new heap using s as storage
func create(s []int) *heap {
	h := heap{s, len(s)}
	for i := (h.size - 1)/2; i >= 0; i-- {
		h.heapify(i)
	}
	return &h
}

func main() {
	s := []int{17,1,34,4,2}
	fmt.Printf("slice: %v\n", s)
	h := create(s)
	val := h.pop()
	fmt.Printf("pop returned %v\n", val)
	h.insert(3)
	fmt.Printf("inserted 3\n")
	for !h.empty() {
		fmt.Printf("pop returned %v\n", h.pop())
	}
}

