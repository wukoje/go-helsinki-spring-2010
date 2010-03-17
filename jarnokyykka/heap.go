package main

import (
	"fmt"
)

type heap struct {
	s []int
	length int
}


func main() {
	s := []int{100,15,12,45,5,99,58,-1}

	fmt.Printf("Test slice: ")
	for _, v := range s {
		fmt.Printf("%d ", v)
	} 
	
	h := create(s)

	fmt.Printf("\nPopping all values!\n")
	for i := h.length; i > 0; i-- {
		fmt.Printf("POP: %d\n", h.pop())
	}

}

// insert adds element e to the heap
func (h *heap) insert(e int) {
	if h.s == nil {
		h.s = make([]int, 5)
	}

	//double the size if needed
	if len(h.s) == h.length {
		fmt.Printf("\n Array full! Doubling the size\n")
		var s2 []int = make([]int, len(h.s)*2) 
		for i, v := range h.s {
			s2[i] = v
		}
		h.s = s2
	} 
	

	h.s[h.length] = e
	h.length = h.length+1
	current := h.length
	parent := getParentIndex(current)	

	//if parent is higher and not head, swap
	for ((current > 1) && (h.getValueAtIndex(current)) <= (h.getValueAtIndex(parent))) {
		h.s[current-1], h.s[parent-1] = h.s[parent-1], h.s[current-1]
		current = parent
		parent = getParentIndex(current)
	}
}

// pop returns the top element from the heap
func (h *heap) pop() int {
	
	returnVal := h.getValueAtIndex(1);
	current :=1 
	
	//set last to first position and decrese length
	h.s[0] = h.getValueAtIndex(h.length)
	h.length = h.length-1

	
	for getLeftIndex(current) <= h.length {
		left := getLeftIndex(current)
		right := getRightIndex(current)
		move := left
		
		//if right is smaller 
		if (right <= h.length && h.getValueAtIndex(left) > h.getValueAtIndex(right)) {
			move = right
		}

		//swap
		if h.getValueAtIndex(current) >= h.getValueAtIndex(move) {
			h.s[current-1], h.s[move-1] = h.s[move-1], h.s[current-1]
		}

		current = move
	}

	return returnVal
}

// create takes the given slice and creates a heap type that uses the slice for storage
func create(s []int) *heap {
	h := new(heap)
	h.length = 0
	for _, v := range s {
		h.insert(v)
	}
	fmt.Printf("HEAP created! Array looks: ")
	for _, v := range h.s {
		fmt.Printf(" %d ", v)
	}
	return h
}

func getParentIndex(node int) int {
	return node/2
}

func getLeftIndex(parent int) int {
	return 2*parent
}

func getRightIndex(parent int) int {
	return ((2*parent)+1)
}

func (h *heap) getValueAtIndex(index int) int {
	return h.s[index-1]
}

