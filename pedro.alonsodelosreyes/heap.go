package main

import (
	"fmt"
)

func main() {

//	fmt.Printf("Epic win!!!11oneoneeleveneleven\n")
	a := []int{8, 20, 9, 7, 17, 19, 25, 36, 100}
	h := create(a)
	fmt.Println(h.print(), h.size)
	//	h.pop()
	//	fmt.Println(h.pop(),h.print())
	//	h.pop()
	//	h.pop()
	fmt.Println("Inserting values 40, 60, 30, 10, 1, 4, 100, 70")
	h.insert(40)
	fmt.Println(h.print())
	h.insert(60)
	fmt.Println(h.print())
	h.insert(30)
	h.insert(10)
	h.insert(1)
	h.insert(4)
	h.insert(100)
	h.insert(70)
	fmt.Println(h.print())
	fmt.Println("Deleting smallest values 1 then 4")
	h.pop()
	fmt.Println(h.print())
	h.pop()
	fmt.Println(h.print())

}


type heap struct {
	capacity, size int
	elem           []int
}

func (h *heap) parent(e int) int {

	p := 0
	if e != 0 {

		p = (e - 1) / 2
	}
	return p
}


// insert adds element e to the heap
func (h *heap) insert(e int) {

	if h.size == h.capacity {

	} else {

		h.elem[h.size] = e
		pos := h.size
		for {
			if pos == 0 {
				break
			}
			if h.elem[pos] < h.elem[h.parent(pos)] {

				a := h.elem[pos]
				h.elem[pos] = h.elem[h.parent(pos)]
				h.elem[h.parent(pos)] = a
				pos = h.parent(pos)

			} else {
				break
			}
		}
		h.size++

	}

}

//pop returns the smallest element from the heap
func (h *heap) pop() int {

	p := 0
	if h.size > 0 {

		p = h.elem[0]
		h.elem[0] = h.elem[h.size-1]
		h.elem[h.size] = 0
		h.size -= 1
		h.heapify(0)

	} else {

		p = 0
	}
	return p

}

func (h *heap) print() []int { return h.elem[0:h.size] }

// create takes the given slice and creates a heap type that uses the slice for storage
func create(s []int) *heap {

	h := new(heap)
	h.elem = s
	h.size = 0
	h.capacity = cap(s)
	return h

}

func (h *heap) heapify(node int) {

	//fmt.Println("test")
	for {

		left := 2*(node+1) - 1
		right := 2 * (node + 1)
		smallest := node
		//fmt.Println(left, " ", h.size ," ", h.elem[left], " > ", h.elem[smallest], " node " , node, "smallest ", smallest)
		if left < h.size && h.elem[left] < h.elem[smallest] {

			smallest = left

		}
		//fmt.Println(right, " ", h.size ," ", h.elem[right], " > ", h.elem[smallest], " node ", node," smallest ", smallest)
		if right < h.size && h.elem[right] < h.elem[smallest] {

			smallest = right

		}
		if smallest == node {

			break
		}
		h.elem[node], h.elem[smallest] = h.elem[smallest], h.elem[node]
		node = smallest
		//fmt.Println(h.print(),"ldlldld")
	}

}
