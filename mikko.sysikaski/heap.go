package main

import "fmt"

// the heap type
type heap []int

// insert adds element e to the heap
func (h *heap) insert(e int) {
	l := len(*h)
	*h = (*h)[0 : l+1]
	s := *h
	s[l] = e
	for i := l; i > 0; i = parent(i) {
		p := parent(i)
		if s[p] <= s[i] {
			break
		}
		s[i], s[p] = s[p], s[i]
	}
}

// pop returns the top element from the heap
func (h *heap) pop() int {
	s := *h
	l := len(s)
	r := s[0]
	s[0] = s[l-1]
	*h = (*h)[0 : l-1]
	s.heapify(0)
	return r
}

// create takes the given slice and creates a heap type that uses the slice for storage
func create(s []int) *heap {
	res := heap(s[0:0])
	return &res
}


func parent(x int) int { return (x - 1) / 2 }
func left(x int) int   { return 2*x + 1 }
func right(x int) int  { return 2*x + 2 }
func (h heap) heapify(i int) {
	small := i
	if left(i) < len(h) && h[left(i)] < h[i] {
		small = left(i)
	}
	if right(i) < len(h) && h[right(i)] < h[small] {
		small = right(i)
	}
	if small != i {
		h[i], h[small] = h[small], h[i]
		h.heapify(small)
	}
}


// Test functions

func tpush(h *heap, x int) {
	fmt.Println("Adding number", x)
	h.insert(x)
	fmt.Println(*h)
}
func tpop(h *heap) {
	fmt.Println("Popping number", h.pop())
	fmt.Println(*h)
}

func main() {
	h := create(make([]int, 100))

	tpush(h, 15)
	tpush(h, 5)
	tpush(h, 73)
	tpush(h, 27)
	tpush(h, 2)
	tpush(h, 17)
	tpop(h)
	tpop(h)
	tpush(h, 62)
	for len(*h) > 0 {
		tpop(h)
	}
}
