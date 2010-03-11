
package main

import (
	"fmt"; 
)
func main() {
	store := make([]int,10)
	myHeap := create(store)
//	fmt.Println(myHeap.front)
//	fmt.Println(myHeap.storage[myHeap.front])
	
	myHeap.insert(2)
	myHeap.insert(2)

	//fmt.Println(myHeap.front)
	//fmt.Println(myHeap.storage[0])
	for i:=0; i<len(store)-1; i++ {
		fmt.Println(myHeap.storage[i])
	}
	val := myHeap.pop()
	fmt.Println("poped value:",val)
	for i:=0; i<len(store)-1; i++ {
		fmt.Println(myHeap.storage[i])
	}
}

type heap struct {
	front int
	storage []int
}
// insert adds element e to the heap
func (h *heap) insert(e int) {
	h.storage[h.front] = e
	h.front++
}

// pop returns the top element from the heap
func (h *heap) pop() int {
	h.front--
	ret := h.storage[h.front]
	h.storage[h.front] = 0
	return ret
}

// create takes the given slice and creates a heap type that uses the slice for storage
func create(s []int) *heap {
	return &heap{0,s}
}
