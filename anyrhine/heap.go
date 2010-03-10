package main

import "fmt"
import "rand"

type heap struct {
	tab []int
}

func (h *heap) insert(e int) {
	last := len(h.tab)
	if last+1 > cap(h.tab) {
		ntab := make([]int, cap(h.tab)+10)
		for i, c := range h.tab {
			ntab[i] = c
		}
		h.tab = ntab
	}
	h.tab = h.tab[0 : last+1]
	var i int
	for i = last; i > 1 && h.tab[i/2] < e; i /= 2 {
		h.tab[i] = h.tab[i/2]
	}
	h.tab[i] = e
}

func (h *heap) pop() int {
	i := 1
	e := h.tab[len(h.tab)-1]
	top := h.tab[i]
	h.tab = h.tab[0 : len(h.tab)-1]
	for i < len(h.tab) {
		var ni int
		if 2*i+1 < len(h.tab) && h.tab[2*i] < h.tab[2*i+1] {
			ni = 2*i + 1
		} else if 2*i < len(h.tab) {
			ni = 2 * i
		} else {
			h.tab[i] = e
			break
		}
		if e < h.tab[ni] {
			h.tab[i] = h.tab[ni]
			i = ni
		} else {
			h.tab[i] = e
			break
		}
	}
	return top
}

func create(s []int) *heap {
	h := new(heap)
	h.tab = s[0:1] // s[-1:0]? no go.
	return h
}

func main() {
	h := create(make([]int, 10))
	for i := 0; i < 100; i++ {
		h.insert(rand.Int() % 1000)
	}
	for i := 0; i < 100; i++ {
		fmt.Printf("%d\n", h.pop())
	}
}
