package main

import "fmt"

func main() {
	h := create(make([]int, 16))
	h.insert(3)
	h.insert(2)
	h.insert(5)
	h.insert(4)
	fmt.Println(h.s)

	h.pop()
	//h.pop()
	fmt.Println(h.s)
}

type heap struct {
	s []int		//slice
	i int		//index
}

// insert adds element e to the heap
func (h *heap) insert(e int) {
	h.s[h.i], h.i= e, h.i+1
	j := h.i-1
	
	for ; h.s[j]>h.s[pN(j)]; j=pN(j) {
		h.s[j], h.s[pN(j)] = h.s[pN(j)], h.s[j]
	}
}

// pop returns the top element from the heap
func (h *heap) pop() (n int) {
	n, h.s[0], h.i = h.s[0], h.s[h.i-1], h.i-1
	h.s[h.i] = 0	//only for test, not necessary
	j := 0
	k := 0
	for {
		if h.s[j]<h.s[lN(j)] && lN(j)<h.i { k = lN(j) }
		if h.s[j]<h.s[rN(j)] && rN(j)<h.i { k = rN(j) }
		if j!=k {
			if h.s[lN(j)]<h.s[rN(j)] {
				k = rN(j)
			} else {
				k = lN(j)
			}
			h.s[j], h.s[k] = h.s[k], h.s[j]
			j = k
		} else {
			break
		}
	}
	return n
}

// create takes the given slice and creates a heap type that uses the slice for storage
func create(s []int) *heap {
	h := new(heap)
	h.s, h.i = s, 0
	return h
}

// tools function
func pN(index int) int { return (index-1)/2 }	//index for the parent node
func lN(index int) int { return 2*index+1 }		//index for the left node
func rN(index int) int { return 2*index+2 }		//index for the right node

