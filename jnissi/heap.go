package main

import "fmt"

type heap struct {
	heap []int
	tail int
}

func (h *heap) insert(e int) {
	if(h.tail == len(h.heap)) {
		return
	}
	h.heap[h.tail] = e;
	pos := h.tail
	for {
		if pos == 0 {
			break
		}
		parentPos := (pos - 1) / 2
		if h.heap[pos] > h.heap[parentPos] {
			h.heap[pos], h.heap[parentPos] = h.heap[parentPos], h.heap[pos]
			pos = parentPos
		} else {
			break
		}
	}
	h.tail++
}

func (h *heap) pop() int {
	if h.tail == 0 && h.heap[h.tail] == 0 {
		return 0
	}
	ret := h.heap[0]
	h.heap[0], h.heap[h.tail - 1] = h.heap[h.tail - 1], 0
	h.tail--
	pos := 0
	for {
		l, r := pos * 2 + 1, pos * 2 + 2
		hi := pos
		if l < h.tail && h.heap[l] > h.heap[hi] {
			hi = l
		}
		if r < h.tail && h.heap[r] > h.heap[hi] {
			hi = r
		}
		if hi == pos {
			break
		}
		h.heap[hi], h.heap[pos] = h.heap[pos], h.heap[hi]
		pos = hi
	}
	
	return ret
}

func Create(s []int) *heap {
	if s == nil {
		return nil
	}
	return &heap{s,0}
}

func main() {
	back := make([]int,10)
	heap := Create(back)
	//Tries to add more than fits. No effect.
	for _,v := range []int{10,20,30,15,50,12,24,33,51,19,21} {
		heap.insert(v)
		fmt.Printf("%v\n", back)
	}
	//Tries to pop more than available. Gets 0 the 11th time.
	for i := 0; i < 11; i++ {
		n := heap.pop()
		fmt.Printf("Popped:%d\n%v\n", n, back)
	}
}
