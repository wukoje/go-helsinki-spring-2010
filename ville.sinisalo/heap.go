package main

import (
	"fmt"
	"rand"
	"time"
)

type heap struct {
	l int
	s []int
}

func main() {
	h := create(make([]int, 300))
        /* Add 300 random numbers [0,1000) */
        rand.Seed(time.Nanoseconds())
	for i:=0; i<300; i++ {
		h.insert(rand.Intn(1000))
	}
	fail := rand.Intn(5); 
	if fail == 4 {
		/* Sometimes try to insert too much */
		h.insert(rand.Intn(1000))
	}
	for i:=0; i<300; i++ {
		fmt.Printf("%d ", h.pop())
	}
	if fail == 3 {
		/* Sometimes try to pop too much */
		fmt.Printf("%d ", h.pop())
	}
	fmt.Printf("\n")

}


// insert adds element e to the heap
func (h *heap) insert(e int) {
	if h.l == len(h.s) {
		panic("Heap overflow")
	}
	i := h.l
	for ;i>0; i/=2 {
		if h.s[i/2] > e {
			h.s[i] = h.s[i/2]
		} else {
			break
		}
	}
	h.l++
	h.s[i] = e
}

// pop returns the top element from the heap
func (h *heap) pop() int {
	if h.l == 0 {
		panic("Attemt to pop from empty heap!")
	}
	ret := h.s[0]
	h.l--
	h.s[0] = h.s[h.l]
	for i:=0;true; {
		var smallest int
		left := i*2
		right := i*2 + 1
		if left < h.l && h.s[left] < h.s[i] {
			smallest = left
		} else {
			smallest = i
		}
		if right < h.l && h.s[right] < h.s[smallest] {
			smallest = right
		}
		if smallest != i {
			h.s[i], h.s[smallest] = h.s[smallest], h.s[i] 
			i = smallest
		} else {
			break;
		}
	}
	return ret
}


// create takes the given slice and creates a heap type that uses the slice for storage
func create(s []int) *heap {
	return &heap{ 0, s }
}
