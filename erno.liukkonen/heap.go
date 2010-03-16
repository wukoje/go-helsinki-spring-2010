package main

import fmt "fmt"

func main() {

	h := create(make([]int,30))
	
	fmt.Printf("Insert %d\n", 15)
	h.insert(10)
	fmt.Printf("Insert %d\n",6)
	h.insert(5)
	fmt.Printf("Insert %d\n",23)
	h.insert(20)

	
	for i:=0; i < 3; i++ {
		fmt.Printf("Pop %d \n",h.pop())
	}

} 

type heap struct {

	array []int
	max int
	size int
}

func create (s []int) *heap {
	h := new (heap)
	h.array = s
	h.size = 0
	h.max = len(s)
	return h
}

func (h * heap) insert(e int) {
	if (h.max >= h.size) {
		h.array[h.size] = e
		h.size++

	 	node := h.size - 1
	 	for {
	 		if node == 0 {
			 	break
			 }
	 
			parent := (node - 1) / 2

			if h.array[parent] < h.array[node] {
			h.array[parent], h.array[node] = h.array[node], h.array[parent]
			node = parent
			} else {
				break
			}
	 			   
		    }
		 }
}


func (h *heap) pop() int {
	value := h.array[0]
	h.array[0] = h.array[h.size-1]
	h.array[h.size-1] = 0
	h.size--

	left := 0
	right := 0
	x := 0

	for {
		parent := h.array[x]
		left = h.array[x*2+1]
		right = h.array[x*2+2]

		 if parent >= left && parent >= right {
		 	break

	 	} else if parent < left || parent < right {

		 	if left >= right {
		 	h.array[x], h.array[x*2+1] = left, parent
			x = x*2+1
		} else {
		 	h.array[x], h.array[x*2+2] = right, parent
		 	x = x*2+2
		 }
		}
	 }
	 return value
}
