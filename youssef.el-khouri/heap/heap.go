package main


import fmt "fmt"

//Tein max-heapi, koska tehtävän annossa ei ollut muuta mainintaa
func main() {

	s := make([]int, 30)
	h := create(s)		
	fmt.Printf("insert:....")
	h.insert(10)
	h.insert(20)
	h.insert(30)
	h.insert(40)
	h.insert(50)
	h.insert(60)
	h.insert(70)
	h.insert(80)
	h.insert(90)
	h.insert(100)
	h.insert(70)
	h.insert(50)
	
	fmt.Printf("done: \n")
	
	for _,j := range h.heap {
		fmt.Printf("%d,", j)
	}
	
	fmt.Printf("\npop: %d\n", h.pop())
	
	for _,j := range h.heap {
		fmt.Printf("%d,", j)
	}
	
	fmt.Printf("\npop: %d\n", h.pop())
	
	for _,j := range h.heap {
		fmt.Printf("%d,", j)
	}

	fmt.Printf("\npop: %d\n", h.pop())
	
	for _,j := range h.heap {
		fmt.Printf("%d,", j)
	}

	fmt.Printf("\npop: %d\n", h.pop())
	
	for _,j := range h.heap {
		fmt.Printf("%d,", j)
	}

	fmt.Printf("\npop: %d\n", h.pop())
	
	for _,j := range h.heap {
		fmt.Printf("%d,", j)
	}

	fmt.Printf("\npop: %d\n", h.pop())
	
	for _,j := range h.heap {
		fmt.Printf("%d,", j)
	}

}

type heap struct {
	heap []int
	tail int
}

// insert adds element e to the heap
func (h *heap) insert(e int) {
	if h != nil {
		h.heap[h.tail] = e
			
		
		current := h.tail
		parent := current/2
		
		for current != parent {	
			if e > h.heap[parent] {
				h.heap[current], h.heap[parent] = h.heap[parent], h.heap[current]
			} else {
				break
			}
			current = parent
			parent = current/2
		}
	}
	h.tail++
}

// pop returns the top element from the heap
func (h *heap) pop() int {
	if h != nil {
		temp := h.heap[0]
		h.tail--		
		h.heap[0], h.heap[h.tail] = h.heap[h.tail], 0
				
		
		left, right, currentMax := 1, 2, 0
		for i := 0; right <= h.tail;{
			if h.heap[right] <= h.heap[left] {
				currentMax = left
			} else {
				currentMax = right
			}
			if h.heap[i] < h.heap[currentMax] {
				h.heap[i], h.heap[currentMax] = h.heap[currentMax], h.heap[i]
				i = currentMax
				left = 2*i+1
				right = 2*i+2			
			} else {
				return temp
			}

			if i == currentMax {
				currentMax = 0
			} else {
				return temp
			}
		}
		return temp
	}
	return -1
}


// create takes the given slice and creates a heap type that uses the slice for storage
func create(s []int) *heap {
	if s != nil {
		heap := heap{s, 0}
		return &heap	
	}
	return nil
	
}

