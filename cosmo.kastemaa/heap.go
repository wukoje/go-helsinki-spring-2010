package main

import "fmt"

func main() {
	h := create(make([]int, 7))
	fmt.Printf("%v\n\n", h)
	testInsert := func(e int) {
		fmt.Printf("Inserting %v... ", e)
		if h.insert(e) {
			fmt.Printf("ok\n")
		} else {
			fmt.Printf("failed\n")
		}
		fmt.Printf("%v\n\n", h)
	}
	testPop := func() {
		fmt.Printf("Popping... ")
		e, ok := h.pop()
		if ok {
			fmt.Printf("ok, got %v\n", e)
		} else {
			fmt.Printf("failed\n")
		}
		fmt.Printf("%v\n\n", h)
	}
	testInsert(4)
	testInsert(8)
	testInsert(6)
	testInsert(5)
	testInsert(7)
	testInsert(3)
	testInsert(1)
	testInsert(0)
	testPop()
	testPop()
	testPop()
	testPop()
	testPop()
	testPop()
	testPop()
	testPop()
}

type heap struct {
	size int
	slice []int
}

func (h *heap) insert(e int) bool {
	if h.size < len(h.slice) {
		i := h.size
		for ; i > 0 && h.slice[parent(i)] > e; i = parent(i) {
			h.slice[i] = h.slice[parent(i)]
		}
		h.slice[i] = e
		h.size++
		return true
	}
	return false
}

func parent(i int) int {
	return i / 2
}

func leftChild(i int) int {
	return 2 * i
}

func rightChild(i int) int {
	return 2 * i + 1
}

func (h *heap) pop() (int, bool) {
	if h.size > 0 {
		min := h.slice[0]
		h.slice[0] = h.slice[h.size - 1]
		h.size--
		h.minHeapify(0)
		return min, true
	}
	return 0, false
}

func (h *heap) minHeapify(i int) {
	l := leftChild(i)
	r := rightChild(i)
	var next int
	if l <= h.size && h.slice[l] < h.slice[i] {
		next = l
	} else {
		next = i
	}
	if r <= h.size && h.slice[r] < h.slice[next] {
		next = r
	}
	if next != i {
		h.slice[i], h.slice[next] = h.slice[next], h.slice[i]
		h.minHeapify(next)
	}
}

func create(s []int) *heap {
	return &heap{size: 0, slice: s}
}

func (h *heap) String() string {
	return fmt.Sprintf("heap(%v, %v)", h.size, h.slice[0:h.size])
}
