// binary min-heap
package main

import (
	"fmt"
)

type heap []int

func (h *heap) left(i int) int   { 
	return 2*i + 1
}

func (h *heap) right(i int) int  { 
	return 2*i + 2
}

func (h *heap) parent(i int) int { 
	return (i - 1) / 2
}

func (h *heap) insert(e int) {
	i := len(*h)
	*h = (*h)[0:len(*h)+1]
	
	for i > 0 && (*h)[h.parent(i)] > e {
		(*h)[i] = (*h)[h.parent(i)]
		i = (*h).parent(i)
	}
	(*h)[i] = e
}

func (h *heap) pop() int {
	tmp := (*h)[0]
	(*h)[0] = (*h)[len(*h)-1]
	*h = (*h)[0:len(*h)-1]
	if len(*h) > 1 {
		(*h).heapify(0)
	}
	return tmp
}

func (h *heap) heapify(i int) {
	j := i
	x := (*h)[i]
	for (*h).left(j) < len(*h) {
		l := (*h).left(j)
		if h.right(j) < len(*h) && (*h)[(*h).right(j)] < (*h)[l] {
			l = (*h).right(j)
		}
		if (*h)[l] < x {
			(*h)[j] = (*h)[l]
			j = l
		} else {
			(*h)[j] = x
			return
		}
	}
	(*h)[j] = x
}

func main() {
	h := make(heap,0, 25)

	for _, v := range []int{9,6,7,4,5,8,10} {
		fmt.Printf("Insert: %v\n", v)
		h.insert(v)
	}
	fmt.Printf("Array: %v\n", h[0:len(h)])
	fmt.Printf("Pop: %v\n", h.pop())
	fmt.Printf("Pop: %v\n", h.pop())
	fmt.Printf("Pop: %v\n", h.pop())
	fmt.Printf("Array: %v\n", h[0:len(h)])
	for _, v := range []int{2,15,1,30,0} {
		fmt.Printf("Insert: %v\n", v)
		h.insert(v)
	}
	fmt.Printf("Array: %v\n", h[0:len(h)])
	for len(h) > 0 {
		fmt.Printf("Pop: %v\n", h.pop())
		fmt.Printf("Array: %v\n", h[0:len(h)])
	}
}
