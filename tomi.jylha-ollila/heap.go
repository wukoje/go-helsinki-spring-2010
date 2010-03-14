package main


import (
	"fmt"
	"rand"
)


type heap struct {
	elems []int
}


func (h *heap) insert(e int) {
	elem_count := len(h.elems)
	if elem_count == cap(h.elems) {
		s := make([]int, elem_count*2)
		for i, e := range h.elems {
			s[i] = e
		}
		h.elems = s
	}
	h.elems = h.elems[0 : elem_count+1]
	h.elems[elem_count] = e
	h.upheap(elem_count)
}


func (h *heap) pop() int {
	elem_count := len(h.elems)
	if elem_count == 0 {
		return -1
	}
	ret := h.elems[0]
	h.elems[0] = h.elems[elem_count-1]
	h.elems = h.elems[0 : elem_count-1]
	h.downheap(0)
	return ret
}


func left_child(index int) int { return 2*index + 1 }


func right_child(index int) int { return 2*index + 2 }


func parent(index int) int { return (index - 1) / 2 }


func (h *heap) downheap(index int) {
	smallest := index
	for left := left_child(index); left < len(h.elems); left = left_child(index) {
		if h.elems[left] < h.elems[smallest] {
			smallest = left
		}
		right := right_child(index)
		if right < len(h.elems) && h.elems[right] < h.elems[smallest] {
			smallest = right
		}
		if index == smallest {
			break
		}
		h.elems[index], h.elems[smallest] = h.elems[smallest], h.elems[index]
		index = smallest
	}
}


func (h *heap) upheap(index int) {
	par := parent(index)
	for index > 0 && h.elems[index] < h.elems[par] {
		h.elems[index], h.elems[par] = h.elems[par], h.elems[index]
		index = par
		par = parent(index)
	}
}


func create(s []int) *heap {
	h := new(heap)
	h.elems = s[0:0]
	return h
}


func main() {
	h := create(make([]int, 10))
	fmt.Printf("Inserting ")
	for i := 0; i < 20; i++ {
		if i > 0 {
			fmt.Printf(", ")
		}
		num := rand.Int() % 100
		fmt.Printf("%d", num)
		h.insert(num)
	}
	fmt.Printf("\n")
	fmt.Printf("Retrieving ")
	for first, num := true, h.pop(); num >= 0; num = h.pop() {
		if !first {
			fmt.Printf(", ")
		}
		first = false
		fmt.Printf("%d", num)
	}
	fmt.Printf("\n")
}
