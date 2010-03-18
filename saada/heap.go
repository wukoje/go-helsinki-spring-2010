package main

import (
	"fmt"
	"math"
	"exp/exception"
)

type heap []int

func (h *heap) insert(e int) {
	if x := exception.Try(func(throw exception.Handler) {
		if len(*h) == cap(*h) {
			throw("overflow")
		}
		i := len(*h)
		*h = (*h)[0:len(*h)+1]
		for i > 0 && (*h)[parent(i)] > e {
			(*h)[i] = (*h)[parent(i)]
			i = parent(i)
		}
		(*h)[i] = e
	}); x != nil {
		fmt.Println(x.Value);
	}
}

func (h *heap) pop() int {
	if len(*h) == 0 {
		return -1
	}
	minimum := (*h)[0]
	(*h)[0] = (*h)[len(*h)-1]
	*h = (*h)[0:len(*h)-1]
	heapify(h, 0)
	return minimum
}

func create(s []int) *heap {
	h := heap(s[0:0])
	return &h
}

func heapify(h *heap, i int) {
	var next int
	l := left(i)
	r := right(i)
	if l < len(*h) && (*h)[l] < (*h)[i] {
		next = l
	} else {
		next = i
	}
	if r < len(*h) && (*h)[r] < (*h)[next] {
		next = r
	}
	if next != i {
		(*h)[i], (*h)[next] = (*h)[next], (*h)[i]
		heapify(h, next)
	}
}

func parent(i int) int {
	return int(math.Floor(float64(i/2)))
}

func left(i int) int {
	return 2*i+1
}

func right(i int) int {
	return 2*i+2
}

func main() {
	h := create(make([]int, 100))
	fmt.Println(h)
	h.insert(11)
	fmt.Println(h)
	h.insert(21)
	fmt.Println(h)
	h.insert(14)
	fmt.Println(h)
	h.insert(7)
	fmt.Println(h)
	fmt.Println(h.pop())
	fmt.Println(h)
	h.insert(9)
	fmt.Println(h)
	fmt.Println(h.pop())
	fmt.Println(h)
	fmt.Println(h.pop())
	fmt.Println(h)
	fmt.Println(h.pop())
	fmt.Println(h)
	fmt.Println(h.pop())
	fmt.Println(h)
	fmt.Println(h.pop())
	fmt.Println(h)
	fmt.Println(h.pop())
	fmt.Println(h)
	h2 := create(make([]int, 2))
	h2.insert(9)
	fmt.Println(h2)
	h2.insert(31)
	fmt.Println(h2)
	h2.insert(4)
	fmt.Println(h2)
}
