package main

import(
	"fmt"
	"math"
	"strconv"
)

type heap struct{
	data []int
	tail int
	head int
}

// insert adds element e to the heap
func (h *heap) insert(e int) {
	if h.tail == h.head { 
		fmt.Println("Error: heap is full!"); return 
}

	h.data[h.tail] = e
	h.siftup(h.tail)

	h.tail++
}


// pop returns the smallest element from the heap
func (h *heap) pop() int {
	if h.tail == 0 {
		fmt.Println("Error: heap is empty!")
		return -1
	}

	elem := h.data[0]
	h.tail--
	h.data[0] = h.data[h.tail]
	h.siftdown(0)

	return elem
}

// sift up the element to maintain heap
func (h *heap) siftup(i int) {
	if i == 0 { return }

	for i > 0 {
		parent_idx := (i-1)/2
		if h.data[parent_idx] > h.data[i] {
			h.data[parent_idx], h.data[i] = h.data[i], h.data[parent_idx]
			i = parent_idx
			continue
		}
		break
	}
}

// sift down the element to maintain heap
func (h *heap) siftdown(i int) {
	if (2*i+1) >= h.tail { return }

	for (2*i+1) < h.tail {
		i = 2*i+1	//index of left child
		if ((i+1) < h.tail) && (h.data[i+1] < h.data[i]) {
			i++
		}

		if h.data[(i-1)/2] > h.data[i] {
			h.data[(i-1)/2],h.data[i] = h.data[i],h.data[(i-1)/2]
			continue
		}
		break
	}
}

// return string to represent information about heap
func (h *heap) String() string {
	if h.tail == 0 { return "Heap: empty\n" }

	s := "Heap: ["
	for i:=0; i<h.tail; i++ {
		s += strconv.Itoa(h.data[i]) + ","
	}
	s = s[0:len(s)-1] + "]\n"

	return s
}

// return string to represent heap as a tree structure
func (h *heap) treeString() string {
	if h.tail == 0 { return "" }

	s := "Tree structure of heap:\n"
	layer := int(math.Log2(float64(h.tail)))

	current := 0
	for i:=0; i<=layer; i++ {
		for j:=0; j<int(math.Pow(2,float64(i))) && current<h.tail; j++ {
			s += strconv.Itoa(h.data[current]) + "\t"
			current++
		}
		s += "\n"
	}
	return s
}

// create takes the given slice and creates a heap type that uses the slice for storage
func create(s []int) *heap {
	data := make([]int, len(s))
	tail := 0
	head := len(s)

	return &heap{data, tail, head}
}

func main(){
	s := make([]int ,30)
	heap := create(s)

	fmt.Println("#Test case 1: insert element to heap")
	fmt.Println("insert order: 25,30,4,18,15,33,10,77")
	heap.insert(25)
	heap.insert(30)
	heap.insert(4)
	heap.insert(18)
	heap.insert(15)
	heap.insert(33)
	heap.insert(10)
	heap.insert(77)
	fmt.Println("Insertion done -- ",heap)
	fmt.Println(heap.treeString())

	fmt.Println("#Test case 2: pop element in heap one by one")
	for i:=1;i<=8;i++ {
	fmt.Println("Pop out: ", heap.pop())
	fmt.Println("after pop -- ",heap)
	}
	fmt.Println("Pop out: ", heap.pop())
	}
