package main

import fmt "fmt"




func main() {
	array := make([]int, 40)
	test := create(array)
	pop := 0
	
	fmt.Println("Inserting 6")
	test.insert(6)
	test.print()
	
	fmt.Println("Inserting 8")
	test.insert(8)
	test.print()
	
	fmt.Println("Inserting 20")
	test.insert(20)
	test.print()
	
	fmt.Println("Inserting 11")
	test.insert(11)
	test.print()
	
	fmt.Println("Inserting 2")
	test.insert(2)
	test.print()
	
	fmt.Println("Inserting 32")
	test.insert(32)
	test.print()
	
	fmt.Println("Inserting 8")
	test.insert(8)
	test.print()
	
	pop = test.pop()
	fmt.Println("Popped number: ", pop)
	test.print()
	
	pop = test.pop()
	fmt.Println("Popped number: ", pop)
	test.print()
	
	fmt.Println("Inserting 8")
	test.insert(8)
	test.print()
	
	fmt.Println("Inserting 88")
	test.insert(88)
	test.print()
	
	fmt.Println("Inserting 15")
	test.insert(15)
	test.print()
	
	fmt.Println("Inserting 56")
	test.insert(56)
	test.print()
	
	fmt.Println("Inserting 99")
	test.insert(99)
	test.print()
	
	pop = test.pop()
	fmt.Println("Popped number: ", pop)
	test.print()
	
	pop = test.pop()
	fmt.Println("Popped number: ", pop)
	test.print()
	
	pop = test.pop()
	fmt.Println("Popped number: ", pop)
	test.print()
	
	pop = test.pop()
	fmt.Println("Popped number: ", pop)
	test.print()
}





type heap struct {
	array []int
	//temp int
	max int
	amount int
}






func create(s []int) *heap {
	h := new(heap)
	h.array = s
	h.max = len(s)
	h.amount = 0
	
	return h
}






func (h *heap) insert(e int) {
	if h.max > h.amount {
		h.array[h.amount] = e
		h.amount++
		
		current := h.amount - 1
		for {
			if current == 0 {
				break
			}
		
			parent := (current - 1) / 2
			
			if h.array[current] > h.array[parent] {
				h.array[parent], h.array[current] = h.array[current], h.array[parent]
				current = parent
			} else {
				break
			}
		}
	}
}






func (h *heap) pop() int {
	popped := h.array[0]
	h.array[0] = h.array[h.amount-1]
	h.array[h.amount-1] = 0
	h.amount--

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
	return popped	
}





//	func (h *heap) String() string
//	Got your comment, on printing in Go.
//	I've never programmed in C,
//	so fmt.Printf("%v", h) looks
//	a bit weird to me =D 
//	Decided to keep the current function
//	this time.

func (h *heap) print() {
	fmt.Println(h.array[0])
	
	level := 2
	counter := 0
	
	for i := 1; i <= h.amount-1; i++ {
		fmt.Print(h.array[i], " ")
		
		counter++
		if level == counter {
			fmt.Println()
			level = level * 2
			counter = 0
		}
	}
	fmt.Println()
	fmt.Println()
}
