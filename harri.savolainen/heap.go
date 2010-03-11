package main

import fmt "fmt" 

func main() { 
	
	// Some data for heap
	s := []int{99,1,777,20,5,40,30,100}
	
	// Create heap
	h := create(make([]int,30))
	
	// Insert values
	for _,v := range s {
		fmt.Printf("Inserting %d\n",v)
		h.insert(v)
	}

	// Pop values
	for i:=0; i<len(s); i++ {
		fmt.Printf("Popping %d \n",h.pop())
	}

} 

type heap struct {
	data []int
	heapSize int
	arraySize int
}

func getLeftChildIndex(nodeIndex int) int {
	return 2*nodeIndex+1
}

func getRightChildIndex(nodeIndex int) int {
	return 2*nodeIndex+2
}

func getParentIndex(nodeIndex int) int {
	return (nodeIndex -1)/2
}

func create(s []int) *heap {
	h := new(heap)
	h.data = s
	h.heapSize = 0
	h.arraySize = len(s)
	return h
}

func (h *heap) getMinimum() int {
	if h.isEmpty() { 
		panic ("heap is empty") 
	} else {
		return h.data[0]
	}
	panic("unreachable")
}

func (h *heap) isEmpty() bool {
	return h.heapSize == 0
}

func (h *heap) shiftUp(nodeIndex int) {
	if (nodeIndex != 0) {
		parentIndex := getParentIndex(nodeIndex)
		if h.data[parentIndex] > h.data[nodeIndex] {
			tmp := h.data[parentIndex]
			h.data[parentIndex] = h.data[nodeIndex]
			h.data[nodeIndex] = tmp
			h.shiftUp(parentIndex)
		}
	}
}


func (h *heap) insert(e int) {
	if h.heapSize == h.arraySize {
		panic ("heap overflow")	
	} else {
		h.heapSize++
		h.data[h.heapSize-1] = e
		h.shiftUp(h.heapSize-1)
	}
}


func (h *heap) shiftDown(nodeIndex int) {
	leftChildIndex := getLeftChildIndex(nodeIndex)
	rightChildIndex := getRightChildIndex(nodeIndex)
	var minIndex int	

	if rightChildIndex >= h.heapSize {
		if leftChildIndex >= h.heapSize {
			return
		} else {
			minIndex = leftChildIndex
		}
	} else {
		if h.data[leftChildIndex] <= h.data[rightChildIndex] {
			minIndex = leftChildIndex
		} else {
			minIndex = rightChildIndex
		}
	} 		

	if h.data[nodeIndex] > h.data[minIndex] {
		tmp := h.data[minIndex]
		h.data[minIndex] = h.data[nodeIndex]
		h.data[nodeIndex] = tmp
		h.shiftDown(minIndex)
	}
}

func (h *heap) pop() int {
	if h.isEmpty() {
		panic("Heap empty")
	} else {
		value := h.data[0]
		h.data[0] = h.data[h.heapSize-1]
		h.heapSize--
		if h.heapSize > 0 {
			h.shiftDown(0)
		}
		return value
	}
	panic ("unreachable")
}
