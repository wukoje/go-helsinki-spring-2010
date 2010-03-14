package main

import (
	"fmt"
	"rand"
	"time"
)

type heap struct {
	s    []int
	size int
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func (h *heap) parent(i int) int { return max((i-1)/2, 0) }

func (h *heap) left(i int) int { return 2*(i+1)-1 }

func (h *heap) right(i int) int { return 2*(i+1) }

func (h *heap) heapify(i int) {
	l := h.left(i)
	r := h.right(i)
	smallest := i
	if l < h.size && h.s[l] < h.s[i] {
		smallest = l
	}
	if r < h.size && h.s[r] < h.s[smallest] {
		smallest = r
	}
	if smallest != i {
		h.s[i], h.s[smallest] = h.s[smallest], h.s[i]
		h.heapify(smallest)
	} else {
	}
}

func (h *heap) decrease(i int, k int) {
	if h.s[i] < k {
		return
	}
	h.s[i] = k
	for h.s[h.parent(i)] > h.s[i] && i > 0 {
		h.s[i], h.s[h.parent(i)] = h.s[h.parent(i)], h.s[i]
		i = h.parent(i)
	}
}

func (h *heap) pop() int {
	if h.size < 1 {
		return -(1 << 31)
	}
	min := h.s[0]
	h.s[0] = h.s[h.size-1]
	h.size--
	h.heapify(0)
	return min
}

func (h *heap) insert(i int) {
	if h.size > len(h.s) {
		return
	}
	h.size++
	h.s[h.size-1] = 1 << 31 - 1
	h.decrease(h.size-1, i)
}

func create(s []int) *heap {
	h := new(heap)
	h.s = s
	h.size = 0
	return h
}

func main() {
	a := []int{2,4,6,1,2,3,5,7,3,16,41,12,11,10,9,1}
	h := create(make([]int, len(a)))
	fmt.Printf("array: %v\n", a)
	fmt.Printf("made a new heap: %v\n", h)
	for _, c := range a {
		h.insert(c)
	}
	fmt.Printf("heap after inserting all in array: %v\n", h)
	fmt.Printf("now popping values: ")
	for i := h.pop(); i >= 0; i = h.pop() {
		fmt.Printf("%v ", i)
	}
	fmt.Printf("\n\n")

	fmt.Printf("testing heap sort with random arrays... ")
	length := 99999
	rounds := 11
	ok := true
	rand.Seed(time.Nanoseconds())
	l: for i := 0; i < rounds; i++ {
		s := rand.Perm(length)
		h = create(make([]int, length))
		for _, c := range s {
			h.insert(c)
		}
		prev := -1
		for p := h.pop(); p >= 0; p = h.pop() {
			if p < prev {
				fmt.Printf("\nERROR: got %v which is smaller than previous %v!\n", p, prev)
				fmt.Printf("array is %v\n", s)
				fmt.Printf("heap is %v\n", h)
				ok = false
				break l
			}
			prev = p
		}
	}
	if ok {
		fmt.Printf("OK!\n")
	}
}

