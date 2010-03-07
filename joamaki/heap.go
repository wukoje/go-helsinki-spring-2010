package main

import (
	"fmt"
	"os"
	"rand"
)

type heap struct {
	arr []int
	n int
	nmax int
}

func error(s string) {
	fmt.Println(s)
	os.Exit(1)
}

func create(s []int) *heap {
	h := new(heap)
	h.arr = s
	h.n = 0
	h.nmax = len(s)
	return h
}

func (h *heap) shiftup(i int) {
	for {
		parent := (i-1)/2 // parent of arr[i] or arr[0] if i==0
		if h.arr[i] < h.arr[parent] {
			h.arr[parent], h.arr[i] = h.arr[i], h.arr[parent]
			i = parent
		} else {
			break
		}
	}
}

func (h *heap) shiftdown(i int) {
	for {
		left := i*2+1  // left child of arr[i]
		right := i*2+2 // right child of arr[i]
		smaller := i

		if left < h.n && h.arr[left] < h.arr[smaller] {
			smaller = left
		}

		if right < h.n && h.arr[right] < h.arr[smaller] {
			smaller = right
		}

		if smaller != i {
			// a child found to be smaller then i, swap it
			h.arr[i], h.arr[smaller] = h.arr[smaller], h.arr[i]
			i = smaller
		} else {
			break
		}
	}
}

func (h *heap) insert(e int) {
	if h.n >= h.nmax {
		error("heap overflow")
	}

	h.arr[h.n] = e
	h.n++
	h.shiftup(h.n-1)
}

func (h *heap) pop() int {
	if h.n <= 0 {
		error("heap underflow")
	}

	e := h.arr[0]
	h.arr[0] = h.arr[h.n-1]
	h.n--

	h.shiftdown(0)

	return e
}

func main() {
	h := create(make([]int, 10000))

	for _, c := range rand.Perm(10000) {
		h.insert(c)
	}

	for i := 0; i < 10000; i++ {
		res := h.pop()
		if (res != i) {
			fmt.Println("Expected ", i, " got ", res)
			return
		}
	}

}
